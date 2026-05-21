package handlers

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"coverage-service/backend/config"
	"coverage-service/backend/database"
	"coverage-service/backend/models"

	"github.com/gin-gonic/gin"
)

// registeredServices is loaded once at startup from services.json
var registeredServices []config.ServiceConfig

func init() {
	svcs, err := config.LoadServices()
	if err != nil {
		log.Println("Warning: could not load services.json:", err)
		registeredServices = []config.ServiceConfig{}
		return
	}
	registeredServices = svcs
}

// RunTestsSSE streams test output line-by-line via Server-Sent Events,
// then saves the parsed result to the database.
func RunTestsSSE(c *gin.Context) {
	serviceName := c.Param("service")

	var svc *config.ServiceConfig
	for i, s := range registeredServices {
		if s.Name == serviceName {
			svc = &registeredServices[i]
			break
		}
	}
	if svc == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found: " + serviceName})
		return
	}

	// SSE headers
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	send := func(line string) {
		fmt.Fprintf(c.Writer, "data: %s\n\n", line)
		c.Writer.Flush()
	}

	send("=== Starting tests for: " + svc.DisplayName + " ===")

	start := time.Now()
	result := &models.Coverage{ServiceName: svc.Name}

	// ── Backend tests ─────────────────────────────────────────────────────────
	if svc.BackendPath != "" && svc.BackendType == "go" {
		send("")
		send("▶ Running backend tests (go test)...")
		bePath := resolvePath(svc.BackendPath)
		send("  Path: " + bePath)

		_, bePassed, beFailed, beFlaky, beCov := runCommand(
			bePath, send,
			"go", "test", "./...", "-v", "-coverpkg=./...", "-cover",
		)
		result.PassedTests += bePassed
		result.FailedTests += beFailed
		result.FlakyTests += beFlaky
		result.TotalTests += bePassed + beFailed
		result.BackendCoverage = beCov
	}

	// ── Frontend tests ────────────────────────────────────────────────────────
	if svc.FrontendPath != "" && svc.FrontendType == "npm" {
		send("")
		send("▶ Running frontend tests (npm test)...")
		fePath := resolvePath(svc.FrontendPath)
		send("  Path: " + fePath)

		// Always install deps to ensure vitest is available
		send("  Installing npm dependencies...")
		runCommand(fePath, send, "npm", "install", "--silent")

		_, fePassed, feFailed, feFlaky, feCov := runCommand(
			fePath, send,
			"npm", "test", "--", "--coverage",
		)
		result.PassedTests += fePassed
		result.FailedTests += feFailed
		result.FlakyTests += feFlaky
		result.TotalTests += fePassed + feFailed
		result.FrontendCoverage = feCov
	}

	// ── Summary ───────────────────────────────────────────────────────────────
	elapsed := time.Since(start).Seconds()
	result.AvgExecutionTime = math.Round(elapsed*1000) / 1000
	result.Timestamp = time.Now()

	send("")
	send("======================================")
	send(fmt.Sprintf("  BE Coverage:  %.1f%%", result.BackendCoverage))
	send(fmt.Sprintf("  FE Coverage:  %.1f%%", result.FrontendCoverage))
	send(fmt.Sprintf("  Total Tests:  %d", result.TotalTests))
	send(fmt.Sprintf("  Passed:       %d", result.PassedTests))
	send(fmt.Sprintf("  Failed:       %d", result.FailedTests))
	send(fmt.Sprintf("  Flaky:        %d", result.FlakyTests))
	send(fmt.Sprintf("  Time:         %.2fs", elapsed))
	send("======================================")

	// ── Save to DB ────────────────────────────────────────────────────────────
	if err := database.DB.Create(result).Error; err != nil {
		send("ERROR saving to database: " + err.Error())
		fmt.Fprintf(c.Writer, "event: error\ndata: %s\n\n", err.Error())
	} else {
		send("✓ Results saved to dashboard")
		fmt.Fprintf(c.Writer, "event: done\ndata: saved\n\n")
	}
	c.Writer.Flush()
}

// runCommand runs a command in dir, streams each line via send(),
// and returns (fullOutput, passed, failed, flaky, coveragePct).
func runCommand(dir string, send func(string), name string, args ...string) (string, int, int, int, float64) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		send("ERROR: " + err.Error())
		return "", 0, 0, 0, 0
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		send("ERROR: " + err.Error())
		return "", 0, 0, 0, 0
	}

	if err := cmd.Start(); err != nil {
		send("ERROR starting command: " + err.Error())
		return "", 0, 0, 0, 0
	}

	var (
		sb     strings.Builder
		passed int
		failed int
		covPct float64
	)

	// Per-test pass/fail tracking for flaky detection
	testPassCount := map[string]int{}
	testFailCount := map[string]int{}

	// Compile regexes once outside the loop
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	reGoCoverage := regexp.MustCompile(`coverage:\s*([\d.]+)%`)
	reVitestCoverage := regexp.MustCompile(`All files\s*\|\s*([\d.]+)`)
	reGoSubtestPass := regexp.MustCompile(`^\s+--- PASS:`)
	reGoTopPass := regexp.MustCompile(`^--- PASS: [^/]+\(`)
	reGoFail := regexp.MustCompile(`--- FAIL:`)
	reVitestPass := regexp.MustCompile(`Tests\s+(\d+)\s+passed`)
	reTestName := regexp.MustCompile(`--- (PASS|FAIL): (\S+)`)

	hasSubtests := false
	reader := io.MultiReader(stdout, stderr)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		// Strip ANSI color codes
		line = ansiRegex.ReplaceAllString(line, "")

		sb.WriteString(line + "\n")
		send(line)

		// ── Coverage extraction ──────────────────────────────────────────────
		if m := reGoCoverage.FindStringSubmatch(line); m != nil {
			if v, err := strconv.ParseFloat(m[1], 64); err == nil && v > covPct {
				covPct = v
			}
		}

		if m := reVitestCoverage.FindStringSubmatch(line); m != nil {
			if v, err := strconv.ParseFloat(m[1], 64); err == nil && v > covPct {
				covPct = v
			}
		}

		// ── Go test pass/fail counting ───────────────────────────────────────
		if reGoSubtestPass.MatchString(line) {
			hasSubtests = true
			passed++
		} else if reGoTopPass.MatchString(line) && !hasSubtests {
			passed++
		}

		if reGoFail.MatchString(line) {
			failed++
		}

		// ── Vitest pass counting ─────────────────────────────────────────────
		if m := reVitestPass.FindStringSubmatch(line); m != nil {
			if v, err := strconv.Atoi(m[1]); err == nil {
				passed += v
			}
		}

		// ── Flaky detection: track per-test pass/fail counts ─────────────────
		// A test is flaky if it appears as both PASS and FAIL in the same run
		// This happens when tests are run with -count=2 or retry flags
		if m := reTestName.FindStringSubmatch(line); m != nil {
			status := m[1]
			testName := m[2]
			switch status {
			case "PASS":
				testPassCount[testName]++
			case "FAIL":
				testFailCount[testName]++
			}
		}
	}

	// Count flaky tests — appeared as both PASS and FAIL in the same run
	flaky := 0
	for name := range testFailCount {
		if testPassCount[name] > 0 {
			flaky++
		}
	}

	cmd.Wait()
	return sb.String(), passed, failed, flaky, covPct
}

// resolvePath resolves a path relative to the config directory.
func resolvePath(p string) string {
	if filepath.IsAbs(p) {
		return p
	}
	_, filename, _, _ := runtime.Caller(0)
	base := filepath.Dir(filepath.Dir(filename)) // project root
	return filepath.Clean(filepath.Join(base, p))
}

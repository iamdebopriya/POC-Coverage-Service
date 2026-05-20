#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

BACKEND_DIR="$ROOT_DIR/backend"
FRONTEND_DIR="$ROOT_DIR/frontend"

TMP_BE="/tmp/backend_test_output.txt"

echo "Inventory Tracker - Coverage Script"
echo ""


# Backend Tests


echo "Running backend tests..."
cd "$BACKEND_DIR" || exit 1

START_BE=$(date +%s)

go test -v -covermode=atomic \
  -coverpkg=inventory-tracker/backend/handlers,inventory-tracker/backend/database,inventory-tracker/backend/router \
  -coverprofile=coverage.out ./tests/... > "$TMP_BE" 2>&1

END_BE=$(date +%s)

BE_TIME=$((END_BE - START_BE))

BACKEND_COVERAGE=$(go tool cover -func=coverage.out 2>/dev/null \
  | grep "total:" \
  | awk '{print $3}' \
  | tr -d '%')

BACKEND_COVERAGE=${BACKEND_COVERAGE:-0}

TOTAL_BE=$(grep -c "^--- " "$TMP_BE")
PASSED_BE=$(grep -c "^--- PASS" "$TMP_BE")
FAILED_BE=$(grep -c "^--- FAIL" "$TMP_BE")

TOTAL_BE=${TOTAL_BE:-0}
PASSED_BE=${PASSED_BE:-0}
FAILED_BE=${FAILED_BE:-0}

echo "Backend Coverage: ${BACKEND_COVERAGE}%"
echo "Backend Tests: $TOTAL_BE total, $PASSED_BE passed, $FAILED_BE failed"
echo "Backend Time: ${BE_TIME}s"


# Frontend Tests


echo ""
echo "Running frontend tests..."

cd "$FRONTEND_DIR" || exit 1

FLAKY_TESTS=0
FRONTEND_FAILED_RUNS=0

START_FE=$(date +%s)

for i in 1 2 3
do
  echo "Frontend test run $i..."

  OUTPUT_FILE="/tmp/frontend_test_output_$i.txt"

  npm run test:coverage > "$OUTPUT_FILE" 2>&1

  FAILED_COUNT=$(grep -oE '[0-9]+ failed' "$OUTPUT_FILE" \
    | head -1 \
    | grep -oE '[0-9]+')

  FAILED_COUNT=${FAILED_COUNT:-0}

  if [ "$FAILED_COUNT" -gt 0 ]; then
    FRONTEND_FAILED_RUNS=$((FRONTEND_FAILED_RUNS + 1))
  fi
done

END_FE=$(date +%s)

TOTAL_FE_TIME=$((END_FE - START_FE))
FE_TIME=$((TOTAL_FE_TIME / 3))

# Detect flaky tests
if [ "$FRONTEND_FAILED_RUNS" -gt 0 ] && [ "$FRONTEND_FAILED_RUNS" -lt 3 ]; then
  FLAKY_TESTS=1
fi

LATEST_OUTPUT="/tmp/frontend_test_output_3.txt"

FRONTEND_COVERAGE=$(grep "All files" "$LATEST_OUTPUT" \
  | head -1 \
  | grep -oE '[0-9]+(\.[0-9]+)?' \
  | head -1)

FRONTEND_COVERAGE=${FRONTEND_COVERAGE:-0}

TEST_LINE=$(grep -E "Tests.*passed" "$LATEST_OUTPUT" | tail -1)

PASSED_FE=$(echo "$TEST_LINE" \
  | grep -oE '[0-9]+ passed' \
  | grep -oE '[0-9]+')

FAILED_FE=$(echo "$TEST_LINE" \
  | grep -oE '[0-9]+ failed' \
  | grep -oE '[0-9]+')

PASSED_FE=${PASSED_FE:-0}
FAILED_FE=${FAILED_FE:-0}

TOTAL_FE=$((PASSED_FE + FAILED_FE))

echo "Frontend Coverage: ${FRONTEND_COVERAGE}%"
echo "Frontend Tests: $TOTAL_FE total, $PASSED_FE passed, $FAILED_FE failed"
echo "Frontend Time: ${FE_TIME}s"


# Combined Results


TOTAL_TIME=$((BE_TIME + FE_TIME))

TOTAL_TESTS=$((TOTAL_BE + TOTAL_FE))
PASSED_TESTS=$((PASSED_BE + PASSED_FE))
FAILED_TESTS=$((FAILED_BE + FAILED_FE))

echo ""
echo "Total Tests    : $TOTAL_TESTS"
echo "Passed         : $PASSED_TESTS"
echo "Failed         : $FAILED_TESTS"
echo "Flaky Tests    : $FLAKY_TESTS"
echo "Backend Cov    : ${BACKEND_COVERAGE}%"
echo "Frontend Cov   : ${FRONTEND_COVERAGE}%"
echo "Avg Time       : ${TOTAL_TIME}s"


# Send Results


echo ""
echo "Sending results to dashboard..."

JSON_PAYLOAD=$(cat <<EOF
{
  "service_name": "inventory-tracker",
  "backend_coverage": $BACKEND_COVERAGE,
  "frontend_coverage": $FRONTEND_COVERAGE,
  "total_tests": $TOTAL_TESTS,
  "passed_tests": $PASSED_TESTS,
  "failed_tests": $FAILED_TESTS,
  "flaky_tests": $FLAKY_TESTS,
  "avg_execution_time": $TOTAL_TIME
}
EOF
)

echo "$JSON_PAYLOAD"

curl -s -X POST http://localhost:8080/api/coverage \
  -H "Content-Type: application/json" \
  -d "$JSON_PAYLOAD"

echo ""
echo "Done! Check dashboard at http://localhost:5173/dashboard"
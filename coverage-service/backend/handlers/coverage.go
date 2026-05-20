package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"coverage-service/backend/database"
	"coverage-service/backend/models"

	"github.com/gin-gonic/gin"
)

func SaveCoverage(c *gin.Context) {
	var coverage models.Coverage
	if err := c.ShouldBindJSON(&coverage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coverage.Timestamp = time.Now()
	database.DB.Create(&coverage)
	c.JSON(http.StatusCreated, coverage)
}

func GetCoverage(c *gin.Context) {
	var coverages []models.Coverage
	query := database.DB
	if from, to := c.Query("from"), c.Query("to"); from != "" && to != "" {
		query = query.Where("timestamp BETWEEN ? AND ?", from, to)
	}
	if s := c.Query("service"); s != "" {
		query = query.Where("service_name = ?", s)
	}
	query.Order("timestamp desc").Find(&coverages)
	c.JSON(http.StatusOK, coverages)
}

func GetRegisteredServices(c *gin.Context) {
	// Returns services from services.json
	c.JSON(http.StatusOK, registeredServices)
}

func GetCoverageServices(c *gin.Context) {
	// Returns services that already have coverage results in DB
	var services []string
	database.DB.Model(&models.Coverage{}).Distinct("service_name").Pluck("service_name", &services)
	c.JSON(http.StatusOK, services)
}

func DownloadCoverage(c *gin.Context) {
	var coverages []models.Coverage
	query := database.DB
	if from, to := c.Query("from"), c.Query("to"); from != "" && to != "" {
		query = query.Where("timestamp BETWEEN ? AND ?", from, to)
	}
	if s := c.Query("service"); s != "" {
		query = query.Where("service_name = ?", s)
	}
	query.Order("timestamp desc").Find(&coverages)

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=coverage_results.csv")
	w := csv.NewWriter(c.Writer)
	defer w.Flush()

	w.Write([]string{"ID", "Service Name", "Backend Coverage", "Frontend Coverage",
		"Total Tests", "Passed Tests", "Failed Tests", "Flaky Tests", "Avg Execution Time", "Timestamp"})

	for _, cov := range coverages {
		w.Write([]string{
			fmt.Sprintf("%d", cov.ID), cov.ServiceName,
			fmt.Sprintf("%.2f", cov.BackendCoverage),
			fmt.Sprintf("%.2f", cov.FrontendCoverage),
			fmt.Sprintf("%d", cov.TotalTests), fmt.Sprintf("%d", cov.PassedTests),
			fmt.Sprintf("%d", cov.FailedTests), fmt.Sprintf("%d", cov.FlakyTests),
			fmt.Sprintf("%.2f", cov.AvgExecutionTime),
			cov.Timestamp.Format("2006-01-02 15:04:05"),
		})
	}
}

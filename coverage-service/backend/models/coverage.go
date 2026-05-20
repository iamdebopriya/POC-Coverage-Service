package models

import "time"

type Coverage struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ServiceName      string    `json:"service_name" gorm:"not null"`
	BackendCoverage  float64   `json:"backend_coverage"`
	FrontendCoverage float64   `json:"frontend_coverage"`
	TotalTests       int       `json:"total_tests"`
	PassedTests      int       `json:"passed_tests"`
	FailedTests      int       `json:"failed_tests"`
	FlakyTests       int       `json:"flaky_tests"`
	AvgExecutionTime float64   `json:"avg_execution_time"`
	Timestamp        time.Time `json:"timestamp"`
}

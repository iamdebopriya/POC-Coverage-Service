package router

import (
	"coverage-service/backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		// Coverage CRUD
		api.POST("/coverage", handlers.SaveCoverage)
		api.GET("/coverage", handlers.GetCoverage)
		api.GET("/coverage/download", handlers.DownloadCoverage)
		api.GET("/coverage/services", handlers.GetCoverageServices)

		// registered services from services.json → populates dropdown
		api.GET("/registered-services", handlers.GetRegisteredServices)

		// run tests + stream output via SSE → saves result to DB
		api.GET("/run/:service", handlers.RunTestsSSE)
	}

	return r
}

package router

import (
	"inventory-tracker/backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-User-ID")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		// Auth routes
		api.POST("/auth/register", handlers.Register)
		api.POST("/auth/login", handlers.Login)

		// Item routes
		api.GET("/items", handlers.GetAllItems)
		api.GET("/items/low-stock", handlers.GetLowStockItems)
		api.GET("/items/:id", handlers.GetItem)
		api.POST("/items", handlers.CreateItem)
		api.PUT("/items/:id", handlers.UpdateItem)
		api.DELETE("/items/:id", handlers.DeleteItem)

		// Coverage routes
		// api.POST("/coverage", handlers.SaveCoverage)
		// api.GET("/coverage", handlers.GetCoverage)
		// api.GET("/coverage/download", handlers.DownloadCoverage)
	}

	return r
}

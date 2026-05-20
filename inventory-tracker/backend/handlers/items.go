package handlers

import (
	"net/http"
	"strconv"

	"inventory-tracker/backend/database"
	"inventory-tracker/backend/models"

	"github.com/gin-gonic/gin"
)

func getUserID(c *gin.Context) uint {
	userIDStr := c.GetHeader("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return 0
	}
	return uint(userID)
}

func GetAllItems(c *gin.Context) {
	userID := getUserID(c)
	var items []models.Item
	database.DB.Where("user_id = ?", userID).Find(&items)
	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	userID := getUserID(c)
	id := c.Param("id")
	var item models.Item
	result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&item)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.UserID = userID
	database.DB.Create(&item)
	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	userID := getUserID(c)
	id := c.Param("id")
	var item models.Item
	result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&item)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.UserID = userID
	database.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	userID := getUserID(c)
	id := c.Param("id")
	var item models.Item
	result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&item)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	database.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func GetLowStockItems(c *gin.Context) {
	userID := getUserID(c)
	var items []models.Item
	database.DB.Where("user_id = ? AND quantity < ?", userID, 10).Find(&items)
	c.JSON(http.StatusOK, items)
}

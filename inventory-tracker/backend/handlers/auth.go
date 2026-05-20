package handlers

import (
	"net/http"

	"inventory-tracker/backend/database"
	"inventory-tracker/backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		StoreName string `json:"store_name"`
		Password  string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		StoreName: input.StoreName,
		Password:  string(hashedPassword),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Store name already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         user.ID,
		"store_name": user.StoreName,
	})
}

func Login(c *gin.Context) {
	var input struct {
		StoreName string `json:"store_name"`
		Password  string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := database.DB.Where("store_name = ?", input.StoreName).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid store name or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid store name or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"store_name": user.StoreName,
	})
}

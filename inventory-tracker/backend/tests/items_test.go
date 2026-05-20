package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"inventory-tracker/backend/database"
	"inventory-tracker/backend/models"
	"inventory-tracker/backend/router"

	"github.com/joho/godotenv"
)

func setupTestDB() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	os.Setenv("DB_NAME", "inventory_tracker_test")
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Item{})
	// database.DB.AutoMigrate(&models.Coverage{})
}

func TestGetAllItems(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/api/items", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}

func TestCreateItem(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	database.DB.Where("store_name = ?", "CreateTestStore").Delete(&models.User{})
	registerBody, _ := json.Marshal(map[string]string{
		"store_name": "CreateTestStore",
		"password":   "123456",
	})
	req1, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(registerBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)

	var userResponse map[string]interface{}
	json.Unmarshal(w1.Body.Bytes(), &userResponse)
	userID := fmt.Sprintf("%.0f", userResponse["id"].(float64))

	item := map[string]interface{}{
		"name":     "Test Item",
		"category": "Test Category",
		"quantity": 5,
		"price":    99.99,
	}
	body, _ := json.Marshal(item)
	req, _ := http.NewRequest("POST", "/api/items", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User-ID", userID)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201 got %d", w.Code)
	}
}
func TestGetItem(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	item := models.Item{
		Name:     "Get Test",
		Category: "Test",
		Quantity: 3,
		Price:    10.00,
	}
	database.DB.Create(&item)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/items/%d", item.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}

func TestGetLowStockItems(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/api/items/low-stock", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}

func TestDeleteItem(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	item := models.Item{
		Name:     "Delete Test",
		Category: "Test",
		Quantity: 3,
		Price:    10.00,
	}
	database.DB.Create(&item)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/items/%d", item.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}

func TestUpdateItem(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	item := models.Item{
		Name:     "Update Test",
		Category: "Test",
		Quantity: 3,
		Price:    10.00,
	}
	database.DB.Create(&item)

	updatedItem := map[string]interface{}{
		"name":     "Updated Item",
		"category": "Updated Category",
		"quantity": 20,
		"price":    199.99,
	}

	body, _ := json.Marshal(updatedItem)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/items/%d", item.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}

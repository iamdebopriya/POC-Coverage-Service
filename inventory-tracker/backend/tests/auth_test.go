package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"inventory-tracker/backend/database"
	"inventory-tracker/backend/models"
	"inventory-tracker/backend/router"
)

func TestRegister(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	database.DB.Where("store_name = ?", "TestStore").Delete(&models.User{})

	body, _ := json.Marshal(map[string]string{
		"store_name": "TestStore",
		"password":   "123456",
	})

	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201 got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["store_name"] != "TestStore" {
		t.Errorf("Expected store_name TestStore got %v", response["store_name"])
	}
}

func TestRegisterDuplicate(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	database.DB.Where("store_name = ?", "DuplicateStore").Delete(&models.User{})

	body, _ := json.Marshal(map[string]string{
		"store_name": "DuplicateStore",
		"password":   "123456",
	})

	req1, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)

	body2, _ := json.Marshal(map[string]string{
		"store_name": "DuplicateStore",
		"password":   "123456",
	})
	req2, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body2))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 got %d", w2.Code)
	}
}

func TestLogin(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	database.DB.Where("store_name = ?", "LoginStore").Delete(&models.User{})

	registerBody, _ := json.Marshal(map[string]string{
		"store_name": "LoginStore",
		"password":   "123456",
	})
	req1, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(registerBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)

	loginBody, _ := json.Marshal(map[string]string{
		"store_name": "LoginStore",
		"password":   "123456",
	})
	req2, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(loginBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Errorf("Expected status 200 got %d", w2.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &response)

	if response["store_name"] != "LoginStore" {
		t.Errorf("Expected store_name LoginStore got %v", response["store_name"])
	}
}

func TestLoginWrongPassword(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	database.DB.Where("store_name = ?", "WrongPassStore").Delete(&models.User{})

	registerBody, _ := json.Marshal(map[string]string{
		"store_name": "WrongPassStore",
		"password":   "123456",
	})
	req1, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(registerBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)

	loginBody, _ := json.Marshal(map[string]string{
		"store_name": "WrongPassStore",
		"password":   "wrongpassword",
	})
	req2, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(loginBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401 got %d", w2.Code)
	}
}

func TestLoginWrongStoreName(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	loginBody, _ := json.Marshal(map[string]string{
		"store_name": "NonExistentStore",
		"password":   "123456",
	})
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401 got %d", w.Code)
	}
}

func TestCreateItemWithUser(t *testing.T) {
	setupTestDB()
	r := router.SetupRouter()

	database.DB.Where("store_name = ?", "ItemStore").Delete(&models.User{})

	registerBody, _ := json.Marshal(map[string]string{
		"store_name": "ItemStore",
		"password":   "123456",
	})
	req1, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(registerBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)

	var userResponse map[string]interface{}
	json.Unmarshal(w1.Body.Bytes(), &userResponse)
	userID := fmt.Sprintf("%.0f", userResponse["id"].(float64))

	itemBody, _ := json.Marshal(map[string]interface{}{
		"name":     "Test Item",
		"category": "Test",
		"quantity": 10,
		"price":    99.99,
	})
	req2, _ := http.NewRequest("POST", "/api/items", bytes.NewBuffer(itemBody))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusCreated {
		t.Errorf("Expected status 201 got %d", w2.Code)
	}
}
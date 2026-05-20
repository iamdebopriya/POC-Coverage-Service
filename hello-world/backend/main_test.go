package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Run("returns 200 and correct JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
		w := httptest.NewRecorder()

		helloHandler(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", res.StatusCode)
		}

		var body HelloResponse
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if body.Message != "Hello, World!" {
			t.Errorf("expected message 'Hello, World!', got '%s'", body.Message)
		}
		if body.Status != "ok" {
			t.Errorf("expected status 'ok', got '%s'", body.Status)
		}
	})

	t.Run("sets CORS headers", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
		w := httptest.NewRecorder()

		helloHandler(w, req)

		res := w.Result()
		if res.Header.Get("Access-Control-Allow-Origin") != "*" {
			t.Error("expected CORS header Access-Control-Allow-Origin: *")
		}
	})

	t.Run("rejects POST with 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/hello", nil)
		w := httptest.NewRecorder()

		helloHandler(w, req)

		if w.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected 405, got %d", w.Result().StatusCode)
		}
	})

	t.Run("handles OPTIONS preflight", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodOptions, "/api/hello", nil)
		w := httptest.NewRecorder()

		helloHandler(w, req)

		if w.Result().StatusCode != http.StatusNoContent {
			t.Errorf("expected 204, got %d", w.Result().StatusCode)
		}
	})
}

func TestHealthHandler(t *testing.T) {
	t.Run("returns 200 and healthy status", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
		w := httptest.NewRecorder()

		healthHandler(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", res.StatusCode)
		}

		var body map[string]string
		json.NewDecoder(res.Body).Decode(&body)
		if body["status"] != "healthy" {
			t.Errorf("expected status 'healthy', got '%s'", body["status"])
		}
	})
}

func TestNewMux(t *testing.T) {
	t.Run("routes /api/hello", func(t *testing.T) {
		mux := NewMux()
		req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Result().StatusCode)
		}
	})

	t.Run("routes /api/health", func(t *testing.T) {
		mux := NewMux()
		req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Result().StatusCode)
		}
	})
}

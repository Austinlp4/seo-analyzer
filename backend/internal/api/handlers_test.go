package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"automated-seo-analyzer/backend/internal/models"
)

func TestHandleAnalyze(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedURL    string
	}{
		{"Valid URL", "https://example.com", http.StatusOK, "https://example.com"},
		{"Invalid URL", "not-a-url", http.StatusBadRequest, ""},
		{"Empty URL", "", http.StatusBadRequest, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request body
			reqBody := models.AnalysisRequest{URL: tt.url}
			body, _ := json.Marshal(reqBody)

			// Create a request
			req, _ := http.NewRequest("POST", "/api/analyze", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Create a Gin context
			c, _ := gin.CreateTestContext(rr)
			c.Request = req

			// Call the handler
			handleAnalyze(c)

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Check the response body
			if tt.expectedStatus == http.StatusOK {
				var result models.AnalysisResult
				json.Unmarshal(rr.Body.Bytes(), &result)
				if result.URL != tt.expectedURL {
					t.Errorf("handler returned unexpected URL: got %v want %v", result.URL, tt.expectedURL)
				}
			}
		})
	}
}

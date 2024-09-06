package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Austinlp4/seo-analyzer/backend/internal/models"
)

func TestHandleAnalyze(t *testing.T) {
	// Create a request body
	reqBody := models.AnalysisRequest{URL: "https://example.com"}
	body, _ := json.Marshal(reqBody)

	// Create a request
	req, _ := http.NewRequest("POST", "/api/analyze", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handleAnalyze(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var result models.AnalysisResult
	json.Unmarshal(rr.Body.Bytes(), &result)
	if result.URL != reqBody.URL {
		t.Errorf("handler returned unexpected body: got %v want %v", result.URL, reqBody.URL)
	}
}

package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Austinlp4/seo-analyzer/backend/internal/models"
	"github.com/Austinlp4/seo-analyzer/backend/internal/seo"
)

func handleAnalyze(w http.ResponseWriter, r *http.Request) {
	var req models.AnalysisRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := seo.Analyze(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func HandleStaticFiles(w http.ResponseWriter, r *http.Request) {
	// Serve static files from the React build directory
	fs := http.FileServer(http.Dir("./frontend/build"))

	// If the file exists, serve it directly
	if _, err := os.Stat("./frontend/build" + r.URL.Path); err == nil {
		fs.ServeHTTP(w, r)
		return
	}

	// Otherwise, serve the index.html file
	http.ServeFile(w, r, "./frontend/build/index.html")
}

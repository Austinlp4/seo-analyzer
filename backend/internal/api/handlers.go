package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Austinlp4/seo-analyzer/backend/internal/auth"
	"github.com/Austinlp4/seo-analyzer/backend/internal/cache"
	"github.com/Austinlp4/seo-analyzer/backend/internal/database"
	"github.com/Austinlp4/seo-analyzer/backend/internal/models"
	"github.com/Austinlp4/seo-analyzer/backend/internal/seo"
)

func handleAnalyze(w http.ResponseWriter, r *http.Request) {
	var req models.AnalysisRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Validate URL format
	_, err = url.ParseRequestURI(req.URL)
	if err != nil {
		http.Error(w, "Invalid URL format: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check cache
	cache := cache.GetCache()
	if cachedResult, found := cache.Get(req.URL); found {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cachedResult)
		return
	}

	result, err := seo.Analyze(req.URL)
	if err != nil {
		switch err.(type) {
		case *url.Error:
			http.Error(w, "Error accessing URL: "+err.Error(), http.StatusBadRequest)
		case *seo.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			log.Printf("Error analyzing URL: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Convert AnalysisResponse to AnalysisResult
	analysisResult := &models.AnalysisResult{
		URL:      req.URL,
		Analysis: result,
	}

	// Store the analysis result in the database
	err = database.StoreAnalysisResult(analysisResult)
	if err != nil {
		log.Printf("Error storing analysis result: %v", err)
		// Continue execution, as this error shouldn't prevent sending the result to the client
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

	http.ServeFile(w, r, "./frontend/build/index.html")
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = auth.CreateUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := auth.GetUserByUsername(req.Username)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"automated-seo-analyzer/backend/internal/auth"
	"automated-seo-analyzer/backend/internal/database"
	"automated-seo-analyzer/backend/internal/models"
	"automated-seo-analyzer/backend/internal/seo"
)

func handleAnalyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// Get user ID from the JWT token
	userID, err := auth.GetUserIDFromToken(r)
	if err != nil {
		log.Printf("Error getting user ID from token: %v", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Convert AnalysisResponse to AnalysisResult
	analysisResult := &models.AnalysisResult{
		URL:      req.URL,
		Analysis: result,
	}

	// Store the analysis result in the database
	err = database.StoreAnalysisResult(userID, analysisResult)
	if err != nil {
		log.Printf("Error storing analysis result: %v", err)
		// Continue execution, as this error shouldn't prevent sending the result to the client
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func HandleStaticFiles(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./frontend/build"))

	if _, err := os.Stat("./frontend/build" + r.URL.Path); err == nil {
		fs.ServeHTTP(w, r)
		return
	}

	http.ServeFile(w, r, "./frontend/build/index.html")
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func handleGetUserAnalyses(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering handleGetUserAnalyses function")

	// Get user ID from the JWT token
	userID, err := auth.GetUserIDFromToken(r)
	if err != nil {
		log.Printf("Error getting user ID from token: %v", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	log.Printf("User ID retrieved: %d", userID)

	analyses, err := database.GetUserAnalyses(userID)
	if err != nil {
		log.Printf("Error fetching user analyses: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("Retrieved %d analyses for user", len(analyses))

	// Set headers to prevent caching
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	log.Println("Sending JSON response")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analyses)
}

func handleCollectSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var sessionData models.SessionData
	err := json.NewDecoder(r.Body).Decode(&sessionData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = database.StoreSessionData(&sessionData)
	if err != nil {
		log.Printf("Error storing session data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleCreateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := auth.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	project, err := database.CreateProject(userID, req.Name, req.URL)
	if err != nil {
		log.Printf("Error creating project: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func handleGetUserProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projects, err := database.GetUserProjects(userID)
	if err != nil {
		log.Printf("Error getting user projects: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

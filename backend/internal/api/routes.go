package api

import (
	"log"
	"net/http"
	"strings"
)

func SetupRoutes(mux *http.ServeMux) {
	log.Println("Setting up routes")

	// API routes
	mux.HandleFunc("/api/analyze", handleAnalyze)
	mux.HandleFunc("/api/register", handleRegister)
	mux.HandleFunc("/api/login", handleLogin)
	mux.HandleFunc("/api/user/analyses", authMiddleware(handleGetUserAnalyses))
	mux.HandleFunc("/api/collect-session", handleCollectSession)
	mux.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			authMiddleware(handleCreateProject)(w, r)
		case http.MethodGet:
			authMiddleware(handleGetUserProjects)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Serve static files
	fs := http.FileServer(http.Dir("./frontend/build"))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		fs.ServeHTTP(w, r)
	}))

	log.Println("Routes setup completed")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "frontend/public/index.html")
		return
	}

	if strings.HasPrefix(r.URL.Path, "/public/") {
		file := filepath.Join("frontend", r.URL.Path)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, file)
		return
	}

	// Serve index.html for other routes (client-side routing)
	http.ServeFile(w, r, "frontend/public/index.html")
}

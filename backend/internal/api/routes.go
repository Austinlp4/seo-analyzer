package api

import "net/http"

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api/analyze", handleAnalyze)
}

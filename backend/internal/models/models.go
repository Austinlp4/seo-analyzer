package models

type AnalysisRequest struct {
	URL string `json:"url"`
}

type AnalysisResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StatusCode  int    `json:"statusCode"`
}

type AnalysisResult struct {
	URL string
	// Add other fields as needed
}

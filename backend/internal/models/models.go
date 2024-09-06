package models

type AnalysisRequest struct {
	URL string `json:"url"`
}

type AnalysisResponse struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StatusCode  int    `json:"statusCode"`
	H1Count     int    `json:"h1Count"`
	WordCount   int    `json:"wordCount"`
}

type AnalysisResult struct {
	URL string
	// Add other fields as needed
}

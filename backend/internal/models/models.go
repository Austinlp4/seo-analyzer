package models

type AnalysisRequest struct {
	URL string `json:"url"`
}

type AnalysisResponse struct {
	URL               string  `json:"url"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	StatusCode        int     `json:"statusCode"`
	H1Count           int     `json:"h1Count"`
	WordCount         int     `json:"wordCount"`
	PageLoadSpeed     float64 `json:"pageLoadSpeed"`
	MobileFriendly    bool    `json:"mobileFriendly"`
	ResponsiveDesign  bool    `json:"responsiveDesign"`
	SSLCertificate    bool    `json:"sslCertificate"`
	MetaRobotsContent string  `json:"metaRobotsContent"`
	SEOScore          int     `json:"seoScore"`
}

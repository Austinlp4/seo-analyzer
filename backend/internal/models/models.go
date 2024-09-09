package models

import "time"

type AnalysisRequest struct {
	URL string `json:"url"`
}

type AnalysisResponse struct {
	URL                  string  `json:"url"`
	Title                string  `json:"title"`
	Description          string  `json:"description"`
	StatusCode           int     `json:"statusCode"`
	H1Count              int     `json:"h1Count"`
	WordCount            int     `json:"wordCount"`
	PageLoadSpeed        float64 `json:"pageLoadSpeed"`
	MobileFriendly       bool    `json:"mobileFriendly"`
	ResponsiveDesign     bool    `json:"responsiveDesign"`
	SSLCertificate       bool    `json:"sslCertificate"`
	MetaRobotsContent    string  `json:"metaRobotsContent"`
	SEOScore             int     `json:"seoScore"`
	PageSpeedScore       float64 `json:"pageSpeedScore"`
	FirstContentfulPaint float64 `json:"firstContentfulPaint"`
	TimeToInteractive    float64 `json:"timeToInteractive"`
}

type SessionData struct {
	SessionID    string    `json:"sessionId"`
	UserID       string    `json:"userId"`
	ProjectID    string    `json:"projectId"`
	URL          string    `json:"url"`
	Duration     int64     `json:"duration"`
	ScreenWidth  int       `json:"screenWidth"`
	ScreenHeight int       `json:"screenHeight"`
	UserAgent    string    `json:"userAgent"`
	Timestamp    time.Time `json:"timestamp"`
}

type Project struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
}

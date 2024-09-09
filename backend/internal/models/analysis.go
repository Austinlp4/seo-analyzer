package models

import (
	"time"
)

type AnalysisResult struct {
	URL                  string
	Title                string
	Description          string
	StatusCode           int
	H1Count              int
	WordCount            int
	PageLoadSpeed        float64
	MobileFriendly       bool
	ResponsiveDesign     bool
	SSLCertificate       bool
	MetaRobotsContent    string
	SEOScore             float64
	Analysis             *AnalysisResponse
	PageSpeedScore       float64 // Add this line
	CreatedAt            time.Time
	FirstContentfulPaint float64
	TimeToInteractive    float64
}

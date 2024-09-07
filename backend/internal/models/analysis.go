package models

type AnalysisResult struct {
	URL               string
	Title             string
	Description       string
	StatusCode        int
	H1Count           int
	WordCount         int
	PageLoadSpeed     float64
	MobileFriendly    bool
	ResponsiveDesign  bool
	SSLCertificate    bool
	MetaRobotsContent string
	SEOScore          float64
	Analysis          *AnalysisResponse
}

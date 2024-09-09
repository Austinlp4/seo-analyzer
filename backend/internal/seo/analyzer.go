package seo

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"encoding/json"
	"io/ioutil"
	"os"

	"automated-seo-analyzer/backend/internal/models"

	"golang.org/x/net/html"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func Analyze(urlString string) (*models.AnalysisResponse, error) {
	parsedURL, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, &ValidationError{Message: "Invalid URL format"}
	}

	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	start := time.Now()
	resp, err := client.Get(urlString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &ValidationError{Message: fmt.Sprintf("HTTP error: %d", resp.StatusCode)}
	}

	pageLoadSpeed := time.Since(start).Seconds()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	pageSpeedData, err := getPageSpeedInsights(urlString)
	if err != nil {
		log.Printf("Error getting PageSpeed Insights: %v", err)
	}

	result := &models.AnalysisResponse{
		URL:                  urlString,
		Title:                extractTitle(doc),
		Description:          extractMetaDescription(doc),
		StatusCode:           resp.StatusCode,
		H1Count:              countH1Tags(doc),
		WordCount:            countWords(doc),
		PageLoadSpeed:        pageLoadSpeed,
		MobileFriendly:       checkMobileFriendliness(doc),
		ResponsiveDesign:     checkResponsiveDesign(doc),
		SSLCertificate:       parsedURL.Scheme == "https",
		MetaRobotsContent:    extractMetaRobots(doc),
		PageSpeedScore:       getPageSpeedScore(pageSpeedData),
		FirstContentfulPaint: getFirstContentfulPaint(pageSpeedData),
		TimeToInteractive:    getTimeToInteractive(pageSpeedData),
	}

	result.SEOScore = calculateSEOScore(result)

	return result, nil
}

func getPageSpeedInsights(url string) (map[string]interface{}, error) {
	apiKey := os.Getenv("PAGESPEED_API_KEY")
	apiURL := fmt.Sprintf("https://www.googleapis.com/pagespeedonline/v5/runPagespeed?url=%s&key=%s", url, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getPageSpeedScore(data map[string]interface{}) float64 {
	if data == nil {
		return 0
	}
	if lighthouseResult, ok := data["lighthouseResult"].(map[string]interface{}); ok {
		if categories, ok := lighthouseResult["categories"].(map[string]interface{}); ok {
			if performance, ok := categories["performance"].(map[string]interface{}); ok {
				if score, ok := performance["score"].(float64); ok {
					return score * 100
				}
			}
		}
	}
	return 0
}

func getFirstContentfulPaint(data map[string]interface{}) float64 {
	return getAuditScore(data, "first-contentful-paint")
}

func getTimeToInteractive(data map[string]interface{}) float64 {
	return getAuditScore(data, "interactive")
}

func getAuditScore(data map[string]interface{}, auditName string) float64 {
	if data == nil {
		return 0
	}
	if lighthouseResult, ok := data["lighthouseResult"].(map[string]interface{}); ok {
		if audits, ok := lighthouseResult["audits"].(map[string]interface{}); ok {
			if audit, ok := audits[auditName].(map[string]interface{}); ok {
				if numericValue, ok := audit["numericValue"].(float64); ok {
					return numericValue / 1000 // Convert ms to seconds
				}
			}
		}
	}
	return 0
}

func calculateSEOScore(result *models.AnalysisResponse) int {
	score := 0

	// Title length (ideal: 50-60 characters)
	titleLength := len(result.Title)
	if titleLength >= 50 && titleLength <= 60 {
		score += 10
	} else if titleLength > 0 {
		score += 5
	}

	// Description length (ideal: 150-160 characters)
	descLength := len(result.Description)
	if descLength >= 150 && descLength <= 160 {
		score += 10
	} else if descLength > 0 {
		score += 5
	}

	// H1 count (ideal: 1)
	if result.H1Count == 1 {
		score += 10
	} else if result.H1Count > 1 {
		score += 5
	}

	// Word count (minimum: 300)
	if result.WordCount >= 300 {
		score += 10
	}

	// Page load speed (ideal: < 3 seconds)
	if result.PageLoadSpeed < 3 {
		score += 20
	} else if result.PageLoadSpeed < 5 {
		score += 10
	}

	// Mobile-friendly
	if result.MobileFriendly {
		score += 15
	}

	// Responsive design
	if result.ResponsiveDesign {
		score += 15
	}

	// SSL certificate
	if result.SSLCertificate {
		score += 10
	}

	// PageSpeed score (0-100)
	score += int(result.PageSpeedScore / 2)

	// First Contentful Paint (ideal: < 1.8 seconds)
	if result.FirstContentfulPaint < 1.8 {
		score += 10
	} else if result.FirstContentfulPaint < 3 {
		score += 5
	}

	// Time to Interactive (ideal: < 3.8 seconds)
	if result.TimeToInteractive < 3.8 {
		score += 10
	} else if result.TimeToInteractive < 7.3 {
		score += 5
	}

	return score
}

func checkMobileFriendliness(n *html.Node) bool {
	viewport := extractMetaContent(n, "viewport")
	return strings.Contains(viewport, "width=device-width")
}

func checkResponsiveDesign(n *html.Node) bool {
	mediaQueries := extractStyleContent(n)
	return strings.Contains(mediaQueries, "@media")
}

func extractMetaRobots(n *html.Node) string {
	return extractMetaContent(n, "robots")
}

func extractMetaContent(n *html.Node, name string) string {
	if n.Type == html.ElementNode && n.Data == "meta" {
		var metaName, content string
		for _, a := range n.Attr {
			if a.Key == "name" && a.Val == name {
				metaName = a.Val
			}
			if a.Key == "content" {
				content = a.Val
			}
		}
		if metaName == name {
			return content
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := extractMetaContent(c, name); result != "" {
			return result
		}
	}
	return ""
}

func extractStyleContent(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "style" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := extractStyleContent(c); result != "" {
			return result
		}
	}
	return ""
}

func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title := extractTitle(c); title != "" {
			return title
		}
	}
	return ""
}

func extractMetaDescription(n *html.Node) string {
	return extractMetaContent(n, "description")
}

func countH1Tags(n *html.Node) int {
	count := 0
	if n.Type == html.ElementNode && n.Data == "h1" {
		count++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count += countH1Tags(c)
	}
	return count
}

func countWords(n *html.Node) int {
	if n.Type == html.TextNode {
		return len(strings.Fields(n.Data))
	}
	count := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count += countWords(c)
	}
	return count
}

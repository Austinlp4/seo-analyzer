package seo

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Austinlp4/seo-analyzer/backend/internal/cache"
	"github.com/Austinlp4/seo-analyzer/backend/internal/models"
	"golang.org/x/net/html"
)

func Analyze(urlString string) (*models.AnalysisResponse, error) {
	cache := cache.GetCache()

	// Check if the result is in the cache
	if cachedResult, found := cache.Get(urlString); found {
		return cachedResult, nil
	}

	parsedURL, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
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

	pageLoadSpeed := time.Since(start).Seconds()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &models.AnalysisResponse{
		URL:               urlString,
		Title:             extractTitle(doc),
		Description:       extractMetaDescription(doc),
		StatusCode:        resp.StatusCode,
		H1Count:           countH1Tags(doc),
		WordCount:         countWords(doc),
		PageLoadSpeed:     pageLoadSpeed,
		MobileFriendly:    checkMobileFriendliness(doc),
		ResponsiveDesign:  checkResponsiveDesign(doc),
		SSLCertificate:    parsedURL.Scheme == "https",
		MetaRobotsContent: extractMetaRobots(doc),
	}

	// Store the result in the cache
	cache.Set(urlString, result)

	return result, nil
}

// Existing functions: extractTitle, extractMetaDescription, countH1Tags, countWords

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

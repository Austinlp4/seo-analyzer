package seo

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/Austinlp4/seo-analyzer/backend/internal/models"
	"golang.org/x/net/html"
)

func Analyze(urlString string) (*models.AnalysisResponse, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(urlString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	title := extractTitle(doc)
	description := extractMetaDescription(doc)
	h1Count := countH1Tags(doc)
	wordCount := countWords(doc)

	return &models.AnalysisResponse{
		URL:         urlString,
		Title:       title,
		Description: description,
		StatusCode:  resp.StatusCode,
		H1Count:     h1Count,
		WordCount:   wordCount,
	}, nil
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
	if n.Type == html.ElementNode && n.Data == "meta" {
		for _, a := range n.Attr {
			if a.Key == "name" && a.Val == "description" {
				for _, a := range n.Attr {
					if a.Key == "content" {
						return a.Val
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if desc := extractMetaDescription(c); desc != "" {
			return desc
		}
	}
	return ""
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

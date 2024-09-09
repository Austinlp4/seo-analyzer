package database

import (
	"database/sql"
	"fmt"

	"automated-seo-analyzer/backend/internal/models"
)

func GetUserAnalyses(userID int) ([]models.AnalysisResult, error) {
	rows, err := DB.Query(`
        SELECT url, title, description, status_code, h1_count, word_count,
               page_load_speed, mobile_friendly, responsive_design,
               ssl_certificate, meta_robots_content, seo_score,
               page_speed_score, first_contentful_paint, time_to_interactive,
               created_at
        FROM analysis_results
        WHERE user_id = $1
        ORDER BY created_at DESC
        LIMIT 10
    `, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query user analyses: %v", err)
	}
	defer rows.Close()

	var analyses []models.AnalysisResult
	for rows.Next() {
		var analysis models.AnalysisResult
		var pageSpeedScore, firstContentfulPaint, timeToInteractive sql.NullFloat64
		err := rows.Scan(
			&analysis.URL, &analysis.Title, &analysis.Description,
			&analysis.StatusCode, &analysis.H1Count, &analysis.WordCount,
			&analysis.PageLoadSpeed, &analysis.MobileFriendly, &analysis.ResponsiveDesign,
			&analysis.SSLCertificate, &analysis.MetaRobotsContent, &analysis.SEOScore,
			&pageSpeedScore, &firstContentfulPaint, &timeToInteractive,
			&analysis.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan analysis result: %v", err)
		}

		analysis.PageSpeedScore = nullFloatToFloat64(pageSpeedScore)
		analysis.FirstContentfulPaint = nullFloatToFloat64(firstContentfulPaint)
		analysis.TimeToInteractive = nullFloatToFloat64(timeToInteractive)

		analyses = append(analyses, analysis)
	}

	return analyses, nil
}

func nullFloatToFloat64(nf sql.NullFloat64) float64 {
	if nf.Valid {
		return nf.Float64
	}
	return 0
}

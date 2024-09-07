package database

import (
	"database/sql"
	"log"

	"github.com/Austinlp4/seo-analyzer/backend/internal/models"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "postgresql://seo-analytics_owner:gDGSJ1lX2yph@ep-holy-block-a5e3lzo0.us-east-2.aws.neon.tech/seo-analytics?sslmode=require"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Successfully connected to the database")

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`

	createAnalysisResultsTable := `
	CREATE TABLE IF NOT EXISTS analysis_results (
		id SERIAL PRIMARY KEY,
		url TEXT NOT NULL,
		title TEXT,
		description TEXT,
		status_code INT,
		h1_count INT,
		word_count INT,
		page_load_speed FLOAT,
		mobile_friendly BOOLEAN,
		responsive_design BOOLEAN,
		ssl_certificate BOOLEAN,
		meta_robots_content TEXT,
		seo_score FLOAT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	_, err = DB.Exec(createAnalysisResultsTable)
	if err != nil {
		log.Fatalf("Error creating analysis_results table: %v", err)
	}

	log.Println("Tables created successfully")
}

func StoreAnalysisResult(result *models.AnalysisResult) error {
	query := `
	INSERT INTO analysis_results (
		url, title, description, status_code, h1_count, word_count,
		page_load_speed, mobile_friendly, responsive_design,
		ssl_certificate, meta_robots_content, seo_score
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := DB.Exec(query,
		result.URL,
		result.Title,
		result.Description,
		result.StatusCode,
		result.H1Count,
		result.WordCount,
		result.PageLoadSpeed,
		result.MobileFriendly,
		result.ResponsiveDesign,
		result.SSLCertificate,
		result.MetaRobotsContent,
		result.SEOScore,
	)

	if err != nil {
		log.Printf("Error storing analysis result: %v", err)
		return err
	}

	log.Printf("Successfully stored analysis result for URL: %s", result.URL)
	return nil
}

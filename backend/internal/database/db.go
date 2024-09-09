package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"automated-seo-analyzer/backend/internal/models"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := os.Getenv("REACT_APP_DB_CONNECTION_STRING")
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
	RunMigrations()
	return DB, err
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
		user_id INT NOT NULL,
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
		page_speed_score FLOAT,
		first_contentful_paint FLOAT,
		time_to_interactive FLOAT,
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

func StoreAnalysisResult(userID int, result *models.AnalysisResult) error {
	query := `
	INSERT INTO analysis_results (
		user_id, url, title, description, status_code, h1_count, word_count,
		page_load_speed, mobile_friendly, responsive_design,
		ssl_certificate, meta_robots_content, seo_score, page_speed_score,
		first_contentful_paint, time_to_interactive
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`

	_, err := DB.Exec(query,
		userID,
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
		result.PageSpeedScore,
		result.FirstContentfulPaint,
		result.TimeToInteractive,
	)

	if err != nil {
		log.Printf("Error storing analysis result: %v", err)
		return err
	}

	return nil
}

func StoreSessionData(data *models.SessionData) error {
	query := `
	INSERT INTO session_data (session_id, user_id, project_id, url, duration, screen_width, screen_height, user_agent, timestamp)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := DB.Exec(query,
		data.SessionID,
		data.UserID,
		data.ProjectID,
		data.URL,
		data.Duration,
		data.ScreenWidth,
		data.ScreenHeight,
		data.UserAgent,
		time.Now(),
	)
	return err
}

func CreateProject(userID int, name string, url string) (*models.Project, error) {
	query := `
    INSERT INTO projects (user_id, name, url)
    VALUES ($1, $2, $3)
    RETURNING id, created_at
    `
	var project models.Project
	err := DB.QueryRow(query, userID, name, url).Scan(&project.ID, &project.CreatedAt)
	if err != nil {
		return nil, err
	}
	project.UserID = userID
	project.Name = name
	project.URL = url
	return &project, nil
}

func GetUserProjects(userID int) ([]models.Project, error) {
	query := `
    SELECT id, name, url, created_at
    FROM projects
    WHERE user_id = $1
    ORDER BY created_at DESC
    `
	rows, err := DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err := rows.Scan(&p.ID, &p.Name, &p.URL, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		p.UserID = userID
		projects = append(projects, p)
	}
	return projects, nil
}

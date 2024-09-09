package database

import (
	"log"
)

func RunMigrations() {
	addMissingColumns()
	addUserIDColumn()
	createSessionDataTable()
	createProjectsTable()
}

func addMissingColumns() {
	_, err := DB.Exec(`
	ALTER TABLE analysis_results
	ADD COLUMN IF NOT EXISTS first_contentful_paint FLOAT,
	ADD COLUMN IF NOT EXISTS time_to_interactive FLOAT;
	`)
	if err != nil {
		log.Fatalf("Error adding missing columns: %v", err)
	}
	log.Println("Successfully added missing columns to analysis_results table")
}

func addUserIDColumn() {
	_, err := DB.Exec(`
	ALTER TABLE analysis_results
	ADD COLUMN IF NOT EXISTS user_id INT NOT NULL DEFAULT 1;
	`)
	if err != nil {
		log.Fatalf("Error adding user_id column: %v", err)
	}
	log.Println("Successfully added user_id column to analysis_results table")
}

func createSessionDataTable() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS session_data (
		id SERIAL PRIMARY KEY,
		session_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		project_id TEXT NOT NULL,
		url TEXT NOT NULL,
		duration BIGINT NOT NULL,
		screen_width INT NOT NULL,
		screen_height INT NOT NULL,
		user_agent TEXT NOT NULL,
		timestamp TIMESTAMP NOT NULL
	);`)
	if err != nil {
		log.Fatalf("Error creating session_data table: %v", err)
	}
	log.Println("Successfully created session_data table")
}

func createProjectsTable() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS projects (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		name TEXT NOT NULL,
		url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`)
	if err != nil {
		log.Fatalf("Error creating projects table: %v", err)
	}
	log.Println("Successfully created projects table")
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	fmt.Println("Connected to PostgreSQL")
}

func InitTables() {
	seriesTable := `
	CREATE TABLE IF NOT EXISTS series (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		current_episode INTEGER NOT NULL,
		total_episodes INTEGER NOT NULL,
		image_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	ratingsTable := `
	CREATE TABLE IF NOT EXISTS ratings (
		id SERIAL PRIMARY KEY,
		series_id INTEGER REFERENCES series(id) ON DELETE CASCADE,
		rating INTEGER CHECK (rating >= 1 AND rating <= 5)
	);`

	_, err := DB.Exec(seriesTable)
	if err != nil {
		log.Fatal("Error creating series table:", err)
	}

	_, err = DB.Exec(ratingsTable)
	if err != nil {
		log.Fatal("Error creating ratings table:", err)
	}

	fmt.Println("Tables initialized")
}

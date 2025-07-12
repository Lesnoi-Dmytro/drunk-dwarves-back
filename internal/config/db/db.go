package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("No database connection url defined in environment")
	}

	var err error
	DB, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	migrationsPath := filepath.Join("internal", "config", "db", "migrations")
	if err := goose.Up(DB, migrationsPath); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	return DB
}

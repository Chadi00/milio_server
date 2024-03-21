package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectDb() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	tursoUrl := os.Getenv("TURSO_SQL_URL")
	if tursoUrl == "" {
		log.Fatal("Turso URL not set in .env file")
	}
	tursoAuthToken := os.Getenv("TURSO_SQL_AUTH_TOKEN")
	if tursoAuthToken == "" {
		log.Fatal("Turso Auth Token not set in .env file")
	}

	url := tursoUrl + "?authToken=" + tursoAuthToken

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to open db %s: %s", url, err)
	}
	return db, err
}

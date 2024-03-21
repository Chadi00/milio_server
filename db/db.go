package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectDb() (*sql.DB, error) {
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

func AddUsersTable() {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	createUsersTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(createUsersTableSQL); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	} else {
		fmt.Println("Users table created successfully")
	}
}

func AddErrorTable() {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	createErrorTableSQL := `
	CREATE TABLE IF NOT EXISTS errors (
		id SERIAL PRIMARY KEY,
		user_chat TEXT,
		system_chat TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		type_error VARCHAR(255),
		device VARCHAR(255)
	);
	`

	if _, err := db.Exec(createErrorTableSQL); err != nil {
		log.Fatalf("Failed to create error table: %v", err)
	} else {
		fmt.Println("Error table created successfully")
	}
}

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectDb() {

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
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()
}

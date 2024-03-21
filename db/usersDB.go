package db

import (
	"fmt"
	"log"
)

func AddUsersTable() {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	createUsersTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL,
        created_at TEXT DEFAULT CURRENT_TIMESTAMP
    );
    `

	if _, err := db.Exec(createUsersTableSQL); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	} else {
		fmt.Println("Users table created successfully")
	}
}

func DeleteUsersTable() {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	dropTableSQL := `DROP TABLE IF EXISTS users;`

	if _, err := db.Exec(dropTableSQL); err != nil {
		log.Fatalf("Failed to delete errors table: %v", err)
	} else {
		log.Println("Errors table deleted successfully")
	}
}

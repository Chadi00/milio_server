package db

import (
	"database/sql"
	"log"
	"milio/models"
	"time"
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
	}
}

func AddUser(email, password string) error {
	db, err := ConnectDb()
	if err != nil {
		log.Printf("Could not connect to database: %v", err)
		return err
	}
	defer db.Close()

	addUserSQL := `INSERT INTO users(email, password, created_at) VALUES (?, ?, CURRENT_TIMESTAMP);`

	statement, err := db.Prepare(addUserSQL)
	if err != nil {
		log.Printf("Could not prepare SQL statement: %v", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(email, password)
	if err != nil {
		log.Printf("Could not insert user into database: %v", err)
		return err
	}

	//fmt.Println("Added user successfully")
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}
	defer db.Close()

	query := `SELECT id, email, password, created_at FROM users WHERE email = ?`
	var user models.User
	var createdAtString string

	err = db.QueryRow(query, email).Scan(&user.UserID, &user.Email, &user.Password, &createdAtString)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			log.Printf("Failed to query user by email: %v", err)
			return nil, err
		}
	}

	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
	if err != nil {
		log.Printf("Failed to parse created_at timestamp: %v", err)
		return nil, err
	}

	return &user, nil
}

func DeleteUserByEmail(email string) error {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return err
	}
	defer db.Close()

	deleteQuery := `DELETE FROM users WHERE email = ?`
	stmt, err := db.Prepare(deleteQuery)
	if err != nil {
		log.Printf("Could not prepare delete statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email)
	if err != nil {
		log.Printf("Failed to delete user by email: %v", err)
		return err
	}

	return nil
}

func DeleteUsersTable() {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	dropTableSQL := `DROP TABLE IF EXISTS users;`

	if _, err := db.Exec(dropTableSQL); err != nil {
		log.Fatalf("Failed to delete users table: %v", err)
	}
}

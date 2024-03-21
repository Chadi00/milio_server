package db

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type ErrorRecord struct {
	ID         int       `json:"id"`
	UserChat   string    `json:"user_chat"`
	SystemChat string    `json:"system_chat"`
	CreatedAt  time.Time `json:"created_at"`
	TypeError  string    `json:"type_error"`
	Device     string    `json:"device"`
}

func AddErrorTable() {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	createErrorTableSQL := `
    CREATE TABLE IF NOT EXISTS errors (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_chat TEXT,
        system_chat TEXT,
        created_at TEXT DEFAULT CURRENT_TIMESTAMP,
        type_error TEXT,
        device TEXT
    );
    `

	if _, err := db.Exec(createErrorTableSQL); err != nil {
		log.Fatalf("Failed to create error table: %v", err)
	} else {
		fmt.Println("Error table created successfully")
	}
}

func AddError(userChat, systemChat, typeError, device string) {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	insertErrorSQL := `
	INSERT INTO errors (user_chat, system_chat, created_at, type_error, device)
	VALUES (?, ?, ?, ?, ?);
	`

	_, err = db.Exec(insertErrorSQL, userChat, systemChat, time.Now(), typeError, device)
	if err != nil {
		log.Fatalf("Failed to insert error: %v", err)
	} else {
		fmt.Println("Error inserted successfully")
	}
}

func ReadAllErrorsAsString() (string, error) {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	query := `SELECT id, user_chat, system_chat, created_at, type_error, device FROM errors;`

	rows, err := db.Query(query)
	if err != nil {
		return "", fmt.Errorf("failed to query errors: %v", err)
	}
	defer rows.Close()

	var result strings.Builder

	for rows.Next() {
		var id int
		var userChat, systemChat, typeError, device, createdAtStr string
		if err := rows.Scan(&id, &userChat, &systemChat, &createdAtStr, &typeError, &device); err != nil {
			return "", fmt.Errorf("failed to scan error row: %v", err)
		}

		// Parse the createdAtStr with timezone offset
		createdAt, err := time.Parse("2006-01-02 15:04:05.999999-07:00", createdAtStr)
		if err != nil {
			// If parsing fails, try without the timezone offset
			createdAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
			if err != nil {
				return "", fmt.Errorf("failed to parse created_at datetime: %v", err)
			}
		}

		line := fmt.Sprintf("ID: %d, User Chat: %s, System Chat: %s, Created At: %s, Type Error: %s, Device: %s",
			id, userChat, systemChat, createdAt.Format(time.RFC3339), typeError, device)
		result.WriteString(line + "\n")
	}

	if err = rows.Err(); err != nil {
		return "", fmt.Errorf("error iterating rows: %v", err)
	}

	return result.String(), nil
}

func DeleteErrorsTable() {
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	dropTableSQL := `DROP TABLE IF EXISTS errors;`

	if _, err := db.Exec(dropTableSQL); err != nil {
		log.Fatalf("Failed to delete errors table: %v", err)
	} else {
		log.Println("Errors table deleted successfully")
	}
}

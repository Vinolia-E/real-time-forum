package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"real_time_forum/backend/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		// fmt.Println("Error opening database:", err)
		return fmt.Errorf("Error opening database: %v", err)
	}
	// defer db.Close()

	handlers.DB = db
	//  Read schema file.
	schema, err := os.ReadFile("../backend/sqlite/schema.sql")
	if err != nil {
		return fmt.Errorf("Error reading the database schema: %v", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("Error executing schema: %w", err)
	}

	fmt.Println("Database initialized successfully.")
	return nil
}

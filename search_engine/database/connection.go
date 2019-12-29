package database

import (
	"os"
	"fmt"
	"errors"
	"database/sql"
)

// ConnectDB ...
func ConnectDB() (*sql.DB, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host     := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")

	dbConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, database)
	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		return nil, errors.New("could not open database")
	}

	return db, nil
}

package datastore

import (
    "os"
    "fmt"
    "database/sql"
)

// connectDB は環境変数を参照してデータベースに接続し、DBハンドラを返す。
func connectDB() (*sql.DB, error) {
    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host     := os.Getenv("DB_HOST")
    database := os.Getenv("DB_DATABASE")

    dbConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, database)
    db, err := sql.Open("mysql", dbConfig)
    if err != nil {
        return nil, fmt.Errorf("could not open database '%s'", dbConfig)
    }

    return db, nil
}
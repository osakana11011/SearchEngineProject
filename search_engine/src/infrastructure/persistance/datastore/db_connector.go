package datastore

import (
    "os"
    "fmt"
    "github.com/jinzhu/gorm"
)

// NewGormDBConnection は環境変数の設定を参照して、DBコネクションを確立してハンドラを返す。
func NewGormDBConnection() (*gorm.DB, error) {
    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host     := os.Getenv("DB_HOST")
    database := os.Getenv("DB_DATABASE")
    dbConfig := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

    db, err := gorm.Open("mysql", dbConfig)
    if err != nil {
        return nil, err
    }

    // 接続成功
    return db, nil
}

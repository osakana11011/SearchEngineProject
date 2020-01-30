package datastore

import (
    "search_engine_project/crawler/src/domain/model/entity"
)

// MigrateAll はGormを利用してマイグレーション処理を実行する
func MigrateAll() error {
    // 接続処理
    db, err := NewGormDBConnection()
    if err != nil {
        return err
    }
    defer db.Close()

    // マイグレーション処理
    db.AutoMigrate(&entity.Document{})
    db.AutoMigrate(&entity.Domain{})
    db.AutoMigrate(&entity.Token{})
    db.AutoMigrate(&entity.CrawlWaiting{})
    db.AutoMigrate(&entity.InvertedData{})

    return nil
}

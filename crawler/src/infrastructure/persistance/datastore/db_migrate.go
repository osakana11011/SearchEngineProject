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

// DropAll はGormを利用して全てのテーブル情報を削除する
func DropAll() error {
    // 接続処理
    db, err := NewGormDBConnection()
    if err != nil {
        return err
    }
    defer db.Close()

    // マイグレーション処理
    db.DropTableIfExists(&entity.Document{})
    db.DropTableIfExists(&entity.Domain{})
    db.DropTableIfExists(&entity.Token{})
    db.DropTableIfExists(&entity.CrawlWaiting{})
    db.DropTableIfExists(&entity.InvertedData{})

    return nil
}

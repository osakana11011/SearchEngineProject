package datastore

import (
	"search_engine_project/crawler/src/domain/model/newentity"
)

// MigrateAll はGormを利用してマイグレーション処理を実行する
func MigrateAll() error {
	// 接続処理
	db, err := connectGormDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// マイグレーション処理
	db.AutoMigrate(&newentity.Document{})
	db.AutoMigrate(&newentity.Domain{})
	db.AutoMigrate(&newentity.Token{})
	db.AutoMigrate(&newentity.CrawlWaiting{})
	db.AutoMigrate(&newentity.InvertedData{})

	return nil
}

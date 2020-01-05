package datastore

import (
	"search_engine_project/crawler/domain/model/entity"
	"search_engine_project/crawler/domain/repository"
)

// PageRepository ...
type PageRepository struct {}

// NewPageRepository ...
func NewPageRepository() repository.PageRepository {
	return &PageRepository{}
}

// Regist ...
func (r *PageRepository) Regist(page entity.Page) (int64, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	// 登録処理
	res, err := db.Exec("INSERT INTO pages(title, url, created_at, updated_at) VALUES(?, ?, NOW(), NOW())", page.Title, page.URL)
	if err != nil {
		return -1, err
	}

	// 登録したばかりのIDを取得して返す
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil
}

// GetCountsByURL ...
func (r *PageRepository) GetCountsByURL(url string) (int, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// 件数の取得
	var counts int
	db.QueryRow("SELECT COUNT(id) FROM pages WHERE url = ?", url).Scan(&counts)

	return counts, nil
}

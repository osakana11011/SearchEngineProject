package datastore

import (
	"search_engine_project/crawler/domain/repository"
)

// InvertedIndexRepository ...
type InvertedIndexRepository struct {}

// NewInvertedIndexRepository ...
func NewInvertedIndexRepository() repository.InvertedIndexRepository {
	return &InvertedIndexRepository{}
}

// Regist ...
func (r *InvertedIndexRepository) Regist(pageID int64, wordID int64, counts int) error {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// 登録処理
	db.Exec("INSERT INTO inverted_index(page_id, word_id, counts, created_at, updated_at) VALUES(?, ?, ?, NOW(), NOW())", pageID, wordID, counts)

	return nil
}

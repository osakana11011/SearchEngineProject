package datastore

import (
	"search_engine_project/crawler/domain/repository"
)

// WordRepository ...
type WordRepository struct {}

// NewWordRepository ...
func NewWordRepository() repository.WordRepository {
	return &WordRepository{}
}

// Regist ...
func (r *WordRepository) Regist(word string) (int64, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	// 登録処理
	res, err := db.Exec("INSERT INTO words(word, created_at, updated_at) VALUES(?, NOW(), NOW())", word)
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

// GetID ...
func (r *WordRepository) GetID(word string) (int64, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// IDの取得
	var id int64
	db.QueryRow("SELECT id FROM words WHERE word = ? LIMIT 1", word).Scan(&id)

	return id, nil
}

// GetCounts ...
func (r *WordRepository) GetCounts(word string) (int, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// 件数の取得
	var counts int
	db.QueryRow("SELECT COUNT(id) FROM words WHERE word = ?", word).Scan(&counts)

	return counts, nil
}

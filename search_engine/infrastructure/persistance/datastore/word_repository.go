package datastore

import (
	"search_engine_project/search_engine/domain/repository"
)

// WordRepository ...
type WordRepository struct {}

// NewWordRepository ...
func NewWordRepository() repository.WordRepository {
	return &WordRepository{}
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

// GetIDs ...
func (r *WordRepository) GetIDs(words []string) ([]int64, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return []int64{}, err
	}
	defer db.Close()

	// IDの取得
	// var id int64
	// db.QueryRow("SELECT id FROM words WHERE word = ? LIMIT 1", word).Scan(&id)

	return []int64{}, nil
}

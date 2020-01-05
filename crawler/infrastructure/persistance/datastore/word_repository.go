package datastore

import (
	"fmt"
	"math"

	"search_engine_project/crawler/domain/repository"
)

// WordRepository ...
type WordRepository struct {}

// NewWordRepository ...
func NewWordRepository() repository.WordRepository {
	return &WordRepository{}
}

// BulkInsert ...
func (r *WordRepository) BulkInsert(words []string) error {
	// 登録する単語が0の時はそのまま返す
	if len(words) == 0 {
		return nil
	}

	// DB接続
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// バルクインサートする回数
	bulkNum := (int)(len(words) / 100.0) + 1
	for i := 0; i < bulkNum; i++ {
		// バルクインサート用のSQLを構築
		bulkInsertSQL := "INSERT IGNORE INTO words (word, created_at, updated_at) VALUES "
		to := i * 100
		from := (int)(math.Min((float64)(i*100+100), (float64)(len(words))))
		fmt.Println(to, from)
		wordsMass := words[to:from]
		for _, word := range wordsMass {
			if word == "" {
				continue
			}
			bulkInsertSQL += fmt.Sprintf("('%s', NOW(), NOW()), ", word)
		}
		bulkInsertSQL = bulkInsertSQL[:len(bulkInsertSQL)-2]

		// 登録処理
		fmt.Println(bulkInsertSQL)
		_, err := db.Exec(bulkInsertSQL)
		if err != nil {
			return err
		}
	}

	return nil
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

package datastore

import (
	"fmt"
	"math"
	"strings"

	"search_engine_project/crawler/src/domain/repository"
)

// TokenRepository ...
type TokenRepository struct {}

// NewTokenRepository ...
func NewTokenRepository() repository.TokenRepository {
	return &TokenRepository{}
}

// BulkInsert ...
func (r *TokenRepository) BulkInsert(tokens []string) error {
	// 登録する単語が0の時はそのまま返す
	if len(tokens) == 0 {
		return nil
	}

	// DB接続
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// バルクインサートする回数
	bulkNum := (int)(len(tokens) / 100.0) + 1
	for i := 0; i < bulkNum; i++ {
		// バルクインサート用のSQLを構築
		bulkInsertSQL := "INSERT IGNORE INTO tokens (token, created_at, updated_at) VALUES "
		to := i * 100
		from := (int)(math.Min((float64)(i*100+100), (float64)(len(tokens))))
		tokensMass := tokens[to:from]
		for _, token := range tokensMass {
			token = strings.NewReplacer("'", "",).Replace(token)
			bulkInsertSQL += fmt.Sprintf("('%s', NOW(), NOW()), ", token)
		}
		bulkInsertSQL = bulkInsertSQL[:len(bulkInsertSQL)-2]

		// 登録処理
		_, err := db.Exec(bulkInsertSQL)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetID ...
func (r *TokenRepository) GetID(token string) (int64, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// IDの取得
	var id int64
	db.QueryRow("SELECT id FROM tokens WHERE token = ? LIMIT 1", token).Scan(&id)

	return id, nil
}

// GetCounts ...
func (r *TokenRepository) GetCounts(token string) (int, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// 件数の取得
	var counts int
	db.QueryRow("SELECT COUNT(id) FROM tokens WHERE token = ?", token).Scan(&counts)

	return counts, nil
}

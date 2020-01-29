package datastore

import (
	"fmt"
	"strings"
	"math"
	"search_engine_project/crawler/src/domain/model/newentity"
	"search_engine_project/crawler/src/domain/repository"
)

// NewTokenRepository はトークンに関するDB操作を提供する構造体を生成して返す。
func NewTokenRepository() repository.TokenRepository {
	return &tokenRepository{}
}

type tokenRepository struct {}

func (r *tokenRepository) Insert(token newentity.Token) error {
	// DB接続
	db, err := connectGormDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Create(&token)

	return nil
}

func (r *tokenRepository) BulkInsert(tokens []newentity.Token) error {
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
    bulkNum := (int)(math.Ceil(float64(len(tokens)) / 100.0))
    for i := 0; i < bulkNum; i++ {
        // バルクインサート用のSQLを構築
        bulkInsertSQL := "INSERT IGNORE INTO tokens (name, created_at, updated_at) VALUES "
        to := i * 100
        from := (int)(math.Min((float64)(i*100+100), (float64)(len(tokens))))
        tokensMass := tokens[to:from]
        for _, token := range tokensMass {
            tokenName := strings.NewReplacer("'", "",).Replace(token.Name)
            bulkInsertSQL += fmt.Sprintf("('%s', NOW(), NOW()), ", tokenName)
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

func (r *tokenRepository) GetTokensByName(tokenNames []string) ([]newentity.Token, error) {
    // DB接続
    db, err := connectGormDB()
    if err != nil {
        return []newentity.Token{}, err
    }
    defer db.Close()

    var tokens []newentity.Token
    db.Where("name IN (?)", tokenNames).Find(&tokens)

	return tokens, nil
}

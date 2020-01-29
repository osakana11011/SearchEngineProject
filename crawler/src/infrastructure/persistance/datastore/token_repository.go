package datastore

import (
	"fmt"
	"strings"
	"math"
	"search_engine_project/crawler/src/domain/model/entity"
    "search_engine_project/crawler/src/domain/repository"
    "github.com/jinzhu/gorm"
)

// NewTokenRepository はトークンに関するDB操作を提供する構造体を生成して返す。
func NewTokenRepository(db *gorm.DB) repository.TokenRepository {
	return &tokenRepository{db: db}
}

type tokenRepository struct {
    db *gorm.DB
}

func (r *tokenRepository) Insert(token entity.Token) error {
	r.db.Create(&token)

	return nil
}

func (r *tokenRepository) BulkInsert(tokens []entity.Token) error {
	// 登録する単語が0の時はそのまま返す
    if len(tokens) == 0 {
        return nil
    }

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
        r.db.Exec(bulkInsertSQL)
    }

    return nil
}

func (r *tokenRepository) GetTokensByName(tokenNames []string) ([]entity.Token, error) {
    var tokens []entity.Token
    r.db.Where("name IN (?)", tokenNames).Find(&tokens)

	return tokens, nil
}

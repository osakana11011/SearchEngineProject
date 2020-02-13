package datastore

import (
	"search_engine_project/search_engine/src/domain/model/entity"
	"search_engine_project/search_engine/src/domain/repository"
	"github.com/jinzhu/gorm"
)

// NewTokenRepository はrepository.TokenRepository実装した構造体を返す
func NewTokenRepository(db *gorm.DB) repository.TokenRepository {
	return &tokenRepository{db: db}
}

type tokenRepository struct {
	db *gorm.DB
}

func (r *tokenRepository) GetByTokenName(tokenName string) (entity.Token, error) {
	var token entity.Token

	r.db.Take(&token)

	return token, nil
}

func (r *tokenRepository) GetByTokenNames(tokenNames []string) ([]entity.Token, error) {
	var tokens []entity.Token

	r.db.Where("name IN (?)", tokenNames).Find(&tokens)

	return tokens, nil
}

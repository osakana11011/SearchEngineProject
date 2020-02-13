package repository

import (
	"search_engine_project/search_engine/src/domain/model/entity"
)

// TokenRepository はトークンに関するDB操作を提供するインターフェース
type TokenRepository interface {
	GetByTokenName(tokenName string) (entity.Token, error)
	GetByTokenNames(tokenNames []string) ([]entity.Token, error)
}

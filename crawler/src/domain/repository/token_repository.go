package repository

import (
    "search_engine_project/crawler/src/domain/model/entity"
)

// TokenRepository はトークンに関するDB操作を抽象化するインターフェース
type TokenRepository interface {
    Insert(token entity.Token) error
    BulkInsert(tokens []entity.Token) error
    GetTokensByName(tokenNames []string) ([]entity.Token, error)
}

package repository

import (
    "search_engine_project/crawler/src/domain/model/newentity"
)

// TokenRepository はトークンに関するDB操作を抽象化するインターフェース
type TokenRepository interface {
    Insert(token newentity.Token) error
    BulkInsert(tokens []newentity.Token) error
    GetTokensByName(tokenNames []string) ([]newentity.Token, error)
}

package repository

// TokenRepository はトークンのDB操作に関するインターフェース
type TokenRepository interface {
    BulkInsert(tokens []string) error
    GetIDs(tokens []string) (map[string]int64, error)
    GetID(token string) (int64, error)
    GetCounts(token string) (int, error)
}

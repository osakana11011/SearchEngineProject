package repository

// TokenRepository ...
type TokenRepository interface {
	BulkInsert(tokens []string) error
	GetID(token string) (int64, error)
	GetCounts(token string) (int, error)
}

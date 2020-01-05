package repository

// WordRepository ...
type WordRepository interface {
	BulkInsert(words []string) error
	GetID(word string) (int64, error)
	GetCounts(word string) (int, error)
}

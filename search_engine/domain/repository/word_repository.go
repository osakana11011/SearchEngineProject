package repository

// WordRepository ...
type WordRepository interface {
	GetID(word string) (int64, error)
	GetIDs(words []string) ([]int64, error)
}

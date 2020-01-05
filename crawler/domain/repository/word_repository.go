package repository

// WordRepository ...
type WordRepository interface {
	Regist(word string) (int64, error)
	GetCounts(word string) (int, error)
}

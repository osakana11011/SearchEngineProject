package repository

// InvertedIndexRepository ...
type InvertedIndexRepository interface {
	Regist(pageID int64, wordID int64, counts int) error
}

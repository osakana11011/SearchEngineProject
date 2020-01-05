package repository

type WordRepository interface {
	Regist (word string) (int64, error)
}

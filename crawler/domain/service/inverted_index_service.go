package service

import (
	"search_engine_project/crawler/domain/model/entity"
	"search_engine_project/crawler/infrastructure/persistance/datastore"
)

// InvertedIndex ...
type InvertedIndex interface {
	Regist(*entity.InvertedIndex)
}

type invertedIndex struct {}

// NewInvertedIndex ...
func NewInvertedIndex() InvertedIndex {
	return &invertedIndex{}
}

// Regist ...
func (x *invertedIndex) Regist(invertedIndex *entity.InvertedIndex) {
	wordRepository := datastore.NewWordRepository()
	invertedIndexRepository := datastore.NewInvertedListRepository()

	wordRepository.BulkInsert(invertedIndex.Words)
	invertedIndexRepository.BulkInsert(invertedIndex.InvertedList)
}

package service

import (
	"search_engine_project/crawler/domain/model/entity"
	"search_engine_project/crawler/infrastructure/persistance/datastore"
)

// InvertedIndexService ...
type InvertedIndexService interface {
	Regist(*entity.InvertedIndex)
}

type invertedIndexService struct {}

// NewInvertedIndexService ...
func NewInvertedIndexService() InvertedIndexService {
	return &invertedIndexService{}
}

// Regist ...
func (x *invertedIndexService) Regist(invertedIndex *entity.InvertedIndex) {
	wordRepository := datastore.NewWordRepository()
	invertedListRepository := datastore.NewInvertedListRepository()

	wordRepository.BulkInsert(invertedIndex.WordDictionary)
	invertedListRepository.BulkInsert(invertedIndex.InvertedList)
}

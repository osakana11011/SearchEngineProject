package service

import (
	"search_engine_project/crawler/src/domain/model/entity"
	"search_engine_project/crawler/src/infrastructure/persistance/datastore"
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
	tokenRepository := datastore.NewTokenRepository()
	invertedListRepository := datastore.NewInvertedListRepository()

	tokenRepository.BulkInsert(invertedIndex.TokenDictionary)
	invertedListRepository.BulkInsert(invertedIndex.InvertedList)
}

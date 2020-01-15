package service

import (
	"search_engine_project/search_engine/infrastructure/persistance/datastore"
	"search_engine_project/search_engine/domain/model/entity"
)

// SearchService は文書に関する様々な処理を呼び出す為の窓口。
type SearchService interface {
	Search(q string) ([]entity.Document, error)
}

type searchService struct {}

// NewSearchService はSearchServiceを扱うインスタンスを提供する。
func NewSearchService() SearchService {
	return &searchService{}
}

func (x *searchService) Search(q string) ([]entity.Document, error) {
	documentRepository := datastore.NewDocumentRepository()

	documents, err := documentRepository.GetDocuments(q)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

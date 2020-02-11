package service

import (
    "search_engine_project/search_engine/src/domain/model/entity"
    "search_engine_project/search_engine/src/domain/repository"
)

// SearchService は文書に関する様々な処理を呼び出す為の窓口。
type SearchService interface {
    Search(title string) ([]entity.Document, error)
}

type searchService struct {
    documentRepo repository.DocumentRepository
}

// NewSearchService はSearchServiceを扱うインスタンスを提供する。
func NewSearchService(documentRepo repository.DocumentRepository) SearchService {
    return &searchService{documentRepo: documentRepo}
}

func (ss *searchService) Search(title string) ([]entity.Document, error) {
    return ss.documentRepo.GetByTitle(title)
}

package service

import (
    "fmt"
    "search_engine_project/search_engine/src/domain/model/entity"
    "search_engine_project/search_engine/src/domain/model/valueobject"
    "search_engine_project/search_engine/src/domain/repository"
)

// SearchService は文書に関する様々な処理を呼び出す為の窓口。
type SearchService interface {
    Search(title string) ([]entity.Document, error)
}

type searchService struct {
    tokenRepo repository.TokenRepository
    documentRepo repository.DocumentRepository
}

// NewSearchService はSearchServiceを扱うインスタンスを提供する。
func NewSearchService(tokenRepo repository.TokenRepository, documentRepo repository.DocumentRepository) SearchService {
    return &searchService{tokenRepo: tokenRepo, documentRepo: documentRepo}
}

func (ss *searchService) Search(q string) ([]entity.Document, error) {
    query := valueobject.NewQueryFromString(q)

    // fmt.Println(query.QueryStrings)

    tokens, err := ss.tokenRepo.GetByTokenNames(query.QueryStrings)
    if err != nil {
        return []entity.Document{}, err
    }

    fmt.Println(tokens)

    return []entity.Document{}, nil
}

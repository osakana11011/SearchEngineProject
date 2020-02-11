package usecase

import (
    "search_engine_project/search_engine/src/domain/model/entity"
    "search_engine_project/search_engine/src/domain/service"
)

// SearchUseCase はPresentation層から呼ぶユースケース機能を提供する。
type SearchUseCase interface {
    Search(q string) ([]entity.Document, error)
}

// NewSearchUseCase はSearchUseCaseインターフェースを実装した構造体を返す。
func NewSearchUseCase(searchService service.SearchService) SearchUseCase {
    return &searchUseCase{searchService: searchService}
}

type searchUseCase struct {
    searchService service.SearchService
}

func (u *searchUseCase) Search(q string) ([]entity.Document, error) {
    return u.searchService.Search(q)
}

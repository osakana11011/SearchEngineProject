package repository

import (
    "search_engine_project/search_engine/src/domain/model/entity"
)

// DocumentRepository は文書に関するDB操作を抽象化するインターフェース
type DocumentRepository interface {
	GetByIDs(documentIDs []uint) ([]entity.Document, error)
	GetByTitle(title string) ([]entity.Document, error)
}

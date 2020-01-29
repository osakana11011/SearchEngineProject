package repository

import (
    "search_engine_project/crawler/src/domain/model/entity"
)

// DocumentRepository はDocumentに関するDB操作を抽象化するインターフェース
type DocumentRepository interface {
    Insert(document entity.Document) (uint, error)
    GetByURL(url string) (entity.Document, error)
    Update(document entity.Document) error
}

package repository

import (
    "search_engine_project/crawler/src/domain/model/newentity"
)

// DocumentRepository はDocumentに関するDB操作を抽象化するインターフェース
type DocumentRepository interface {
    Insert(document newentity.Document) (uint, error)
    GetByURL(url string) (newentity.Document, error)
    Update(document newentity.Document) error
}

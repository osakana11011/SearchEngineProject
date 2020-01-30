package repository

import (
    "search_engine_project/crawler/src/domain/model/entity"
)

// DocumentRepository は文書に関するDB操作を抽象化するインターフェース
type DocumentRepository interface {
    Insert(document entity.Document) (uint, error)
    GetByURL(url string) (entity.Document, error)
}

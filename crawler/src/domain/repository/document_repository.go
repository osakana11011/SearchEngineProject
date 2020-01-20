package repository

import (
    "search_engine_project/crawler/src/domain/model/entity"
)

// DocumentRepository は文書のDB操作に関するインターフェース
type DocumentRepository interface {
    Regist(page entity.Document) (int64, error)
    GetCountsByURL(url string) (int, error)
}

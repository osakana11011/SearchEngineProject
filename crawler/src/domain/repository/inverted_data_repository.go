package repository

import (
    "search_engine_project/crawler/src/domain/model/newentity"
)

// InvertedDataRepository は転置データに関するDB操作を抽象化するインターフェース
type InvertedDataRepository interface {
    Insert(invertedData newentity.InvertedData) error
    BulkInsert(invertedData []newentity.InvertedData) error
}

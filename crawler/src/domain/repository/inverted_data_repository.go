package repository

import (
    "search_engine_project/crawler/src/domain/model/entity"
)

// InvertedDataRepository は転置データに関するDB操作を抽象化するインターフェース
type InvertedDataRepository interface {
    Insert(invertedData entity.InvertedData) error
    BulkInsert(invertedData []entity.InvertedData) error
}

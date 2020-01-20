package repository

import (
    "search_engine_project/crawler/src/domain/model/entity"
)

// InvertedListRepository は転置リストのDBに関するインターフェース
type InvertedListRepository interface {
    BulkInsert(invertedList entity.InvertedList) error
}

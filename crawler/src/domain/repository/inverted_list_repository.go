package repository

import (
	"search_engine_project/crawler/src/domain/model/entity"
)

// InvertedListRepository ...
type InvertedListRepository interface {
	BulkInsert(invertedList entity.InvertedList) error
}

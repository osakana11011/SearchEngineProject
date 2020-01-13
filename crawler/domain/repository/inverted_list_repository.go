package repository

import (
	"search_engine_project/crawler/domain/model/entity"
)

// InvertedListRepository ...
type InvertedListRepository interface {
	BulkInsert(invertedList entity.InvertedList) error
}

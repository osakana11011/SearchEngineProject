package repository

import (
	"search_engine_project/crawler/domain/model/entity"
)

// DocumentRepository ...
type DocumentRepository interface {
	Regist(page entity.Document) (int64, error)
	GetCountsByURL(url string) (int, error)
}

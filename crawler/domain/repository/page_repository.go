package repository

import (
	"search_engine_project/crawler/domain/model/entity"
)

type PageRepository interface {
	Regist(page entity.Page) (int64, error)
	GetCountsByURL(url string) (int, error)
}

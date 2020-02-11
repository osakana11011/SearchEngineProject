package repository

import (
	"search_engine_project/search_engine/src/domain/model/entity"
)

type InvertedDataRepository interface {
	GetByToken(token entity.Token, page int, n int) int
}

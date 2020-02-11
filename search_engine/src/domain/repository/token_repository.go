package repository

import (
	"search_engine_project/search_engine/src/domain/model/entity"
)

type TokenRepository interface {
	GetByTokenName(tokenName string) (entity.Token, error)
}

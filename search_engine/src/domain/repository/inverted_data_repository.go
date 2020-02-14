package repository

import (
	"search_engine_project/search_engine/src/domain/model/entity"
)

// InvertedDataRepository は転置データに関するDB操作を提供するインターフェース
type InvertedDataRepository interface {
	GetByToken(token entity.Token) ([]entity.InvertedData, error)
}

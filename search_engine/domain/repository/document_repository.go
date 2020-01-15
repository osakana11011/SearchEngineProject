package repository

import (
	"search_engine_project/search_engine/domain/model/entity"
)

// DocumentRepository ...
type DocumentRepository interface {
	GetDocuments(q string) ([]entity.Document, error)
}

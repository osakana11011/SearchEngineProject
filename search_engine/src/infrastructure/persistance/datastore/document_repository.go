package datastore

import (
	"search_engine_project/search_engine/src/domain/model/entity"
	"search_engine_project/search_engine/src/domain/repository"
	"github.com/jinzhu/gorm"
)

// NewDocumentRepository repository.DocumentRepositoryを実装した構造体を返す
func NewDocumentRepository(db *gorm.DB) repository.DocumentRepository {
	return &documentRepository{db: db}
}

// DocumentRepository は文書に関するDB操作を提供する
type documentRepository struct {
	db *gorm.DB
}

func (r *documentRepository) GetByTitle(title string) ([]entity.Document, error) {
	var documents []entity.Document
	r.db.Limit(10).Select("*").Find(&documents)

	return documents, nil
}

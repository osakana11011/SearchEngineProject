package datastore

import (
    "search_engine_project/search_engine/src/domain/model/entity"
    "search_engine_project/search_engine/src/domain/repository"
    "github.com/jinzhu/gorm"
)

// NewInvertedDataRepository はrepository.InvertedDataRepositoryを実装した構造体を返す
func NewInvertedDataRepository(db *gorm.DB) repository.InvertedDataRepository {
    return &invertedDataRepository{db: db}
}

type invertedDataRepository struct {
    db *gorm.DB
}

func (r *invertedDataRepository) GetByToken(token entity.Token) ([]entity.InvertedData, error) {
    var invertedList []entity.InvertedData
    r.db.Where("token_id = ?", token.ID).Find(&invertedList)

    return invertedList, nil
}

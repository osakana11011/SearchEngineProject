package datastore

import (
    "search_engine_project/crawler/src/domain/model/entity"
    "search_engine_project/crawler/src/domain/repository"
    "github.com/jinzhu/gorm"
    "github.com/t-tiger/gorm-bulk-insert"
)

// NewInvertedDataRepository はrepository.InvertedDataRepositoryを実装した構造体を返す
func NewInvertedDataRepository(db *gorm.DB) repository.InvertedDataRepository {
    return &invertedDataRepository{db: db}
}

type invertedDataRepository struct {
    db *gorm.DB
}

func (r *invertedDataRepository) Insert(invertedData entity.InvertedData) error {
    r.db.Create(&invertedData)

    return nil
}

func (r *invertedDataRepository) BulkInsert(invertedData []entity.InvertedData) error {
    var insertRecords []interface{}
    for _, data := range invertedData {
        insertRecords = append(insertRecords, data)
    }

    if err := gormbulk.BulkInsert(r.db, insertRecords, 2000); err != nil {
		return err
	}

    return nil
}

package datastore

import (
    "search_engine_project/crawler/src/domain/model/newentity"
    "search_engine_project/crawler/src/domain/repository"
    "github.com/t-tiger/gorm-bulk-insert"
)

func NewInvertedDataRepository() repository.InvertedDataRepository {
    return &invertedDataRepository{}
}

type invertedDataRepository struct {}

func (r *invertedDataRepository) Insert(invertedData newentity.InvertedData) error {
    // DB接続
    db, err := connectGormDB()
    if err != nil {
        return err
    }
    defer db.Close()

    db.Create(&invertedData)

    return nil
}

func (r *invertedDataRepository) BulkInsert(invertedData []newentity.InvertedData) error {
    // DB接続
    db, err := connectGormDB()
    if err != nil {
        return err
    }
    defer db.Close()

    var insertRecords []interface{}
    for _, data := range invertedData {
        insertRecords = append(insertRecords, data)
    }

    if err := gormbulk.BulkInsert(db, insertRecords, 2000); err != nil {
		return err
	}

    return nil
}

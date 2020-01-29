package datastore

import (
    "search_engine_project/crawler/src/domain/model/newentity"
    "search_engine_project/crawler/src/domain/repository"
)


// NewDocumentRepository はDocumentRepositoryハンドラを返す。
func NewDocumentRepository() repository.DocumentRepository {NewDomainRepository()
    return &documentRepository{NewDomainRepository()}
}

type documentRepository struct {
    DomainRepository repository.DomainRepository
}

// Insert は文書情報をDBに登録する。
// 登録に成功した場合、その文書IDも返す。
func (r *documentRepository) Insert(document newentity.Document) (uint, error) {
    // DB接続
    db, err := connectGormDB()
    if err != nil {
        return 0, err
    }
    defer db.Close()

    document.Domain, err = r.DomainRepository.FirstOrCreate(document.Domain.Name)
    if err != nil {
        return 0, err
    }
    db.Create(&document)

    return document.ID, nil
}

// GetCountsByURL は該当するURLを持つデータの個数を返す。
func (r *documentRepository) GetByURL(url string) (newentity.Document, error) {
    // DB接続
    db, err := connectGormDB()
    if err != nil {
        return newentity.Document{}, err
    }
    defer db.Close()

    var document newentity.Document
    db.Where("url = ?", url).Take(&document)

    return document, nil
}

func (r *documentRepository) Update(document newentity.Document) error {
    // DB接続
    db, err := connectGormDB()
    if err != nil {
        return err
    }
    defer db.Close()

    db.Save(&document)

    return nil
}

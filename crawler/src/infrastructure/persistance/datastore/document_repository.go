package datastore

import (
    "search_engine_project/crawler/src/domain/model/entity"
    "search_engine_project/crawler/src/domain/repository"
    "github.com/jinzhu/gorm"
)


// NewDocumentRepository はDocumentRepositoryハンドラを返す。
func NewDocumentRepository(db *gorm.DB, domainRepo repository.DomainRepository) repository.DocumentRepository {
    return &documentRepository{db: db, domainRepo: domainRepo}
}

type documentRepository struct {
    db *gorm.DB
    domainRepo repository.DomainRepository
}

// Insert は文書情報をDBに登録する。
// 登録に成功した場合、その文書IDも返す。
func (r *documentRepository) Insert(document entity.Document) (uint, error) {
    document.Domain, _ = r.domainRepo.FirstOrCreate(document.Domain.Name)
    r.db.Create(&document)

    return document.ID, nil
}

// GetCountsByURL は該当するURLを持つデータの個数を返す。
func (r *documentRepository) GetByURL(url string) (entity.Document, error) {
    var document entity.Document
    r.db.Where("url = ?", url).Take(&document)

    return document, nil
}

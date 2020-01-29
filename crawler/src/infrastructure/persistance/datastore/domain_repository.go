package datastore

import (
    "search_engine_project/crawler/src/domain/model/entity"
    "search_engine_project/crawler/src/domain/repository"
    "github.com/jinzhu/gorm"
)

// NewDomainRepository はrepository.DomainRepositoryを実装した構造体を返す。
func NewDomainRepository(db *gorm.DB) repository.DomainRepository {
	return &domainRepository{db: db}
}

type domainRepository struct {
    db *gorm.DB
}

func (r *domainRepository) Insert(domain entity.Domain) (uint, error) {
	r.db.Create(&domain)

    return domain.ID, nil
}

func (r *domainRepository) GetByDomainName(domainName string) (entity.Domain, error) {
    var domain entity.Domain
    r.db.Where("name = ?", domainName).Take(&domain)

    return domain, nil
}

func (r *domainRepository) FirstOrCreate(domainName string) (entity.Domain, error) {
    var domain entity.Domain
    r.db.Where("name = ?", domainName).Take(&domain)

    // テーブル内に存在しない場合は、新しく作成して返す
    if (domain.ID == 0) {
        domain.Name = domainName
        r.db.Create(&domain)
    }

    return domain, nil
}

package datastore

import (
    "search_engine_project/crawler/src/domain/model/newentity"
    "search_engine_project/crawler/src/domain/repository"
)

func NewDomainRepository() repository.DomainRepository {
	return &domainRepository{}
}

type domainRepository struct {}

func (r *domainRepository) Insert(domain newentity.Domain) (uint, error) {
	// DB接続
    db, err := connectGormDB()
    if err != nil {
        return 0, err
    }
    defer db.Close()

    db.Create(&domain)

    return domain.ID, nil
}

func (r *domainRepository) GetByDomainName(domainName string) (newentity.Domain, error) {
    // DB接続
    db, err := connectGormDB()
    if err != nil {
        return newentity.Domain{}, err
    }
    defer db.Close()

    var domain newentity.Domain
    db.Where("name = ?", domainName).Take(&domain)

    return domain, nil
}

func (r *domainRepository) FirstOrCreate(domainName string) (newentity.Domain, error) {
    // DB接続
    db, err := connectGormDB()
    if err != nil {
        return newentity.Domain{}, err
    }
    defer db.Close()

    // テーブル内に与えられたドメイン名を持つレコードが既にある場合はそれを返す
    var domain newentity.Domain
    db.Where("name = ?", domainName).Take(&domain)

    // テーブル内に存在しない場合は、新しく作成して返す
    if (domain.ID == 0) {
        domain.Name = domainName
        db.Create(&domain)
    }

    return domain, nil
}

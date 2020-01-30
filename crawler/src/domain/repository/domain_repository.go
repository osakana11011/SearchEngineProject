package repository

import (
    "search_engine_project/crawler/src/domain/model/entity"
)

// DomainRepository はドメインに関するDB操作を抽象化するインターフェース
type DomainRepository interface {
    Insert(domain entity.Domain) (uint, error)
    GetByDomainName(domainName string) (entity.Domain, error)
    FirstOrCreate(domainName string) (entity.Domain, error)
}

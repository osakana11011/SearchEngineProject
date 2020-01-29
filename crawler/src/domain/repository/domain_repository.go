package repository

import (
	"search_engine_project/crawler/src/domain/model/newentity"
)

type DomainRepository interface {
	Insert(domain newentity.Domain) (uint, error)
	GetByDomainName(domainName string) (newentity.Domain, error)
	FirstOrCreate(domainName string) (newentity.Domain, error)
}

package repository

import (
	"search_engine_project/crawler/src/domain/model/entity"
)

// CrawlWaitingRepository はCrawlWaitingに関するDB操作を抽象化するインターフェース
type CrawlWaitingRepository interface {
	Insert(crawlWaiting entity.CrawlWaiting) error
	BulkInsert(crawlWaiting []entity.CrawlWaiting) error
	GetTopPriority() (entity.CrawlWaiting, error)
	Update(crawlWaiting entity.CrawlWaiting) error
	Delete(crawlWaiting entity.CrawlWaiting) error
}
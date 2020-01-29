package repository

import (
	"search_engine_project/crawler/src/domain/model/newentity"
)

// CrawlWaitingRepository はCrawlWaitingに関するDB操作を抽象化するインターフェース
type CrawlWaitingRepository interface {
	Insert(crawlWaiting newentity.CrawlWaiting) error
	BulkInsert(crawlWaiting []newentity.CrawlWaiting) error
	GetTopPriority() (newentity.CrawlWaiting, error)
	Update(crawlWaiting newentity.CrawlWaiting) error
	Delete(crawlWaiting newentity.CrawlWaiting) error
}

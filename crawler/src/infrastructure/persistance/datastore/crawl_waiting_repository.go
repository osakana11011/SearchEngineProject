package datastore

import (
	"search_engine_project/crawler/src/domain/model/newentity"
	"search_engine_project/crawler/src/domain/repository"
	"github.com/t-tiger/gorm-bulk-insert"
)

type crawlWaitingRepository struct {}

func NewCrawlWaitingRepository() repository.CrawlWaitingRepository {
	return &crawlWaitingRepository{}
}

func (r *crawlWaitingRepository) Insert(crawlWaiting newentity.CrawlWaiting) error {
	// DB接続
	db, err := connectGormDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Create(&crawlWaiting)

	return nil
}

func (r *crawlWaitingRepository) BulkInsert(crawlWaitings []newentity.CrawlWaiting) error {
	// DB接続
	db, err := connectGormDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// バルクインサート
	var insertRecords []interface{}
	for _, crawlWaiting := range crawlWaitings {
		insertRecords = append(insertRecords, crawlWaiting)
	}
	if err := gormbulk.BulkInsert(db, insertRecords, 2000); err != nil {
		return err
	}

	return nil
}

func (r *crawlWaitingRepository) GetTopPriority() (newentity.CrawlWaiting, error) {
	// DB接続
	db, err := connectGormDB()
	if err != nil {
		return newentity.CrawlWaiting{}, err
	}
	defer db.Close()

	var crawlWaiting newentity.CrawlWaiting
	db.Order("is_priority DESC").Take(&crawlWaiting)

	return crawlWaiting, nil
}

func (r *crawlWaitingRepository) Update(crawlWaiting newentity.CrawlWaiting) error {
	// DB接続
	db, err := connectGormDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// 更新処理
	db.Save(&crawlWaiting)

	return nil
}

func (r *crawlWaitingRepository) Delete(crawlWaiting newentity.CrawlWaiting) error {
	// DB接続
	db, err := connectGormDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Delete(&crawlWaiting)

	return nil
}

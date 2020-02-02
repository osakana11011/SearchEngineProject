package datastore

import (
    "fmt"
    "search_engine_project/crawler/src/domain/model/entity"
    "search_engine_project/crawler/src/domain/repository"
    "github.com/jinzhu/gorm"
    "github.com/t-tiger/gorm-bulk-insert"
)

// NewCrawlWaitingRepository はrepository.CrawlWaitingRepositoryを実装した構造体を返す。
func NewCrawlWaitingRepository(db *gorm.DB) repository.CrawlWaitingRepository {
    return &crawlWaitingRepository{db: db}
}

type crawlWaitingRepository struct {
    db *gorm.DB
}

func (r *crawlWaitingRepository) Insert(crawlWaiting entity.CrawlWaiting) error {
    r.db.Create(&crawlWaiting)

    return nil
}

func (r *crawlWaitingRepository) BulkInsert(crawlWaitings []entity.CrawlWaiting) error {
    fmt.Println(crawlWaitings)
    var insertRecords []interface{}
    for _, crawlWaiting := range crawlWaitings {
        insertRecords = append(insertRecords, crawlWaiting)
    }

    if err := gormbulk.BulkInsert(r.db, insertRecords, 2000); err != nil {
        return err
    }

    return nil
}

func (r *crawlWaitingRepository) GetTopPriority() (entity.CrawlWaiting, error) {
    var crawlWaiting entity.CrawlWaiting
    r.db.Order("is_priority DESC").Take(&crawlWaiting)

    return crawlWaiting, nil
}

func (r *crawlWaitingRepository) Delete(crawlWaiting entity.CrawlWaiting) error {
    r.db.Delete(&crawlWaiting)

    return nil
}

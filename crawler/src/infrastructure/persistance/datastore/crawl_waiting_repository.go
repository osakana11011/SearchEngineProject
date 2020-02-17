package datastore

import (
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

func (r *crawlWaitingRepository) GetCounts() int {
    var count int
    var crawlWaitings []entity.CrawlWaiting
    r.db.Find(&crawlWaitings).Count(&count)

    return count
}

func (r *crawlWaitingRepository) Insert(crawlWaiting entity.CrawlWaiting) error {
    r.db.Create(&crawlWaiting)

    return nil
}

func (r *crawlWaitingRepository) BulkInsert(crawlWaitings []entity.CrawlWaiting) error {
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
    r.db.Where("deleted_at IS NULL").Order("is_priority DESC").Order("id ASC").Take(&crawlWaiting)

    return crawlWaiting, nil
}

func (r *crawlWaitingRepository) Delete(crawlWaiting entity.CrawlWaiting) error {
    r.db.Delete(&crawlWaiting)

    return nil
}

func (r *crawlWaitingRepository) HardDelete(crawlWaiting entity.CrawlWaiting) error {
    r.db.Unscoped().Delete(&crawlWaiting)

    return nil
}

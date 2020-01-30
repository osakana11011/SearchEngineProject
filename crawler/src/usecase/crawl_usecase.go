package usecase

import (
    "search_engine_project/crawler/src/domain/service"
)

// NewCrawlUsecase はクロールに関するユースケースを提供する構造体を返す
func NewCrawlUsecase(crawlService service.CrawlService, crawlWaitingService service.CrawlWaitingService) CrawlUsecase {
    return &crawlUsecase{
        crawlWaitingService: crawlWaitingService,
        crawlService: crawlService,
    }
}

// CrawlUsecase はクロールに関するユースケースを定義したインターフェース
type CrawlUsecase interface {
    ExecCrawlService() error
}

type crawlUsecase struct {
    crawlWaitingService service.CrawlWaitingService
    crawlService service.CrawlService
}

// ExecCrawlService はクロールサービスを開始する
func (u *crawlUsecase) ExecCrawlService() error {
    // クロールする対象を一つ取得する
    crawlWaiting, err := u.crawlWaitingService.GetValidTopPriority()
    if err != nil {
        return err
    }

    // クローリングを行う
    if err := u.crawlService.Crawl(crawlWaiting); err != nil {
        return err
    }

    return nil
}

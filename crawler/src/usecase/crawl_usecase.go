package usecase

import (
	"search_engine_project/crawler/src/domain/service"
)

// ExecCrawlService はクロールサービスを開始する
func ExecCrawlService() error {
	crawlWaitingService := service.NewCrawlWaitingService()
	crawlService := service.NewCrawlService()

	// クロールする対象を一つ取得する
	crawlWaiting, err := crawlWaitingService.GetValidTopPriority()
	if err != nil {
		return err
	}

	// クローリングを行う
	if err := crawlService.Crawl(crawlWaiting); err != nil {
		return err
	}


	return nil
}

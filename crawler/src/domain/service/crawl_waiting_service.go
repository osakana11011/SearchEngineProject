package service

import (
	"search_engine_project/crawler/src/domain/model/entity"
	"search_engine_project/crawler/src/domain/repository"
)

// CrawlWaitingService はクロール待ち情報を管理するサービスです。
type CrawlWaitingService interface {
	GetValidTopPriority() (entity.CrawlWaiting, error)
}

// NewCrawlWaitingService はCrawlWaitingServiceを動かす構造体を返す。
func NewCrawlWaitingService(crawlWaitingRepo repository.CrawlWaitingRepository) CrawlWaitingService {
	return &crawlWaitingService{crawlWaitingRepo: crawlWaitingRepo}
}

type crawlWaitingService struct {
	crawlWaitingRepo repository.CrawlWaitingRepository
}

// GetValidTopPriority は有効なクロール待ち情報から最も優先度の高いものを取得して返す。
func (s *crawlWaitingService) GetValidTopPriority() (entity.CrawlWaiting, error) {
	for {
		// 最も優先度の高いクロール待ち情報を取得
		topPriorityCrawlWaiting, err := s.crawlWaitingRepo.GetTopPriority()
		if err != nil {
			return entity.CrawlWaiting{}, err
		}

		// クロール済みであることを証明する為に、CrawledAtに現在日付を入れて更新
		s.crawlWaitingRepo.Delete(topPriorityCrawlWaiting)

		// 取り出した情報が有効であれば、それを返す。
		if topPriorityCrawlWaiting.IsValid() {
			return topPriorityCrawlWaiting, nil
		}
	}
}

package service

import (
	"search_engine_project/crawler/src/domain/model/newentity"
	"search_engine_project/crawler/src/domain/repository"
	"search_engine_project/crawler/src/infrastructure/persistance/datastore"
)

// CrawlWaitingService はクロール待ち情報を管理するサービスです。
type CrawlWaitingService interface {
	GetValidTopPriority() (newentity.CrawlWaiting, error)
}

// NewCrawlWaitingService はCrawlWaitingServiceを動かす構造体を返す。
func NewCrawlWaitingService() CrawlWaitingService {
	crawlWaitingRepository := datastore.NewCrawlWaitingRepository()
	return &crawlWaitingService{crawlWaitingRepository}
}

type crawlWaitingService struct {
	crawlWaitingRepository repository.CrawlWaitingRepository
}

// GetValidTopPriority は有効なクロール待ち情報から最も優先度の高いものを取得して返す。
func (s *crawlWaitingService) GetValidTopPriority() (newentity.CrawlWaiting, error) {
	for {
		// 最も優先度の高いクロール待ち情報を取得
		topPriorityCrawlWaiting, err := s.crawlWaitingRepository.GetTopPriority()
		if err != nil {
			return newentity.CrawlWaiting{}, err
		}

		// クロール済みであることを証明する為に、CrawledAtに現在日付を入れて更新
		s.crawlWaitingRepository.Delete(topPriorityCrawlWaiting)

		// 取り出した情報が有効であれば、それを返す。
		if topPriorityCrawlWaiting.IsValid() {
			return topPriorityCrawlWaiting, nil
		}
	}
}

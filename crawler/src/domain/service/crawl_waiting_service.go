package service

import (
    "time"
    "search_engine_project/crawler/src/domain/model/entity"
    "search_engine_project/crawler/src/domain/repository"
)

// CrawlWaitingService はクロール待ち情報を管理するサービスです。
type CrawlWaitingService interface {
    GetValidTopPriority() (entity.CrawlWaiting, error)
}

// NewCrawlWaitingService はCrawlWaitingServiceを動かす構造体を返す。
func NewCrawlWaitingService(
    crawlWaitingRepo repository.CrawlWaitingRepository,
    documentRepo repository.DocumentRepository) CrawlWaitingService {
    return &crawlWaitingService{crawlWaitingRepo: crawlWaitingRepo, documentRepo: documentRepo}
}

type crawlWaitingService struct {
    crawlWaitingRepo repository.CrawlWaitingRepository
    documentRepo repository.DocumentRepository
}

// GetValidTopPriority は有効なクロール待ち情報から最も優先度の高いものを取得して返す。
func (s *crawlWaitingService) GetValidTopPriority() (entity.CrawlWaiting, error) {
    for {
        // 最も優先度の高いクロール待ち情報を取得
        topPriorityData, err := s.crawlWaitingRepo.GetTopPriority()
        if err != nil {
            return entity.CrawlWaiting{}, err
        }
        // クロール済みの情報は要らないので、それらをハードデリートする
        s.crawlWaitingRepo.HardDelete(topPriorityData)

        // クロール待ちテーブルにデータが存在しない時、3秒待ってから再度クロール待ちデータを探す
        if topPriorityData.ID == 0 {
            time.Sleep(3 * time.Second)
            continue
        }

        // クロール待ちデータのURLの有効性を検証する
        if !topPriorityData.IsValid() {
            continue
        }

        // クロールしたいデータのURLが既にクロール済みかどうか検証する
        if d, _ := s.documentRepo.GetByURL(topPriorityData.URL); d.ID != 0 {
            continue
        }

        return topPriorityData, nil
    }
}

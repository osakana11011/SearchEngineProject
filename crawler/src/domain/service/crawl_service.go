package service

import (
    "search_engine_project/crawler/src/domain/model/entity"
    "search_engine_project/crawler/src/domain/repository"
)

// CrawlService はクローラーが提供する機能を定義するインターフェース。
type CrawlService interface {
    Crawl(crawlWaiting entity.CrawlWaiting) error
}

// NewCrawlService はクローラーに関するサービスを提供する構造体を返す。
func NewCrawlService(
    documentRepo repository.DocumentRepository,
    crawlWaitingRepo repository.CrawlWaitingRepository,
    tokenRepo repository.TokenRepository,
    invertedDataRepo repository.InvertedDataRepository,
    ) CrawlService {
    return &crawlService{
        documentRepo: documentRepo,
        crawlWaitingRepo: crawlWaitingRepo,
        tokenRepo: tokenRepo,
        invertedDataRepo: invertedDataRepo,
    }
}

type crawlService struct {
    documentRepo repository.DocumentRepository
    crawlWaitingRepo repository.CrawlWaitingRepository
    tokenRepo    repository.TokenRepository
    invertedDataRepo repository.InvertedDataRepository
}

func (s *crawlService) Crawl(crawlWaiting entity.CrawlWaiting) error {
    // 文書情報のクロール
    document, err := entity.GetDocumentByCrawl(crawlWaiting.URL)
    if err != nil {
        return err
    }

    // 文書の登録
    documentID, err := s.documentRepo.Insert(document)
    if err != nil {
        return err
    }

    // クロール対象の子リンク(次クロールする候補対象)を登録
    // クロール対象のデータは指数関数敵に増えるので、テーブル上に一定の閾値を超えている時はデータを入れないようにする
    counts := s.crawlWaitingRepo.GetCounts()
    if counts < 100000 {
        if err := s.crawlWaitingRepo.BulkInsert(document.ChildLinks); err != nil {
            return err
        }
    }

    // トークンの登録
    var uniqueTokens []entity.Token
    for _, tokenName := range document.UniqueTokens {
        uniqueTokens = append(uniqueTokens, entity.Token{Name: tokenName})
    }
    if err := s.tokenRepo.BulkInsert(uniqueTokens); err != nil {
        return err
    }

    // トークン情報からIDの取得
    tokens, err := s.tokenRepo.GetTokensByName(document.UniqueTokens)
    if err != nil {
        return err
    }

    // 転置リストの登録
    invertedList := entity.GetInvertedList(documentID, document, tokens)
    s.invertedDataRepo.BulkInsert(invertedList)

    return nil
}

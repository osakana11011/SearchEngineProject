package service

import (
	"fmt"
	"search_engine_project/crawler/src/domain/model/newentity"
	"search_engine_project/crawler/src/infrastructure/persistance/datastore"
)

// CrawlService はクローラーが提供する機能を定義するインターフェース。
type CrawlService interface {
	Crawl(crawlWaiting newentity.CrawlWaiting) error
}

// NewCrawlService はクローラーに関するサービスを提供する構造体を返す。
func NewCrawlService() CrawlService {
	return &crawlService{}
}

type crawlService struct {}

func (s *crawlService) Crawl(crawlWaiting newentity.CrawlWaiting) error {
	documentRepository := datastore.NewDocumentRepository()
	tokenRepository := datastore.NewTokenRepository()
	invertedDataRepository := datastore.NewInvertedDataRepository()

	// 文書情報のクロール
	document, err := newentity.GetDocumentByCrawl(crawlWaiting.URL)
	if err != nil {
		return err
	}

	// 文書の登録
	documentID, err := documentRepository.Insert(document)
	if err != nil {
		return err
	}
	fmt.Println(documentID)

	// トークンの登録
	var uniqueTokens []newentity.Token
	for _, tokenName := range document.UniqueTokens {
		uniqueTokens = append(uniqueTokens, newentity.Token{Name: tokenName})
	}
	if err := tokenRepository.BulkInsert(uniqueTokens); err != nil {
		return err
	}

	// トークン情報からIDの取得
	tokens, err := tokenRepository.GetTokensByName(document.UniqueTokens)
	if err != nil {
		return err
	}

	// 転置リストの登録
	invertedList := newentity.GetInvertedList(documentID, document, tokens)
	invertedDataRepository.BulkInsert(invertedList)

	return nil
}

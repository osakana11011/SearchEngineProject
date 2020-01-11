package service

import (
	"fmt"
	"regexp"

	"search_engine_project/crawler/domain/model/entity"
	"search_engine_project/crawler/infrastructure/persistance/datastore"
)

// PageService ...
type PageService interface {
	Crawl(url string, depth int) error
}

// pageService ...
type pageService struct {}

// NewPageService ...
func NewPageService() PageService {
	return &pageService{}
}

// Crawl ...
func (x *pageService) Crawl(url string, depth int) error {
	// クロールするドメインに制約を掛ける
	if !isAcceptDomain(url) {
		return nil
	}

	// ページ情報の取得
	page, err := entity.CrawlPage(url)
	if err != nil {
		return err
	}

	// 登録済みのページの場合は、ページ/単語/転置インデックスの更新を行わない
	isRegisted := isRegistedPage(url)
	if err != nil {
		return err
	}
	if !isRegisted {
		err := registPageInfo(page)
		if err != nil {
			return err
		}
	}

	// ページ内のリンクを巡回
	for _, link := range page.Links {
		x.Crawl(link, depth - 1)
	}

	return nil
}

// isAcceptDomain ...
// 日本語版Wikipediaからしか取得しないようにする
func isAcceptDomain(url string) bool {
	r := regexp.MustCompile(`^https://ja.wikipedia.org/`)
	return r.MatchString(url)
}

// isRegistedPage ...
func isRegistedPage(url string) bool {
	pageRepository := datastore.NewPageRepository()

	counts, err := pageRepository.GetCountsByURL(url)
	if err != nil {
		return false
	}

	return (counts > 0)
}

// registPageInfo ...
// TODO: Webページは更新が入ることがあるので、登録済みの場合は過去の情報を破棄して新しい情報を登録するようにしたい
func registPageInfo(page entity.Page) error {
	pageRepository := datastore.NewPageRepository()
	wordRepository := datastore.NewWordRepository()
	invertedIndexRepository := datastore.NewInvertedIndexRepository()

	// ページの登録
	pageID, err := pageRepository.Regist(page)
	if err != nil {
		return err
	}

	// 単語情報の登録(登録済みの単語は登録しない)
	words := []string{}
	for word := range page.NounWords {
		words = append(words, word)
	}
	err = wordRepository.BulkInsert(words)
	if err != nil {
		return err
	}

	// 転置インデックスへの登録
	for word, counts := range page.NounWords {
		wordID, _ := wordRepository.GetID(word)
		invertedIndexRepository.Regist(pageID, wordID, counts)
	}

	return nil
}

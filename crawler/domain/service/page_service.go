package service

import (
	"regexp"
	"errors"
	"fmt"
	"log"

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
	pageRepository := datastore.NewPageRepository()
	wordRepository := datastore.NewWordRepository()
	invertedIndexRepository := datastore.NewInvertedIndexRepository()

	// 日本語版Wikipediaからしか取得しないようにする
	r := regexp.MustCompile(`^https://ja.wikipedia.org/`)
	if !r.MatchString(url) {
		log.Println("Stop crawling because this URL's domain isn't 「ja.wikipecia.org」.")
		return errors.New("crawl only 'ja.wikipecia.org' domain")
	}

	// 既に登録されている場合はクロールしにいかない
	counts, err := pageRepository.GetCountsByURL(url)
	if err != nil {
		return err
	}
	if counts > 0 {
		return fmt.Errorf("already registed %s", url)
	}

	// クローリング
	page, err := entity.CrawlPage(url)
	if err != nil {
		return err
	}
	pageID, _ := pageRepository.Regist(page)

	// 単語の登録(登録済みの単語は登録しない)
	for word := range page.NounWords {
		counts, err := wordRepository.GetCounts(word)
		if err != nil || counts > 0 {
			continue
		}
		wordRepository.Regist(word)
	}

	// 転置インデックスへの登録
	for word, counts := range page.NounWords {
		wordID, _ := wordRepository.GetID(word)
		invertedIndexRepository.Regist(pageID, wordID, counts)
	}

	// ページ内のリンクを巡回
	for _, link := range page.Links {
		x.Crawl(link, depth - 1)
	}

	return nil
}

package service

import (
	"time"
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
	if depth <= 0 {
		return nil
	}

	// クロールするドメインに制約を掛ける
	if !isAcceptDomain(url) {
		return nil
	}

	// 登録済みのページの場合は、ページ/単語/転置インデックスの更新を行わない
	isRegisted := isRegistedPage(url)
	if isRegisted {
		return nil
	}

	// ページ情報の取得
	// サーバに負荷を掛けすぎないように自重
	time.Sleep(1 * time.Second)
	page, err := entity.CrawlPage(url)
	if err != nil {
		return err
	}

	// 登録済みのページの場合は、ページ/単語/転置インデックスの更新を行わない
	err = registPageInfo(page)
	if err != nil {
		return err
	}

	// ミニ転置インデックスに追加されているドキュメント数が一定数以上なら、DBへの登録を行う
	invertedIndex := entity.GetInvertedIndex()
	if invertedIndex.DocumentCounts >= 1 {
		invertedIndexService := NewInvertedIndex()
		invertedIndexService.Regist(invertedIndex)
		entity.InitInvertedIndex()
	}

	// ページ内のリンクを巡回
	for _, link := range page.Links {
		err := x.Crawl(link, depth - 1)

		if err != nil {
			return err
		}
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

	// ページの登録
	pageID, err := pageRepository.Regist(page)
	if err != nil {
		return err
	}

	// メモリ上に単語とミニ転置インデックスを登録
	invertedIndex := entity.GetInvertedIndex()
	invertedIndex.AddDocument(pageID, page)

	return nil
}

package service

import (
	"time"
	"regexp"

	"search_engine_project/crawler/domain/model/entity"
	"search_engine_project/crawler/infrastructure/persistance/datastore"
)

// DocumentService は文書に関する様々な処理を呼び出す為の窓口。
type DocumentService interface {
	Crawl(url string, depth int) error
}

type documentService struct {}

// NewDocumentService はDocumentServiceを扱うインスタンスを提供する。
func NewDocumentService() DocumentService {
	return &documentService{}
}

// Crawl が呼ばれると、urlを起点に深さdepthまで連鎖的にクローリングを行う。
func (x *documentService) Crawl(url string, depth int) error {
	// 最初指定した深さに達した場合は、クローリングを行わない。
	if depth <= 0 {
		return nil
	}

	// クロールするドメインに制約を掛ける。
	if !isAcceptDomain(url) {
		return nil
	}

	// 登録済みのページの場合は、ページ/単語/転置インデックスの更新を行わない
	isRegisted := isRegistedDocument(url)
	if isRegisted {
		return nil
	}

	// ページ情報の取得
	// サーバに負荷を掛けすぎないように自重
	time.Sleep(1 * time.Second)
	document, err := entity.GetDocumentByCrawl(url)
	if err != nil {
		return err
	}

	// 文書情報の登録
	documentRepository := datastore.NewDocumentRepository()
	documentID, err := documentRepository.Regist(document)
	if err != nil {
		return err
	}

	// メモリ上のミニ転置インデックスにドキュメント情報を追加する。
	invertedIndex := entity.GetInvertedIndex()
	invertedIndex.AddDocument(documentID, document)

	// ミニ転置インデックスに追加されているドキュメント数が一定数以上ならDBへの登録を行い、ミニ転置インデックスを初期化する
	if invertedIndex.DocumentCounts >= 1 {
		invertedIndexService := NewInvertedIndexService()
		invertedIndexService.Regist(invertedIndex)
		entity.InitInvertedIndex()
	}

	// ページ内のリンクを巡回
	for _, link := range document.Links {
		err := x.Crawl(link, depth - 1)

		if err != nil {
			return err
		}
	}

	return nil
}


func isAcceptDomain(url string) bool {
	r := regexp.MustCompile(`^https://ja.wikipedia.org/`)
	return r.MatchString(url)
}


func isRegistedDocument(url string) bool {
	documentRepository := datastore.NewDocumentRepository()

	counts, err := documentRepository.GetCountsByURL(url)
	if err != nil {
		return false
	}

	return (counts > 0)
}

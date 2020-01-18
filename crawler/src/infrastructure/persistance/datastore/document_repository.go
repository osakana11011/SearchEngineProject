package datastore

import (
	"search_engine_project/crawler/src/domain/model/entity"
	"search_engine_project/crawler/src/domain/repository"
)

// DocumentRepository はdocumentsテーブルに関するDB操作を提供する。
type DocumentRepository struct {}

// NewDocumentRepository はDocumentRepositoryハンドラを返す。
func NewDocumentRepository() repository.DocumentRepository {
	return &DocumentRepository{}
}

// Regist は文書情報をDBに登録する。
// 登録に成功した場合、その文書IDも返す。
func (r *DocumentRepository) Regist(document entity.Document) (int64, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	// 登録処理
	res, err := db.Exec("INSERT INTO documents(title, url, created_at, updated_at) VALUES(?, ?, NOW(), NOW())", document.Title, document.URL)
	if err != nil {
		return -1, err
	}

	// 登録したばかりのIDを取得して返す
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil
}

// GetCountsByURL は該当するURLを持つデータの個数を返す。
func (r *DocumentRepository) GetCountsByURL(url string) (int, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// 件数の取得
	var counts int
	db.QueryRow("SELECT COUNT(id) FROM documents WHERE url = ?", url).Scan(&counts)

	return counts, nil
}

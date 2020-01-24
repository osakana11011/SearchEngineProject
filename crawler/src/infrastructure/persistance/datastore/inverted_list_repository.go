package datastore

import (
    "fmt"
    "strings"

    "search_engine_project/crawler/src/domain/repository"
    "search_engine_project/crawler/src/domain/model/entity"
)

const (
    maxBuffer = 1000  // maxBufferトークン分のデータをバルクインサートする
)

// InvertedListRepository は転置リストテーブルを操作するハンドラ
type InvertedListRepository struct {}

// NewInvertedListRepository は転置リストテーブルを操作するハンドラを生成する
func NewInvertedListRepository() repository.InvertedListRepository {
    return &InvertedListRepository{}
}

// BulkInsert は転置リストをバルクインサートする
func (r *InvertedListRepository) BulkInsert(invertedList entity.InvertedList) error {
    // 転置リストが空の時は何も返さず終了
    if len(invertedList) == 0 {
        return nil
    }

    // DB接続
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    // トークンのIDリストを一括取得
    tokenRepository := NewTokenRepository()
    tokenLookUpTable, err := tokenRepository.GetIDs(keys(invertedList))
    if err != nil {
        return err
    }

    // 1000データ毎にバルクインサート処理を行う
    bulkInsertSQL := "INSERT IGNORE INTO inverted_list (token_id, document_id, tf, offset_list, created_at, updated_at) VALUES "
    buffN := 0
    for token, documentTokens := range invertedList {
        for documentID, documentToken := range documentTokens {
            offsetList := strings.Join(documentToken.OffsetList, ",")
            bulkInsertSQL += fmt.Sprintf("('%d', '%d', '%f', '%s', NOW(), NOW()), ", tokenLookUpTable[token], documentID, documentToken.TF, offsetList)
            buffN++
        }

        // バッファ値が上限を超えたら登録処理を行う
        if buffN >= maxBuffer {
            bulkInsertSQL = bulkInsertSQL[:len(bulkInsertSQL)-2]
            if _, err = db.Exec(bulkInsertSQL); err != nil {
                return err
            }
            bulkInsertSQL = "INSERT IGNORE INTO inverted_list (token_id, document_id, tf, offset_list, created_at, updated_at) VALUES "
            buffN = 0
        }
    }

    if buffN != 0 {
        bulkInsertSQL = bulkInsertSQL[:len(bulkInsertSQL)-2]
        if _, err := db.Exec(bulkInsertSQL); err != nil {
            return err
        }
    }

    return nil
}

func keys(invertedList entity.InvertedList) []string {
    keys := []string{}
    for key := range invertedList {
        keys = append(keys, key)
    }
    return keys
}

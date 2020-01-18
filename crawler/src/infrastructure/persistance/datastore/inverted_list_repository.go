package datastore

import (
	"fmt"
	"strings"

	"search_engine_project/crawler/src/domain/repository"
	"search_engine_project/crawler/src/domain/model/entity"
)

// InvertedListRepository ...
type InvertedListRepository struct {}

// NewInvertedListRepository ...
func NewInvertedListRepository() repository.InvertedListRepository {
	return &InvertedListRepository{}
}

// BulkInsert ...
func (r *InvertedListRepository) BulkInsert(invertedList entity.InvertedList) error {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tokenRepository := NewTokenRepository()

	for token, documentTokens := range invertedList {
		tokenID, err := tokenRepository.GetID(token)
		if err != nil {
			continue
		}
		bulkInsertSQL := "INSERT IGNORE INTO inverted_list (token_id, document_id, tf, offset_list, created_at, updated_at) VALUES "

		for documentID, documentToken := range documentTokens {
			offsetList := strings.Join(documentToken.OffsetList, ",")
			bulkInsertSQL += fmt.Sprintf("('%d', '%d', '%f', '%s', NOW(), NOW()), ", tokenID, documentID, documentToken.TF, offsetList)
		}
		bulkInsertSQL = bulkInsertSQL[:len(bulkInsertSQL)-2]

		// 登録処理
		_, err = db.Exec(bulkInsertSQL)
		if err != nil {
			return err
		}
	}

	return nil
}

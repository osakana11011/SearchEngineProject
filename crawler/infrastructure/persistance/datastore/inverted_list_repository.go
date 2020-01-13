package datastore

import (
	"fmt"
	"strings"

	"search_engine_project/crawler/domain/repository"
	"search_engine_project/crawler/domain/model/entity"
)

// InvertedListRepository ...
type InvertedListRepository struct {}

// NewInvertedListRepository ...
func NewInvertedListRepository() repository.InvertedListRepository {
	return &InvertedListRepository{}
}

// BulkInsert ...
func (r *InvertedListRepository) BulkInsert(invertedList map[string][]entity.PostingList) error {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	wordRepository := NewWordRepository()

	for word, postingList := range invertedList {
		wordID, err := wordRepository.GetID(word)
		if err != nil {
			continue
		}
		bulkInsertSQL := "INSERT IGNORE INTO inverted_list (word_id, page_id, tf, offset_list, created_at, updated_at) VALUES "
		for _, posting := range postingList {
			offsetList := strings.Join(posting.Word.OffsetList, ",")
			bulkInsertSQL += fmt.Sprintf("('%d', '%d', '%f', '%s', NOW(), NOW()), ", wordID, posting.PageID, posting.TF, offsetList)
		}
		bulkInsertSQL = bulkInsertSQL[:len(bulkInsertSQL)-2]

		// fmt.Println(word, bulkInsertSQL)
		// 登録処理
		_, err = db.Exec(bulkInsertSQL)
		if err != nil {
			return err
		}
	}

	return nil
}

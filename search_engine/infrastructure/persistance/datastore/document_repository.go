package datastore

import (
	"search_engine_project/search_engine/domain/model/entity"
	"search_engine_project/search_engine/domain/repository"
)

// DocumentRepository ...
type DocumentRepository struct {}

// NewDocumentRepository ...
func NewDocumentRepository() repository.DocumentRepository {
	return &DocumentRepository{}
}

// GetDocuments ...
func (r *DocumentRepository) GetDocuments(q string) ([]entity.Document, error) {
	// DB接続
	db, connectionErr := connectDB()
	if connectionErr != nil {
		return nil, connectionErr
	}
	defer db.Close()

	stmt, prepareErr := db.Prepare("SELECT title, url FROM documents WHERE title LIKE ? LIMIT 10")
	if prepareErr != nil {
		return nil, prepareErr
	}
	defer stmt.Close()

	rows, queryErr := stmt.Query("%" + q + "%")
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	var documents []entity.Document
	for rows.Next() {
		document := entity.Document{}
		scanErr := rows.Scan(&document.Title, &document.URL)
		if scanErr != nil {
			return nil, scanErr
		}
		documents = append(documents, document)
	}

	return documents, nil
}

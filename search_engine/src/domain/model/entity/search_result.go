package entity

import (
	"math"
	"strconv"
)

type SearchResult struct {
	Q string
    Documents []Document
    DocumentsN int
	Pages []Page
}

type Page struct {
	URL string
	Number int
	IsCurrent bool
}

func NewSearchResult(q string, documents []Document, documentsN int, page int, limit int) SearchResult {
	floatDocumentsN := float64(documentsN)
	floatLimit := float64(limit)

	pages := []Page{}
	for i := 1; i <= int(math.Min(math.Floor(floatDocumentsN / floatLimit), 10)); i++ {
		url := "/search/?q="+q+"&page="+strconv.Itoa(i)
		pages = append(pages, Page{URL: url, Number: i, IsCurrent: (page == i)})
	}

	return SearchResult{
		Q: q,
		Documents: documents,
		DocumentsN: documentsN,
		Pages: pages,
	}
}

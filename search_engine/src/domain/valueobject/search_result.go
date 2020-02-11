package valueobject

import (
    "time"
)

// SearchResult は検索結果を格納する為の構造体。
type SearchResult struct {
    documents   []Document
    isUsedCache bool
    currentPage int
    maxPage     int
    searchTime  time.Time
}

package valueobject

import (
    "strings"
    "search_engine_project/search_engine/src/util"
)

// Query は検索用の情報をまとめた構造体
type Query struct {
    QueryStrings []string
}

// NewQueryFromString は検索文字列からQuery構造体を生成して返す
func NewQueryFromString(q string) Query {
    normalizedQ := util.Normalize(q)
    queryStrings := strings.Split(normalizedQ, " ")

    return Query{QueryStrings: queryStrings}
}

package entity

import (
    "strconv"
)

// DocumentToken は文書中に出現する単語情報を管理するエンティティ。
type DocumentToken struct {
    Token              string    // 単語文字列。
    DocumentTokenCounts int       // 文書中に出現する単語の総数
    OffsetList         []string  // 文書中に出現する単語のオフセットリスト。(何番目に出現するか。)
    OffsetCounts       int       // 単語の出現数。
    TF                 float64   // 文書に対する単語のTF値。文書における単語の重要度。(= 単語の出現数/文書中の総単語数)
}

// addOffset は文書中に出現する単語のオフセットを追加する。
func (dw *DocumentToken) addOffset(offset int) {
    dw.OffsetList = append(dw.OffsetList, strconv.Itoa(offset))
    dw.OffsetCounts++
    dw.TF = (float64)(dw.OffsetCounts) / (float64)(dw.DocumentTokenCounts)
}

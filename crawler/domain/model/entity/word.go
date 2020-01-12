package entity

import (
	"strconv"
)

// Word ...
type Word struct {
	TF         float64   // 単語のTF値 (単語の出現数 / 文書全体の単語数)
	OffsetList []string  // 単語が出現するオフセットリスト
}

// AddOffset ...
func (word *Word) AddOffset(offset int) {
	word.OffsetList = append(word.OffsetList, strconv.Itoa(offset))
}

// CalcTF ...
func (word *Word) CalcTF(pageWordsCounts int) {
	word.TF = (float64)(len(word.OffsetList)) / (float64)(pageWordsCounts)
}

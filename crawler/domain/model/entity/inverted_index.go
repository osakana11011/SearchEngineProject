package entity

import (
	"errors"
)

// InvertedIndex はメモリ上に構築するミニ転置インデックス。
// メモリ上にミニ転置インデックスを構築してからDBに一気に保存させることで、
// DBとのやり取りの回数が減るため高速化に繋がる。
type InvertedIndex struct {
	DocumentCounts int           // ミニ転置インデックスを構築している文書数。
	WordDictionary []string      // 転置インデックス用の単語辞書。
	InvertedList   InvertedList  // 転置リスト。
}

// InvertedList ...
type InvertedList map[string]map[int64]DocumentWord

// invertedIndex はシングルトンパターンにしたいので、init関数でプログラム起動時に1つだけ生成を行う
var invertedIndex *InvertedIndex
func init() {
	InitInvertedIndex()
}

// InitInvertedIndex ...
func InitInvertedIndex() {
	invertedIndex = &InvertedIndex{0, []string{}, map[string]map[int64]DocumentWord{}}
}

// GetInvertedIndex はプログラム内で1つしか存在しない転置インデックスを返す
func GetInvertedIndex() *InvertedIndex {
	return invertedIndex
}

// AddDocument 文書をミニ転置インデックスに追加する。
func (x *InvertedIndex) AddDocument(documentID int64, document Document) error {
	if documentID == 0 {
		return errors.New("DocumentIDが振られていない為、ミニ転置インデックスに登録できません。")
	}

	// ミニ転置インデックスに単語を登録
	for _, word := range document.Words {
		if !containsWord(x.WordDictionary, word) {
			x.WordDictionary = append(x.WordDictionary, word)
			x.InvertedList[word] = map[int64]DocumentWord{}
		}

		x.InvertedList[word][documentID] = *document.InvertedList[word]
	}

	// 文書数をインクリメント
	x.DocumentCounts++

	return nil
}

func containsWord(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

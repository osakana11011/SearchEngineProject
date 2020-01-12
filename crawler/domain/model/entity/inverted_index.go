package entity

// InvertedIndex ...
type InvertedIndex struct {
	DocumentCounts int
	Words []string
	InvertedList map[string][]PostingList
}

// PostingList ...
type PostingList struct {
	PageID int64
	Word
}

// invertedIndex はシングルトンパターンにしたいので、init関数でプログラム起動時に1つだけ生成を行う
var invertedIndex *InvertedIndex
func init() {
	InitInvertedIndex()
}

// InitInvertedIndex ...
func InitInvertedIndex() {
	invertedIndex = &InvertedIndex{0, []string{}, map[string][]PostingList{}}
}

// GetInvertedIndex はプログラム内で1つしか存在しない転置インデックスを返す
func GetInvertedIndex() *InvertedIndex {
	return invertedIndex
}

// AddDocument ...
func (x *InvertedIndex) AddDocument(pageID int64, page Page) {
	// ミニ転置インデックスに単語を登録
	for word := range page.Words {
		if !containsWord(x.Words, word) {
			x.Words = append(x.Words, word)
			x.InvertedList[word] = []PostingList{}
		}

		// NOTE: 単語配列がユニークである前提。ユニークで無い場合は同じ単語に対して同じpageIDが重複して振られる
		x.InvertedList[word] = append(x.InvertedList[word], PostingList{pageID, *page.Words[word]})
	}

	// 文書数をインクリメント
	x.DocumentCounts++
}

func containsWord(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

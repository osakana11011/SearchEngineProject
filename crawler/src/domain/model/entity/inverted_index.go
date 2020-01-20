package entity

import (
    "errors"
)

// InvertedIndex はメモリ上に構築するミニ転置インデックス。
// メモリ上にミニ転置インデックスを構築してからDBに一気に保存させることで、
// DBとのやり取りの回数が減るため高速化に繋がる。
type InvertedIndex struct {
    DocumentCounts  int           // ミニ転置インデックスを構築している文書数。
    TokenDictionary []string      // 転置インデックス用の単語辞書。
    InvertedList    InvertedList  // 転置リスト。
}

// InvertedList は「トークン => 文書トークンmap」を取得する転置リスト
type InvertedList map[string]map[int64]DocumentToken

// invertedIndex はシングルトンパターンにしたいので、init関数でプログラム起動時に1つだけ生成を行う
var invertedIndex *InvertedIndex
func init() {
    InitInvertedIndex()
}

// InitInvertedIndex は転置初期化した転置インデックスを返す
func InitInvertedIndex() {
    invertedIndex = &InvertedIndex{0, []string{}, map[string]map[int64]DocumentToken{}}
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
    for _, token := range document.Tokens {
        if !containsToken(x.TokenDictionary, token) {
            x.TokenDictionary = append(x.TokenDictionary, token)
            x.InvertedList[token] = map[int64]DocumentToken{}
        }

        x.InvertedList[token][documentID] = *document.InvertedList[token]
    }

    // 文書数をインクリメント
    x.DocumentCounts++

    return nil
}

func containsToken(tokens []string, token string) bool {
    for _, w := range tokens {
        if w == token {
            return true
        }
    }
    return false
}

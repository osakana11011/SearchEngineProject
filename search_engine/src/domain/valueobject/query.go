package valueobject

import (
    "strings"

    "search_engine_project/search_engine/src/util"

    "github.com/bluele/mecab-golang"
)

// Query は検索条件をまとめたreadonlyな構造体。
type Query struct {
    queryString string
    tokens []string
}

// NewQuery は検索文字列qからQuery構造体を生成して返す。
func NewQuery(q string) (Query, error) {
    // TODO: 検索文字列qのバリデーション
    q = util.RemoveEmoji(q)  // 絵文字表現の削除
    q = util.Normalize(q)    // 検索文字列の正規化
    tokens, err := extractNounWords(q)
    if err != nil {
        return Query{}, err
    }

    return Query{queryString: q, tokens: tokens}, nil
}

// QueryString はqueryStringの値を取得する為のメソッド。
func (q *Query) QueryString() string {
    return q.queryString
}

// Tokens はtokensの値を取得する為のメソッド
func (q *Query) Tokens() []string {
    return q.tokens
}

// 名詞のみを抜き出す
func extractNounWords(text string) ([]string, error) {
    // 形態素解析を行う
    m, err := mecab.New("-d /usr/lib64/mecab/dic/mecab-ipadic-neologd")
    if err != nil {
        return nil, err
    }
    defer m.Destroy()

    tg, err := m.NewTagger()
    if err != nil {
        return nil, err
    }
    defer tg.Destroy()

    lt, err := m.NewLattice(text)
    if err != nil {
        return nil, err
    }
    defer lt.Destroy()

    words := []string{}
    node := tg.ParseToNode(lt)
    for {
        // 文末まで行くとnode.Next()でエラーを吐くのでそれを合図にループ終了
        if node.Next() != nil {
            break
        }

        word := node.Surface()
        features := strings.Split(node.Feature(), ",")

        if (word != "") && (features[0] == "名詞") {
            words = append(words, word)
        }
    }

    return words, nil
}

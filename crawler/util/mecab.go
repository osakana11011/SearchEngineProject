package util

import (
	"strings"

	"github.com/bluele/mecab-golang"
)

func prepareMeCab() (*mecab.MeCab, *mecab.Tagger, error) {
	m, err := mecab.New("-d /usr/lib64/mecab/dic/mecab-ipadic-neologd")
	if err != nil {
		return nil, nil, err
	}
	tg, err := m.NewTagger()
	if err != nil {
		return nil, nil, err
	}

	return m, tg, nil
}

// ExtractNounWords text中に存在する名詞単語のみを抽出して返す。
// 重複する単語があったとしても、出現した順番に返す。
func ExtractNounWords(text string) ([]string, error) {
	// 形態素解析の準備
	m, tg, err := prepareMeCab()
	if err != nil {
		return []string{}, err
	}
	defer m.Destroy()
	defer tg.Destroy()

	// 形態素解析
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

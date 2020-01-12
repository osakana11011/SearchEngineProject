package entity

import (
	"strings"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/bluele/mecab-golang"
)

// Page ...
type Page struct {
	Title string           // ページのタイトル
	URL   string           // ページのURL
	Words map[string]*Word  // 単語にポスティングリストが紐づく
	Links []string         // ページ内に存在するリンク
}

// CrawlPage ...
func CrawlPage(url string) (Page, error) {
	page := Page{}

	// URLからページ情報の取得
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return page, err
	}

	// 形態素解析して、ページ内の名詞単語を全て取得
	words, err := extractNounWords(doc)
	if err != nil {
		return page, err
	}

	// ページの情報を詰めて返す
	page.Title = doc.Find("title").Text()
	page.URL = url
	page.Words = words
	page.Links = extractLinks(doc)

	return page, nil
}

// extractNounWords ...
func extractNounWords(doc *goquery.Document) (map[string]*Word, error) {
	// 本文を抜き出す
	bodyText := doc.Find("body").Text()
	bodyText = strings.NewReplacer(
        "\r\n", "",
        "\r", "",
		"\n", "",
		"\t", "",
		" ", "",
		"'", "",
	).Replace(bodyText)

	// 形態素解析を行う行う為にMeCabの準備を行う
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

	// 形態素解析
	lt, err := m.NewLattice(bodyText)
	if err != nil {
		return nil, err
	}
	defer lt.Destroy()

	// 形態素解析の結果から、「名詞」のみ抜き出して単語リストを構築する
	var i int
	words := map[string]*Word{}
	node := tg.ParseToNode(lt)
	for i = 0; ; i++ {
		// 文末まで行くと、node.Next()で警告を吐くようになるので、それが出たらループを抜ける
		if node.Next() != nil {
			break
		}

		word := node.Surface()
		features := strings.Split(node.Feature(), ",")

		if word != "" && features[0] == "名詞" {
			_, isKeyExist := words[word]
			if !isKeyExist {
				words[word] = &Word{}
			}
			wordPointer := words[word]
			wordPointer.AddOffset(i)
		}
	}

	// 各単語についてTF値を求める
	for _, word := range words {
		word.CalcTF(i)
	}

	return words, nil
}

// extractLinks ...
func extractLinks(doc *goquery.Document) []string {
	links := []string{}

	// Wikipedia内のリンクは相対パスで書かれていることがあるので補う
	r := regexp.MustCompile(`^/wiki/*`)

	// aタグを全て見つけ出してhrefを抜き出していく
	anchorSelections := doc.Find("a")
	anchorSelections.Each(func(_ int, anchorSelection *goquery.Selection) {
		url, success := anchorSelection.Attr("href")
		if success {
			if r.MatchString(url) {
				url = "https://ja.wikipedia.org" + url
			}
			links = append(links, url)
		}
	})
	return links
}

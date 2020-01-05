package entity

import (
	"fmt"
	"time"
	"strings"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/bluele/mecab-golang"
)

// Page ...
type Page struct {
	Title     string
	URL       string
	NounWords []string
	Links     []string
}

// CrawlPage ...
func CrawlPage(url string) (Page, error) {
	page := Page{}

	// サーバに負荷を掛けすぎないように自重
	time.Sleep(1 * time.Second)

	// URLからページ情報の取得
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return page, err
	}

	// タイトル
	page.Title = doc.Find("title").Text()

	// URL
	page.URL = url

	// 形態素解析して、ページ内の単語を全て取得
	bodyText := doc.Find("body").Text()
	bodyText = strings.NewReplacer(
        "\r\n", "",
        "\r", "",
		"\n", "",
		"\t", "",
		" ", "",
	).Replace(bodyText)
	nounWords, err := extractNounWords(bodyText)
	if err != nil {
		return page, err
	}
	page.NounWords = nounWords
	fmt.Println(strings.Join(nounWords, ", "))

	// ページ内のリンクを全て取得
	links := []string{}
	anchorSelections := doc.Find("a")
	anchorSelections.Each(func(_ int, anchorSelection *goquery.Selection) {
		childURL, success := anchorSelection.Attr("href")
		if success {
			// Wikipedia内のリンクは相対パスで書かれていることがあるので補う
			r := regexp.MustCompile(`^/wiki/*`)
			if r.MatchString(childURL) {
				childURL = "https://ja.wikipedia.org" + childURL
			}

			links = append(links, childURL)
		}
	})
	page.Links = links

	return page, nil
}

// extractNounWords ...
func extractNounWords(text string) ([]string, error) {
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

	nounWords := []string{}
	node := tg.ParseToNode(lt)
	for {
		word := node.Surface()
		features := strings.Split(node.Feature(), ",")

		if features[0] == "名詞" {
			nounWords = append(nounWords, word)
		}

		if node.Next() != nil {
			break
		}
	}

	return nounWords, nil
}

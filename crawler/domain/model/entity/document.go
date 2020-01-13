package entity

import (
	"fmt"
	"regexp"
	"strconv"

	"search_engine_project/crawler/util"

	"github.com/PuerkitoBio/goquery"
)

// Document はWeb文書1つに対応するデータ構造。
type Document struct {
	Title        string                    // タイトル
	URL          string                    // URL
	Words        []string                  // 文書内に現れる単語リスト(重複無し)
	InvertedList map[string]*DocumentWord  // 文書に対する転置リスト
	Links        []string                  // 文書内に存在するリンク
}

// GetDocumentByCrawl はURL情報を与えて、それに該当するWebページの情報を返す。
func GetDocumentByCrawl(url string) (Document, error) {
	// goqueryを用いてWeb文書を取得する。
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return Document{}, err
	}

	// 本文を抜き出して、名詞単語列を取得する。
	bodyText := util.FormatString(doc.Find("body").Text())
	words, err := util.ExtractNounWords(bodyText)
	if err != nil {
		return Document{}, err
	}

	// Documentの構築
	d := Document{}
	d.Title = doc.Find("title").Text()
	d.URL = url
	d.Words = util.UniqArray(words)
	d.InvertedList = getInvertedList(words)
	d.Links = extractLinks(doc)

	fmt.Println(d.InvertedList)
	return d, nil
}

// getInvertedList は単語のリストから転置リストを構築する。
func getInvertedList(words []string) map[string]*DocumentWord {
	invertedList := map[string]*DocumentWord{}

	for offset, word := range words {
		if _, isExist := invertedList[word]; !isExist {
			invertedList[word] = &DocumentWord{Word: word, DocumentWordCounts: len(words)}
		}
		dw := invertedList[word]
		dw.addOffset(offset)
	}

	return invertedList
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




// DocumentWord は文書中に出現する単語情報を管理するエンティティ。
type DocumentWord struct {
	Word               string    // 単語文字列。
	DocumentWordCounts int       // 文書中に出現する単語の総数
	OffsetList         []string  // 文書中に出現する単語のオフセットリスト。(何番目に出現するか。)
	OffsetCounts       int       // 単語の出現数。
	TF                 float64   // 文書に対する単語のTF値。文書における単語の重要度。(= 単語の出現数/文書中の総単語数)
}

// addOffset は文書中に出現する単語のオフセットを追加する。
func (dw *DocumentWord) addOffset(offset int) {
	dw.OffsetList = append(dw.OffsetList, strconv.Itoa(offset))
	dw.OffsetCounts++
	dw.TF = (float64)(dw.OffsetCounts) / (float64)(dw.DocumentWordCounts)
}

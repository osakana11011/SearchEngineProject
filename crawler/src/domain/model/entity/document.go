package entity

import (
    "regexp"

    "search_engine_project/crawler/src/util"

    "github.com/PuerkitoBio/goquery"
)

// Document はWeb文書1つに対応するデータ構造。
type Document struct {
    Title        string                    // タイトル
    URL          string                    // URL
    Tokens        []string                  // 文書内に現れる単語リスト(重複無し)
    InvertedList map[string]*DocumentToken // 文書に対する転置リスト
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
    bodyText := doc.Find("body").Text()
    bodyText = util.RemoveEmoji(bodyText)
    bodyText = util.Normalize(bodyText)
    tokens, err := util.ExtractNounTokens(bodyText)
    if err != nil {
        return Document{}, err
    }

    // Documentの構築
    d := Document{}
    d.Title = doc.Find("title").Text()
    d.URL = url
    d.Tokens = util.UniqArray(tokens)
    d.InvertedList = getInvertedList(tokens)
    d.Links = extractLinks(doc)

    return d, nil
}

// getInvertedList は単語のリストから転置リストを構築する。
func getInvertedList(tokens []string) map[string]*DocumentToken {
    invertedList := map[string]*DocumentToken{}

    for offset, token := range tokens {
        if _, isExist := invertedList[token]; !isExist {
            invertedList[token] = &DocumentToken{Token: token, DocumentTokenCounts: len(tokens)}
        }
        dw := invertedList[token]
        dw.addOffset(offset)
    }

    return invertedList
}

// extractLinks は文書を解析してリンクリストを返す
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

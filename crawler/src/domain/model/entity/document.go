package entity

import (
    "search_engine_project/crawler/src/util"
    "github.com/PuerkitoBio/goquery"
    "github.com/jinzhu/gorm"
)

// Document は文書情報1つに対応するデータ構造
type Document struct {
    gorm.Model
    Title          string   `gorm:"type:text"`
    URL            string   `gorm:"type:varchar(2083)"`
    Description    string   `gorm:"type:text"`
    DomainID       uint     `gorm:"type:int;index"`
    Domain         Domain
    UnUniqueTokens []string `gorm:"-"`
    UniqueTokens   []string `gorm:"-"`
    ChildLinks     []CrawlWaiting `gorm:"-"`
}

// GetDocumentByCrawl はURL情報を用いてクロールしてDocumentを取得する
func GetDocumentByCrawl(url string) (Document, error) {
    gqdoc, err := goquery.NewDocument(url)
    if err != nil {
        return Document{}, err
    }

    document := Document{}
    document.Title = getTitle(gqdoc)
    document.URL = url
    document.Description = getDescription(gqdoc)
    document.Domain = GetDomainByURL(url)
    document.UnUniqueTokens, _ = getUnUniqueTokens(gqdoc)
    document.UniqueTokens, _ = getUniqueTokens(gqdoc)
    document.ChildLinks, _ = getChildLinks(gqdoc, document.Domain)

    if err != nil {
        return Document{}, err
    }

    return document, nil
}

func getTitle(doc *goquery.Document) string {
    return doc.Find("title").Text()
}

func getDescription(doc *goquery.Document) string {
    bodyText := doc.Find("body").Text()
    bodyText = util.RemoveEmoji(bodyText)
    bodyText = util.Normalize(bodyText)
    bodyText = util.CutStringData(bodyText, 500, "...")

    return bodyText
}

func getUnUniqueTokens(doc *goquery.Document) ([]string, error) {
    bodyText := doc.Find("body").Text()
    bodyText = util.RemoveEmoji(bodyText)
    bodyText = util.Normalize(bodyText)
    tokenNames, err := util.ExtractNounTokens(bodyText)

    if err != nil {
        return []string{}, err
    }

    return tokenNames, nil
}

func getUniqueTokens(doc *goquery.Document) ([]string, error) {
    bodyText := doc.Find("body").Text()
    bodyText = util.RemoveEmoji(bodyText)
    bodyText = util.Normalize(bodyText)
    tokenNames, err := util.ExtractNounTokens(bodyText)

    if err != nil {
        return []string{}, err
    }

    uniqueTokenMemo := make(map[string]string)
    uniqueTokens := []string{}
    for _, tokenName := range tokenNames {
        if _, isExist := uniqueTokenMemo[tokenName]; !isExist {
            uniqueTokenMemo[tokenName] = "isExist"
            uniqueTokens = append(uniqueTokens, tokenName)
        }
    }

    return uniqueTokens, nil
}

func getChildLinks(doc *goquery.Document, domain Domain) ([]CrawlWaiting, error) {
    bodyText := doc.Find("body").Text()
    bodyText = util.RemoveEmoji(bodyText)
    bodyText = util.Normalize(bodyText)

    // aタグを見つけて、href属性を抜き出していく
    var childLinks []CrawlWaiting
    anchorSelections := doc.Find("a")
    anchorSelections.Each(func(_ int, anchorSelection *goquery.Selection) {
        url, success := anchorSelection.Attr("href")
        url = util.NormalizeURL(url, domain.Name)
        if success && (url != "") {
            link := CrawlWaiting{URL: url, IsPriority: false}
            childLinks = append(childLinks, link)
        }
    })

    return childLinks, nil
}

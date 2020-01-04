package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"regexp"
	"net/http"
	"strings"

	"search_engine_project/crawler/database"

	"github.com/joho/godotenv"
	"github.com/PuerkitoBio/goquery"
	"github.com/bluele/mecab-golang"
	_ "github.com/go-sql-driver/mysql"
)

// FormatText ...
func FormatText(text string) string {
	text = strings.NewReplacer(
        "\r\n", "",
        "\r", "",
		"\n", "",
		"\t", "",
		" ", "",
	).Replace(text)

	return text
}

// MorphologicalAnalysis ...
func MorphologicalAnalysis(text string) error {
	m, err := mecab.New("-d /usr/lib64/mecab/dic/mecab-ipadic-neologd")
	if err != nil {
		return err
	}
	defer m.Destroy()

	tg, err := m.NewTagger()
	if err != nil {
		return err
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(text)
	if err != nil {
		return err
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		word := node.Surface()
		features := strings.Split(node.Feature(), ",")

		// まだ登録されていない単語なら、新規登録を行う
		isRegistedWord, err := database.IsRegistedWord(word)
		if err != nil {
			return err
		}
		if !isRegistedWord && features[0] == "名詞" {
			registErr := database.RegistWord(word, features[0])
			if registErr != nil {
				return registErr
			}
		}

		if node.Next() != nil {
			break
		}
	}
	return nil
}

// Crawling ...
func Crawling(url string, depth int) error {
	if depth <= 0 {
		log.Println("Stop crawling because depth is 0.")
		return nil
	}

	// 日本語版Wikipediaしか取得しないようにしたいので、ドメイン名でフィルタリングする
	r := regexp.MustCompile(`^https://ja.wikipedia.org/`)
	if !r.MatchString(url) {
		log.Println("Stop crawling because this URL's domain isn't 「ja.wikipecia.org」.")
		return nil
	}

	isRegistedPageRecently, err := database.IsRegistedPageRecently(url)
	if err != nil {
		return err
	}
	if isRegistedPageRecently {
		return nil
	}

	// リクエスト送り過ぎると目付けられるので自重する
	time.Sleep(500 * time.Millisecond)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// goqueryによるスクレイピングの準備
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	// タイトル
	titleSelection := doc.Find("title")
	title := titleSelection.Text()

	bodySelection := doc.Find("body")
	bodyText := FormatText(bodySelection.Text())

	morphologicalErr := MorphologicalAnalysis(bodyText)
	if morphologicalErr != nil {
		return morphologicalErr
	}

	// まだ登録されていないページなら、新規登録を行う
	isRegistedPage, err := database.IsRegistedPage(url)
	if err != nil {
		return err
	}
	if !isRegistedPage {
		registErr := database.RegistPage(title, url)
		if registErr != nil {
			return registErr
		}
	}

	// ページ内のリンクを全て見つけて、そのページのクローリングを行っていく
	anchorSelections := doc.Find("a")
	anchorSelections.Each(func(_ int, anchorSelection *goquery.Selection) {
		nextURL, success := anchorSelection.Attr("href")
		if success {
			// Wikipedia内のリンクは相対パスで書かれていることがあるので補う
			r := regexp.MustCompile(`^/wiki/*`)
			if r.MatchString(nextURL) {
				nextURL = "https://ja.wikipedia.org" + nextURL
			}
			Crawling(nextURL, depth - 1)
		}
	})

	return nil
}

func main() {
	// dotenvファイルを環境変数にロード
	err := godotenv.Load(fmt.Sprintf(".envfiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed load .envfiles/%s.env", os.Getenv("GO_ENV")))
	}

	// クローリング開始
	crawlingErr := Crawling("https://ja.wikipedia.org/wiki/Google", 10)
	if crawlingErr != nil {
		log.Fatal(crawlingErr)
	}
}

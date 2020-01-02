package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"regexp"
	"net/http"
	// "strings"

	"search_engine_project/crawler/database"

	"github.com/joho/godotenv"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

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

	isRegistedRecently, err := database.IsRegistedRecently(url)
	if err != nil {
		return err
	}
	if isRegistedRecently {
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

	// bodySelection := doc.Find("body")
	// bodyText := bodySelection.Text()
	// bodyText = strings.NewReplacer(
    //     "\r\n", "",
    //     "\r", "",
	// 	"\n", "",
	// 	"\t", "",
	// 	" ", "",
    // ).Replace(bodyText)
	// fmt.Println(bodyText)

	// まだ登録されていないページなら、新規登録を行う
	isRegisted, err := database.IsRegisted(url)
	if err != nil {
		return err
	}
	if !isRegisted {
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

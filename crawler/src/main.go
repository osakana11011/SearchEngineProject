package main

import (
	"os"
	"fmt"
	"search_engine_project/crawler/src/domain/model/newentity"
	"search_engine_project/crawler/src/infrastructure/persistance/datastore"
	"search_engine_project/crawler/src/usecase"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func init () {
	// 環境によってdotenvファイルを取り込む
	env := os.Getenv("ENV")
	if env != "" {
		if err := godotenv.Load(fmt.Sprintf("./.envfiles/%s.env", env)); err != nil {
			panic(err)
		}
	}

	// 自動マイグレーション
	datastore.MigrateAll()
}

func main() {
	crawlWaitingRepository := datastore.NewCrawlWaitingRepository()
	crawlWaiting := newentity.CrawlWaiting{URL: "https://ja.wikipedia.org/wiki/Google", IsPriority: true}
	crawlWaitingRepository.Insert(crawlWaiting)

	if err := usecase.ExecCrawlService(); err != nil {
		fmt.Println(err)
	}
}

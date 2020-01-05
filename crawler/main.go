package main

import (
	"os"
	"fmt"
	"log"

	"search_engine_project/crawler/domain/service"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// dotenvファイルを環境変数にロード
	err := godotenv.Load(fmt.Sprintf(".envfiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed load .envfiles/%s.env", os.Getenv("GO_ENV")))
	}

	pageService := service.NewPageService()
	err = pageService.Crawl("https://ja.wikipedia.org/wiki/%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB:Inkscape_logo_2.svg", 10)
	if err != nil {
		fmt.Println(err)
	}
}

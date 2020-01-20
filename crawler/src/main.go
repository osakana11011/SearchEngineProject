package main

import (
    "os"
    "fmt"
    "log"

    "search_engine_project/crawler/src/domain/service"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

const (
    initDepth = 15
)

func init() {
    // ログ設定
    log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
    logfile, _ := os.OpenFile("/var/www/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    log.SetOutput(logfile)

    // dotenvファイルを環境変数にロード
    err := godotenv.Load(fmt.Sprintf(".envfiles/%s.env", os.Getenv("GO_ENV")))
    if err != nil {
        log.Fatal(fmt.Sprintf("[Fatal] Failed load .envfiles/%s.env", os.Getenv("GO_ENV")))
    }
}

func main() {
    documentService := service.NewDocumentService()

    // クローリング開始
    err := documentService.Crawl("https://ja.wikipedia.org/wiki/Google", initDepth)
    if err != nil {
        log.Fatal(err)
    }
}

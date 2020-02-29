package main

import (
    "os"
    "fmt"
    "log"
    "net/http"

    "search_engine_project/search_engine/src/views"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // dotenvファイルを環境変数にロード
    err := godotenv.Load(fmt.Sprintf(".envfiles/%s.env", os.Getenv("ENV")))
    if err != nil {
        log.Fatal(fmt.Sprintf("failed load .envfiles/%s.env", os.Getenv("ENV")))
    }

    // assets 以下へのアクセス
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

    // ページへのアクセス
    http.HandleFunc("/", views.IndexHandler)
    http.HandleFunc("/search/", views.SearchHandler)
    http.HandleFunc("/console", views.ConsoleHandler)

    http.ListenAndServe(":3000", nil)
}

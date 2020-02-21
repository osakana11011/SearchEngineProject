package main

import (
    "os"
    "fmt"
    "log"
    "html/template"
    "net/http"

    "search_engine_project/search_engine/src/domain/model/entity"
    "search_engine_project/search_engine/src/usecase"
    "search_engine_project/search_engine/src/domain/service"
    "search_engine_project/search_engine/src/infrastructure/persistance/datastore"

    "github.com/joho/godotenv"
    "go.uber.org/dig"
    _ "github.com/go-sql-driver/mysql"
)

var c = dig.New()
func init() {
    c.Provide(datastore.NewGormDBConnection)
    c.Provide(datastore.NewTokenRepository)
    c.Provide(datastore.NewDocumentRepository)
    c.Provide(datastore.NewInvertedDataRepository)

    c.Provide(service.NewSearchService)
    c.Provide(usecase.NewSearchUseCase)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("assets/templates/index.html.tpl"))
    tpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")

    // SearchResult は検索結果をテンプレートに渡す構造体
    type SearchResult struct {
        Q string
        DocumentsN int
        Documents []entity.Document
    }

    // 検索して表示
    c.Invoke(func(searchUsecase usecase.SearchUseCase) {
        documents, _ := searchUsecase.Search(q)
        searchResult := SearchResult{Q: q, DocumentsN: len(documents), Documents: documents}
        fmt.Println(searchResult)

        tpl := template.Must(template.ParseFiles("assets/templates/search.html.tpl"))
        tpl.Execute(w, searchResult)
    })
}

func managementHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("assets/templates/management.html.tpl"))
    tpl.Execute(w, nil)
}

func main() {
    // dotenvファイルを環境変数にロード
    err := godotenv.Load(fmt.Sprintf(".envfiles/%s.env", os.Getenv("ENV")))
    if err != nil {
        log.Fatal(fmt.Sprintf("failed load .envfiles/%s.env", os.Getenv("ENV")))
    }

    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/search/", searchHandler)
    http.HandleFunc("/management", managementHandler)
    http.ListenAndServe(":3000", nil)
}

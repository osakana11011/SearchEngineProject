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

type SearchResult struct {
    Q string
    Documents []entity.Document
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("assets/templates/index.html.tpl"))
    tpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    c := dig.New()

    c.Provide(datastore.NewGormDBConnection)
    c.Provide(datastore.NewDocumentRepository)
    c.Provide(service.NewSearchService)
    c.Provide(usecase.NewSearchUseCase)

    // 検索して表示
    c.Invoke(func(searchUsecase usecase.SearchUseCase) {
        documents, _ := searchUsecase.Search("")
        searchResult := SearchResult{Q: "hoge", Documents: documents}

        tpl := template.Must(template.ParseFiles("assets/templates/search.html.tpl"))
        tpl.Execute(w, searchResult)
    })
}

func main() {
    // dotenvファイルを環境変数にロード
    err := godotenv.Load(fmt.Sprintf(".envfiles/%s.env", os.Getenv("ENV")))
    if err != nil {
        log.Fatal(fmt.Sprintf("failed load .envfiles/%s.env", os.Getenv("ENV")))
    }

    http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css/"))))
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/search/", searchHandler)
    http.ListenAndServe(":3000", nil)
}
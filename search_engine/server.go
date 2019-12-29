package main

import (
    "os"
    "fmt"
    "log"
    "html/template"
    "net/http"
    "search_engine_project/search_engine/database"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("templates/index.html.tpl"))
    m := "https://google.com"
    tpl.Execute(w, m)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    pages, err := database.GetPages(q)
    if err != nil {
        log.Println(err)
    }

    tpl := template.Must(template.ParseFiles("templates/search.html.tpl"))
    data := map[string]interface{}{"q": q, "pages": pages}
    tpl.Execute(w, data)
}

func main() {
    // dotenvファイルを環境変数にロード
	err := godotenv.Load(fmt.Sprintf(".envfiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed load .envfiles/%s.env", os.Getenv("GO_ENV")))
    }

    http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css/"))))
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/search/", searchHandler)
    http.ListenAndServe(":3000", nil)
}

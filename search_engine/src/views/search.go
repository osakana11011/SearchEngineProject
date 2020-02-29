package views

import (
	"strconv"
	"net/http"
	"html/template"

	"search_engine_project/search_engine/src/usecase"
)

// SearchHandler は検索を行い結果を返す関数
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	c := getNewContainer()

    q := r.URL.Query().Get("q")
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page == 0 {
        page = 1
    }

    // 検索して表示
    c.Invoke(func(searchUsecase usecase.SearchUseCase) {
        searchResult, _ := searchUsecase.Search(q, page)

        tpl := template.Must(template.ParseFiles("assets/templates/search.html.tpl"))
        tpl.Execute(w, searchResult)
	})
}

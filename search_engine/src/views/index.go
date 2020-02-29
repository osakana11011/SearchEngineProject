package views

import (
	"net/http"
	"html/template"
)

// IndexHandler は検索TOPページを表示する関数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("assets/templates/index.html.tpl"))
    tpl.Execute(w, nil)
}

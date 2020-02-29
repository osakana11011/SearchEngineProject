package views

import (
	"net/http"
	"html/template"
)

// ConsoleHandler はコンソール画面を表示する関数
func ConsoleHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("assets/templates/console.html.tpl"))
    tpl.Execute(w, nil)
}

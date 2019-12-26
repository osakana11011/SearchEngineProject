package main

import (
    "html/template"
    "net/http"
    // "net/url"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("templates/index.html.tpl"))
    m := "https://google.com"
    tpl.Execute(w, m)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    v := r.URL.Query()
    tpl := template.Must(template.ParseFiles("templates/search.html.tpl"))
    tpl.Execute(w, v)
}

func main() {
    http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css/"))))
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/search/", searchHandler)
    http.ListenAndServe(":3000", nil)
}

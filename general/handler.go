package general

import (
	"html/template"
	"net/http"
)

var Mux *http.ServeMux

func init() {
	Mux = http.NewServeMux()
	Mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	Mux.HandleFunc("/test", testHandler)
	Mux.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("tmpl/index.tmpl")).Execute(w, nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("tmpl/test.tmpl")).Execute(w, nil)
}

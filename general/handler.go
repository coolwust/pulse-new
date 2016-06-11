package general

import (
	"html/template"
	"net/http"
)

var Mux *http.ServeMux

func init() {
	Mux = http.NewServeMux()
	Mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	Mux.HandleFunc("/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("tmpl/index.tmpl")).Execute(w, nil)
}

package main

import (
	"net/http"
	"golang.org/x/net/http2"
	"github.com/coldume/pulse/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/", handler.IndexHandler)
	serv := &http.Server{Addr: ":80", Handler: mux}
	http2.ConfigureServer(serv, &http2.Server{})
	serv.ListenAndServe()
}

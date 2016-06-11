package main

import (
	"net/http"

	"github.com/coldume/pulse/general"
	"github.com/coldume/pulse/signup"
	"golang.org/x/net/http2"
)

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/sign-up/api/resolve/", signup.ResolveHandler)
	mux.Handle("/api/sign-up/", signup.Mux)
	mux.Handle("/", general.Mux)
	serv := &http.Server{Addr: ":80", Handler: mux}
	http2.ConfigureServer(serv, &http2.Server{})
	if err := serv.ListenAndServe(); err != nil {
		panic(err)
	}
}

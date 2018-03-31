package main

import (
	"net/http"
	"time"
)

func isIndex(url string) bool {
	if url[len(url)-1] != '/' {
		return true
	}
	return false
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if isIndex(r.URL.Path) {
		http.ServeFile(w, r, r.URL.Path[1:])
	} else {
		http.NotFound(w, r)
	}
}

var server = &http.Server{
	Addr:           ":8081",
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

func main() {
	http.HandleFunc("/asset/", FileHandler)
	panic(server.ListenAndServe())
}

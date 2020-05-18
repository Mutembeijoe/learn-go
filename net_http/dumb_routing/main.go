package main

import (
	"io"
	"log"
	"net/http"
)

type customHandler int

func (c customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	switch urlPath {
	case "/main":
		io.WriteString(w, "Hello World, main Path")
	case "/other":
		io.WriteString(w, "Hello world, other path")

	default:
		io.WriteString(w, "404 - Not Found\n This page is a wild card")
	}

}

func main() {
	var d customHandler

	log.Fatal(http.ListenAndServe(":8080", d))
}

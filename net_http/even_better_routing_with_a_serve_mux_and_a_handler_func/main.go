package main

import (
	"net/http"
)

func handlerFunc1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This response is written from handler fucn 1"))
}

func handlerFunc2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This response is written from handler fucn 2"))
}

func main() {
	mux := http.NewServeMux() //a default serve-mux would work here too

	mux.HandleFunc("/route1", handlerFunc1)
	mux.HandleFunc("/route2", handlerFunc2)

	http.ListenAndServe(":8080", mux)
}

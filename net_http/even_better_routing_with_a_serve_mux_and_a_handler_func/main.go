package main

import (
	"net/http"
)

func handlerFunc1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This response is written from handler func 1"))
}

func handlerFunc2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This response is written from handler func 2"))
}

func main() {
	mux := http.NewServeMux() //a default serve-mux would work here too

	// A http.HandlerFunc converts any function with the signature
	// func(w http.ResponseWrite, r *http.Request){}
	// to a handlerFunc which implements the ServeHttp method thus making it a handler.
	//Notice how the mux.Handle vs mux.HandleFunc are used below

	mux.Handle("/route1", http.HandlerFunc(handlerFunc1))
	mux.HandleFunc("/route2", handlerFunc2)

	http.ListenAndServe(":8080", mux)
}

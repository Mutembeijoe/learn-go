package main

import (
	"fmt"
	"net/http"
)

type customHandler1 int
type customHandler2 string //can be of any type, doesn't matter

func (ch customHandler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This router is handled by custom handler 1")
}
func (ch customHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This router is handled by custom handler 2")
}

func main() {
	var cH1 customHandler1
	var cH2 customHandler2

	http.Handle("/route1/", cH1) // NB. In /route1/ the trailing catch will catch all routes as long as they start with /route1
	http.Handle("/route2", cH2)  //Using the defaulr serve mux, a custom serve mux would work here too!

	http.ListenAndServe(":8080", nil)

	//Visting a none existing route will automatically trigger a 404 response from the server
}

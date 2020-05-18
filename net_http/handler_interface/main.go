package main

import (
	"fmt"
	"net/http"
)

type customHandler int

func (c customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World, This a custom Handler")
}

func main() {

	var cH customHandler

	http.ListenAndServe(":8080", cH)

}

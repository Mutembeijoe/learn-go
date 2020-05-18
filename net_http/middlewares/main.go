package main

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I get called before the request hits the main handler")
		handler.ServeHTTP(w, r)
		fmt.Println("I get called after the main handler dispaches the response")
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("Executing main handler....")
	w.Write([]byte("This is the response"))
}

func main() {
	http.Handle("/route", middleware(http.HandlerFunc(mainHandler)))
	http.ListenAndServe(":8080", nil)
}

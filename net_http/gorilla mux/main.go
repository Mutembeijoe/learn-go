package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\t Book category : %v \n\t Book id : %v", vars["category"], vars["id"])
}

func main() {
	gMux := mux.NewRouter()

	gMux.HandleFunc("/api/books/{category}/{id:[0-9]+}", getBook)

	s := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second, // Good practice: enforce timeouts for servers you create!
		ReadTimeout:  10 * time.Second,
		Handler:      gMux,
	}

	log.Fatal(s.ListenAndServe())
}

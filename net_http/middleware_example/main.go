package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area int64
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported media type"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func addServerTimeToResponse(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}

		http.SetCookie(w, &cookie)
	})
}

func cityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var myCity city
		decorder := json.NewDecoder(r.Body)
		err := decorder.Decode(&myCity)

		if err != nil {
			panic(err)
		}

		r.Body.Close()

		fmt.Printf(" Name : %v \n Area: %v", myCity.Name, myCity.Area)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 -Method Not allowed"))
	}
}

func main() {

	http.Handle("/route", addServerTimeToResponse(filterContentType(http.HandlerFunc(cityHandler))))

	s := http.Server{
		Addr:         ":8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      nil,
	}

	s.ListenAndServe()
}

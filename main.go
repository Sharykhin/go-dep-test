package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/me", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", r))

}

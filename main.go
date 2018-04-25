package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type User struct {
	Email string `valid:"email"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := govalidator.ValidateStruct(User{
			Email: "c@c",
		})
		fmt.Println(err)
		w.Write([]byte("OK"))
	}).Methods("GET")
	fmt.Println("Listening on localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

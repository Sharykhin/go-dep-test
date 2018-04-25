package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	db, err := sql.Open("mysql", "test:test@tcp(cr-mysql-grpc-server:3306)/test?parseTime=true")
	if err != nil {
		log.Fatalf("Could not to open a connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping a database: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", r))

}

package main

import (
	"database/sql"
	"fmt"
	"golang-rest-api/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var repo repository.PostRepository

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./db/postDB.db")
	if err != nil {
		panic("Failed to open the database: " + err.Error())
	}

	if err = db.Ping(); err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
}

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	repo = repository.NewPostRepository(db)

	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

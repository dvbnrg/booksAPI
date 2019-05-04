package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book is the struct containing the info
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publishdate"`
	Rating      int    `json:"rating"`
	Status      bool   `json:"status"`
}

var err error

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetBooks)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book

	json.NewEncoder(w).Encode(&books)
}

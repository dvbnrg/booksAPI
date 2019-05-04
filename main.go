package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Book is the struct containing the info
type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publishdate"`
	Rating      int    `json:"rating"`
	Status      bool   `json:"status"`
}

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open(
		"db",
		"host="+os.Getenv("HOST")+"user="+os.Getenv("USER")+
			"dbname="+os.Getenv("NAME")+" sslmode=disable password="+os.Getenv("PASSWORD"))

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&Book{})
	router.HandleFunc("/", GetBooks)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(&books)
}

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
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

// var tmpl = template.Must(template.ParseGlob("form/*"))
// var err error

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", GetBooks)
	// router.HandleFunc("/")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// GetBooks retrieves a complete list of books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	var book Book

	db := dbConn()

	results, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}

	for results.Next() {
		var title, author, publisher, publishdate string
		var rating int
		var checkin bool
		err = results.Scan(&title, &author, &publisher, &publishdate, &rating, &checkin)
		if err != nil {
			panic(err)
		}

		book.Title = title
		book.Author = author
		book.Publisher = publisher
		book.PublishDate = publishdate
		book.Rating = rating
		book.Status = checkin
		books = append(books, book)
	}

	// tmpl.ExecuteTemplate(w, "GetBooks", results)
	json.NewEncoder(w).Encode(&books)
	defer db.Close()

}

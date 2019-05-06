package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

// Book is the struct containing the info
type Book struct {
	ID          int    `json:"id"`
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
	router.HandleFunc("/getbooks", GetBooks)
	router.HandleFunc("/getbook/{id}", GetBook)
	router.HandleFunc("/updatebook", UpdateBook)
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
		w.WriteHeader(404)
	}

	for results.Next() {
		var title, author, publisher, publishdate string
		var id, rating int
		var checkin bool
		err = results.Scan(&id, &title, &author, &publisher, &publishdate, &rating, &checkin)
		if err != nil {
			panic(err)
			w.WriteHeader(404)
		}

		book.ID = id
		book.Title = title
		book.Author = author
		book.Publisher = publisher
		book.PublishDate = publishdate
		book.Rating = rating
		book.Status = checkin
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(&books)
	defer db.Close()
	w.WriteHeader(200)
}

// GetBook grabs the book that
func GetBook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	id := mux.Vars(r)

	results, err := db.Query("SELECT * FROM books WHERE id=?", id["id"])
	if err != nil {
		panic(err)
		w.WriteHeader(404)
	}

	book := Book{}

	for results.Next() {
		var title, author, publisher, publishdate string
		var id, rating int
		var checkin bool
		err := results.Scan(&id, &title, &author, &publisher, &publishdate, &rating, &checkin)
		if err != nil {
			panic(err)
			w.WriteHeader(404)
		}
		book.ID = id
		book.Title = title
		book.Author = author
		book.Publisher = publisher
		book.PublishDate = publishdate
		book.Rating = rating
		book.Status = checkin
	}
	json.NewEncoder(w).Encode(&book)
	defer db.Close()
	w.WriteHeader(200)
}

// UpdateBook takes in a json payload and will overwrite any existing entry
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	if r.Method == "POST" {

		read, err := r.GetBody()

		if err != nil {
			panic(err)
			w.WriteHeader(418)
		}

		body, err := ioutil.ReadAll(read)

		book := Book{}

		err = json.Unmarshal(body, &book)

		query := fmt.Sprintf("INSERT INTO books VALUES (%d,%s,%s,%s,%s,%d,%t);", book.ID, book.Title, book.Author, book.Publisher, book.PublishDate, book.Rating, book.Status)

		results, err := db.Prepare(query)
		if err != nil {
			panic(err)
			w.WriteHeader(404)
		}
		results.Exec()
		log.Printf("INSERT for %v Successful!", book)
	}

	defer db.Close()
	w.WriteHeader(200)

}

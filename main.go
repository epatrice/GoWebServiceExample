package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const maxId = 1000   // Max Unique Id (0 - maxId-1) to generate for BookId
const port = ":8080" // Port number to start this web service

// Model for a book
// Notice the json tag. Go Json package use this tag name to create JSON keys
// (e.g the JSON obj keys are bookId, bookName, price and author)
type Book struct {
	BookId    int     `json:"bookId"`
	BookName  string  `json:"bookName"`
	BookPrice int     `json:"price"`
	Author    *Author `json:"author"`
}

// Model for Author
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB. Storing all the "DB" data in in-memory struct
var books []Book

func (b *Book) IsEmpty() bool {
	return b.BookName == ""
}

func main() {
	fmt.Println("Starting example web service")
	r := mux.NewRouter()

	// Insert data in the "fake db"
	books = append(books,
		Book{BookId: 1, BookName: "Intro to C++", BookPrice: 50,
			Author: &Author{Fullname: "Bjarne Stroustrup", Website: "https://fakeCPPAuthor.com"}})

	books = append(books,
		Book{BookId: 2, BookName: "Intro to Go", BookPrice: 20,
			Author: &Author{Fullname: "Rob Pike", Website: "https://fakeGoAuthor.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/books", getAllBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	// listen to the port specified in the constant above
	log.Fatal(http.ListenAndServe(port, r))

}

// controllers

// serve home route ("/")
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my Go Web Service</h1>"))
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllBooks called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getOneBook called")
	w.Header().Set("Content-Type", "application/json")

	// get id from request url
	params := mux.Vars(r)
	idReceived, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Printf("ID received from request param cannot be converted to an integer. The Id passed in is %s", params["id"])
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID received from request param cannot be converted to an integer")
		return
	}

	// find matching id and return the response
	for _, book := range books {
		if book.BookId == idReceived {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	// still return a response if no book is found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No book found with this given id")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createOneBook called")
	w.Header().Set("Content-Type", "application/json")

	// Handle when no body was sent in the request
	if r.Body == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Empty request body received. Body required")
	}

	// Handle the case when request is empty json
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	if book.IsEmpty() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Empty JSON in request body received.")
		return
	}

	// Generate unique Id for the book
	rand.Seed(time.Now().UnixNano())
	book.BookId = rand.Intn(maxId)

	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateBook called")
	w.Header().Set("Content-Type", "application/json")

	// get id from request url
	params := mux.Vars(r)
	idReceived, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Printf("ID received from request param cannot be converted to an integer. The Id passed in is %s", params["id"])
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID received from request param cannot be converted to an integer")
		return
	}

	// search for the book with this id. If exist, first remove the book with the id and add a new book from the request
	for index, book := range books {
		if book.BookId == idReceived {
			// delete the current book
			books = append(books[:index], books[index+1:]...)
			var bookFromRequest Book
			_ = json.NewDecoder(r.Body).Decode(&bookFromRequest)
			bookFromRequest.BookId = idReceived
			books = append(books, bookFromRequest)
			json.NewEncoder(w).Encode(bookFromRequest)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Book Not Found")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteBook called")
	w.Header().Set("Content-Type", "application/json")

	// get id from request url
	params := mux.Vars(r)
	idReceived, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Printf("ID received from request param cannot be converted to an integer. The Id passed in is %s", params["id"])
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID received from request param cannot be converted to an integer")
		return
	}

	// search for the book with this id. If exist, first remove the book with the id and add a new book from the request
	for index, book := range books {
		if book.BookId == idReceived {
			// delete the current book
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode("Delete successful")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Book Not Found")
}

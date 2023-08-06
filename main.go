package webServiceExample

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const maxId = 1000

// Model for a book
// Notice the json tag. Go Json package use this tag name to create JSON keys (e.g one of the JSON key is bookId)
type Book struct {
	BookId int `json:"bookId"`  
	BookName string `json:"bookName"`
	BookPrice int `json:"price"`   
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`   
}

// fake DB. Storing all the "DB" data in in-memory struct
var books []Book

func(b *Book) IsEmpty() bool {
	return b.BookName == ""
}

func main(){

}


// controllers 

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to my Go Web Service</h1>"))
}

func getAllBooks(w http.ResponseWriter, r *http.Request){
	fmt.Println("getAllBooks called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getOneBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("getOneBook called")
	w.Header().Set("Content-Type", "application/json")

	// get id from request url
	params := mux.Vars(r)

	// find matching id and return the response
	for _, book := range books{

		idReceived, err := strconv.Atoi(params["id"])
		if(err != nil){
			fmt.Printf("ID received from request param cannot be converted to an integer. The Id passed in is %s", params["id"]))
			json.NewEncoder(w).Encode("ID received from request param cannot be converted to an integer")
	        return
		}

		if book.BookId == idReceived{
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	// still return a response if no book is found
	json.NewEncoder(w).Encode("No book found with this given id")
	return
}

func createOneBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("createOneBook called")
	w.Header().Set("Content-Type", "application/json")

	// Handle when no body was sent in the request
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty request body received. Body required")
	}

	// Request is empty json

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	if book.IsEmpty(){
		json.NewEncoder(w).Encode("Empty JSON in request body received.")
		return
	}

	// Generate unique Id for the book
	rand.Seed(time.Now().UnixNano())
	book.BookId = rand.Intn(maxId)

	books = append(books, book)
	json.NewEncoder(w).Encode("book")
	return
}
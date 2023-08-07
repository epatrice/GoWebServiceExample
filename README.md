# GoWebServiceExample

## Intro
This is an example of creating Rest API web service using Golang. This code was used by me to present how to set up a Restful web service using Go. This is intended to be a simple web service. The DB is mocked by a struct.

This uses Gorilla Mux. See [Gorilla Mux go.dev docs](https://pkg.go.dev/github.com/gorilla/mux ) and [Gorilla Mux github page](https://github.com/gorilla/mux)


## Setup
You need to first run 
`go get -u github.com/gorilla/mux` to pull the Gorilla Mux code. It is reflected in the `go.mod` and `go.sum` files.

## Starting the server
You can build and start the code with this command
`go run main.go`

You should be able to view the welcome page on `http://localhost:8080`

## APIs
Each API is tested with Postman. Please see the screenshots for each API below.

HTTP Method: Get  
Endpoint: `http://localhost:8080`  
Status Code: 200  
Response: Welcome string returned  

See screenshot:
![alt text](/screenshots/getBooksMainPage.png)


HTTP Method: Get  
Endpoint: `http://localhost:8080/books`  
Status Code: 200  
Response: All books in the "db" returned in the response as JSON  

See screenshot:
![alt text](/screenshots/getAllBooks.png)


HTTP Method: Get  
Endpoint: `http://localhost:8080/book/{id}`  
Status Code: 200 when book with id found. 404 when book not found.
Response: Book info with the specified id returned in the response as JSON when book is found. A string saying book is not found is returned when book is not found.

See screenshots:
Success  
![alt text](/screenshots/getBookId1.png)  


Not found  
![alt text](/screenshots/getBookIdNotExist.png) 


HTTP Method: Post  
Endpoint: `http://localhost:8080/book`  
Request Body: JSON object of the book. BookId is not required. The server generates a random book ID. Example body:
```
{
    "bookName": "Java",
    "price" : 100,
    "author" : {
        "fullname": "Java Author Name",
        "website": "htts://fakeUrl123.com"
    }
}
```
Status Code: 201 when book is created (having the correct request body). 404 when request body is incorrect.  
Response: Book info for the book created is returned in the response as JSON.

See screenshots:  
Success  
![alt text](/screenshots/CreateBook.png)  


Below is the result of getting all the books after creating the book using the example request body above.
![alt text](/screenshots/getAllBooksAfterCreate.png)


HTTP Method: Put  
Endpoint: `http://localhost:8080/book/{id}`  
Request Body: JSON object of the book. BookId is required in the URL only. It is not required in the request body.  

Example body:
```
{
    "bookName": "Go v2",
    "price" : 80,
    "author" : {
        "fullname": "Go Author Name",
        "website": "htts://fakeUrl456.com"
    }
}
```
Status Code: 200 when book is updated (having the correct request body). 404 when request body is incorrect or book Id doesn't exist.  
Response: Book info for the book created is returned in the response as JSON when success, else "Book Not Found".


See screenshots:  
Success
![alt text](/screenshots/updateBook.png)  

Not found  
![alt text](/screenshots/UpdateBookNotFound.png)  


HTTP Method: Delete    
Endpoint: `http://localhost:8080/book/{id}`  
Status Code: 200 when book with id found and deleted. 404 when book not found.  

See screenshots: 
Success  
![alt text](/screenshots/deleteBookSuccess.png) 

Not found  
![alt text](/screenshots/deleteBookNotFound.png)  


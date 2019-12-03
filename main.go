package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)


// book struct we pass around the program
// and marshall/unmarshal into json
type book struct {
	Name string `json:"BookName"`
	Author string `json:"BookAuthor"`
	Id string `json:"BookId"`
	Publisher string `json:"Publisher"`
	IsIssued bool `json:"IsIssued"`
	IssuedTo string `json:"IssuedTo"`
	TimesIssued int32 `json:"TimesIssued"`
	PopularityRank int32
}

// you could read this from a database such as mySQL
type allBooks []book
var books = allBooks{
	{
		Name: "Harry Potter and the Sorcerer's Stone",
		Author:"J.K. Rowling",
		Id:"000-000-000",
		Publisher:"Aditya's Amazing Publishers",
		IsIssued:false,
		TimesIssued:1,
		PopularityRank:51,
	}, {
		Name:"The Adventures of Aditya Prerepa",
		Author:"Aditya Prerepa",
		Id:"000-000-001",
		Publisher:"Aditya's Splendid Publishers",
		IsIssued:true,
		TimesIssued:32,
		IssuedTo:"John Doe",
		PopularityRank:21,
	}, {
		Name:"The Horrors of a MSJHS student",
		Author:"Corey Brown",
		Id:"000-000-002",
		Publisher:"We Hate Publishing Inc.",
		IsIssued:true,
		IssuedTo:"Jane Doe",
		TimesIssued:31,
		PopularityRank:100,
	},
}

// home page
func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome! Library assistant here. Available Routes:\nGET /books : get all books\nGET /books/{id} : get a book by and id" +
		"\nGET /bookNames/{name} : get book data identified by name (separate name with spaces)\nPOST /books : insert a book into the book database" +
		"\nDELETE /books : Delete a book from the database\nGET /mostPopular : get the most popular book right now\nGET /mostIssued : get the most issued book so far.")
}

func getMostPopularBook(w http.ResponseWriter, r *http.Request) {
	//eventID := mux.Vars(r)["id"]
	var highest int32
	var highestId string
	fmt.Println("GetMostPopular Called.")
	// find highest id
	for _, book := range books {
		if book.PopularityRank > highest {
			highest = book.PopularityRank
			highestId = book.Id
		}
	}
	fmt.Println(highest)
	for _, book := range books {
		if book.Id == highestId {
			_ = json.NewEncoder(w).Encode(book)
		}
	}
}

// get a book given the book ID
func getBookFromId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBookFromId called")
	bookId := mux.Vars(r)["id"]
	fmt.Printf("Book Id: %s\n", bookId)
	for _, book := range books {
		if book.Id == bookId {
			_ = json.NewEncoder(w).Encode(book)
		}
	}
}

// get a book given the name
func getBookFromName(w http.ResponseWriter, r *http.Request) {
	bookName := mux.Vars(r)["name"]
	for _, book := range books {
		if book.Name == bookName {
			_ = json.NewEncoder(w).Encode(book)
		}
	}
}

func getMostIssuedBook(w http.ResponseWriter, r *http.Request) {
	var mostIssuedId string
	var mostIssued int32 = 0
	var bookMostIssued book
	for _, book := range books {
		if book.TimesIssued > mostIssued {
			mostIssued = book.TimesIssued
			mostIssuedId = book.Id
		}
	}
	for _, book := range books {
		if book.Id == mostIssuedId {
			bookMostIssued = book
		}
	}
	_ = json.NewEncoder(w).Encode(bookMostIssued)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func postEvent(w http.ResponseWriter, r *http.Request) {
	var newBook book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please do it right.")
	}

	json.Unmarshal(reqBody, &newBook)
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["BookId"]

	for i, book := range books {
		if book.Id == bookID {
			books = append(books[:i], books[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", bookID)
		}
	}
}

func main() {
	fmt.Println("Starting Server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/books", getAllBooks)
	router.HandleFunc("/books/{id}", getBookFromId).Methods("GET")
	router.HandleFunc("/bookNames/{name}", getBookFromName).Methods("GET")
	router.HandleFunc("/books", postEvent).Methods("POST")
	router.HandleFunc("/books", deleteBook).Methods("DELETE")
	router.HandleFunc("/mostPopular", getMostPopularBook).Methods("GET")
	router.HandleFunc("/mostIssued", getMostIssuedBook).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
# Go-Gorilla-Rest-API
A Rest API for librarians to help with books. GCI 2020.

## Install needed Libraries
Install Gorilla Mux using:
`go get -u github.com/gorilla/mux`

## Running the program
To run the program, clone this repository, and run `go build` in the directory of `main.go`. This will produce `Go-Rest-Api`, an executable, which you then run with `./Go-Rest-Api`.

## Endpoints
- `GET /books : get all books`
- `GET /books/{id} : get a book by and id`
- `GET /bookNames/{name} : get book data identified by name (separate name with spaces)`
- `POST /books : insert a book into the book database`
- `DELETE /books : Delete a book from the database`
- `GET /mostPopular : get the most popular book right now`
- `GET /mostIssued : get the most issued book so far.`

## Room For Improvement
Right now it uses a dynamically-allocated golang slice to store books, it might be smart to move this to a persistent datastore such as MySQL in the future.

## Author
Aditya Prerepa. GCI 2020.

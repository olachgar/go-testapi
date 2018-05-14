package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	//"encoding/json"
)

// the main function
func main() {

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/books/", ReadBooks).Methods("GET")
	router.HandleFunc("/books/{id}", ReadBook).Methods("GET")
	router.HandleFunc("/books/{id}", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	// 
	log.Fatal(http.ListenAndServe(":80", router))
}

// Read Books
func ReadBooks(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "You've requested to Read all books")
}

// Read Book
func ReadBook(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested to Read the book: %s \n", id)
}

// Create Book
func CreateBook(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested to Create the book : %s \n", id)
}


// Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested to Update the book : %s \n", id)
}


// Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested to delete the book : %s \n", id)
}

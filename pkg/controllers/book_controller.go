package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aaryanraj/ebook-store-api/pkg/clients"
	"github.com/aaryanraj/ebook-store-api/pkg/models"
	"github.com/aaryanraj/ebook-store-api/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func ListBooks(w http.ResponseWriter, r *http.Request) {
	books := clients.GetAllBooks()
	resp, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	books, db := clients.GetBookByID(id)
	if db.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		response := models.UserResponse{Status_Code: http.StatusNotFound,
			Message:   "Book Not Found",
			More_info: "Try to list the books and see if the book id you are trying to find is in the list"}

		resp, _ := json.Marshal(response)
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &clients.Book{}
	if err := utils.ParseBody(r, createBook); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book := createBook.CreateBook()
	resp, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, db := clients.DeleteBook(id)
	if db.RowsAffected < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		response := models.UserResponse{Status_Code: http.StatusNotFound,
			Message:   "Book Not Found",
			More_info: "Try to list the books and see if the book id you are trying to find is in the list"}

		resp, _ := json.Marshal(response)
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	updateBook := &clients.Book{}
	utils.ParseBody(r, updateBook)
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book, db := clients.GetBookByID(id)

	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	db.Save(book)
	resp, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resp)
}

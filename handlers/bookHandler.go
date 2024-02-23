package handlers

import (
	"encoding/json"
	"fmt"
	"grpc-go2/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("error : %v", err)
	}
	var book models.Book
	json.Unmarshal(body, &book)

	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	var updatedBook, book models.Book
	json.Unmarshal(body, &updatedBook)

	if result := h.DB.Model(&book).Where("id = ?", id).Updates(updatedBook); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Updated")
}

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	if result := h.DB.Delete(&book); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Deleted")
}

package main

import (
	"grpc-go2/database"
	"grpc-go2/handlers"
	"grpc-go2/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	DB := database.Init()
	handler := handlers.New(DB)
	router := mux.NewRouter()
	router.Use(middlewares.LoggingMiddleware)
	router.HandleFunc("/books", handler.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books", handler.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", handler.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}

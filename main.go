package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/db"
	"github.com/tutorials/go/crud/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/movies", h.GetAllMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", h.GetMovie).Methods(http.MethodGet)
	router.HandleFunc("/movies", h.AddMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", h.UpdateMovie).Methods(http.MethodPut)
	router.HandleFunc("/movies/{id}", h.DeleteMovie).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

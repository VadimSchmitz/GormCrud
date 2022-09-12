package main

import (
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/cli"
	"github.com/VadimSchmitz/GormCrud/pkg/db"
	"github.com/VadimSchmitz/GormCrud/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/movies", h.GetAllMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", h.GetMovie).Methods(http.MethodGet)
	router.HandleFunc("/titles", h.GetAllMovieTitles).Methods(http.MethodGet)

	router.HandleFunc("/movies", h.AddMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", h.UpdateMovie).Methods(http.MethodPut)
	router.HandleFunc("/movies/{id}", h.DeleteMovie).Methods(http.MethodDelete)

    done := make(chan bool)

	go func() {
		http.ListenAndServe(":8090", router)
        done <- false
	}()

    go cli.CliHandler(done)
    <-done
}

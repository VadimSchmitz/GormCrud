package main

import (
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/cli"
	"github.com/VadimSchmitz/GormCrud/pkg/db"
	handlers "github.com/VadimSchmitz/GormCrud/pkg/handlers/routes"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/movies", h.GetAllMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", h.GetMovie).Methods(http.MethodGet)
	router.HandleFunc("/ids", h.GetAllMovieTitles).Methods(http.MethodGet)
	
	router.HandleFunc("/movieid", h.AddMovieWithId).Methods(http.MethodPost)
	router.HandleFunc("/movies", h.AddMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", h.UpdateMovie).Methods(http.MethodPatch)
	router.HandleFunc("/movies/{id}", h.DeleteMovie).Methods(http.MethodDelete)

    done := make(chan bool)

	go func() {
		http.ListenAndServe("127.0.0.1:8090", router)
        done <- false
	}()

    go cli.CliHandler(done)
    <-done
}

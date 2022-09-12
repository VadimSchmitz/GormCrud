package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedMovie models.Movie
	json.Unmarshal(body, &updatedMovie)

	var movie models.Movie

	if result := h.DB.Where("imdb_id = ?", id).Find(&movie); result.Error != nil {
		fmt.Println(result.Error)
	}

	movie.IMDb_id = updatedMovie.IMDb_id
	movie.Title = updatedMovie.Title
	movie.Rating = updatedMovie.Rating
	movie.Year = updatedMovie.Year

	h.DB.Save(&movie)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
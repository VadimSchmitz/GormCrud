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

	if updatedMovie.IMDb_id != "" {
		movie.IMDb_id = updatedMovie.IMDb_id
	}
	if updatedMovie.Title != "" {
		movie.Title = updatedMovie.Title
	}
	if updatedMovie.Year != 0 {
		movie.Year = updatedMovie.Year
	}
	if updatedMovie.Rating != 0 {
		movie.Rating = updatedMovie.Rating
	}
	if updatedMovie.Summary != "" {
		movie.Summary = updatedMovie.Summary
	}

	h.DB.Model(&movie).Updates(models.Movie{
		IMDb_id: movie.IMDb_id,
		Title: movie.Title, 
		Year: movie.Year, 
		Rating: movie.Rating, 
		Summary: movie.Summary})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

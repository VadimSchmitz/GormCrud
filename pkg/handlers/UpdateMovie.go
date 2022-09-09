package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

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

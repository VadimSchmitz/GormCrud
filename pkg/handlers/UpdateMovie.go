package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedMovie models.Movie
	json.Unmarshal(body, &updatedMovie)

	var movie models.Movie

	if result := h.DB.First(&movie, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	movie.Imdb_id = updatedMovie.Imdb_id
	movie.Title = updatedMovie.Title
	movie.Rating = updatedMovie.Rating
	movie.Year = updatedMovie.Year

	h.DB.Save(&movie)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/models"
)
func (h handler) GetMovie(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]
	// Find movie by Id
	var movie models.Movie
	
	if result := h.DB.Where("imdb_id = ?", id).Find(&movie); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}
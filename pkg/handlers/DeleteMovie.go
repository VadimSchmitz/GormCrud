package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/VadimSchmitz/GormCrud/pkg/models"
)

func (h handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Find the movie by Id

	var movie models.Movie
	
	if result := h.DB.Where("imdb_id = ?", id).Find(&movie); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	h.DB.Delete(&movie)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

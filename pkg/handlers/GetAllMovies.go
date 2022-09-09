package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	if result := h.DB.Find(&movies); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

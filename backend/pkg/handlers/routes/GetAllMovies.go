package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
)

func (h handler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	if result := h.DB.Find(&movies); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

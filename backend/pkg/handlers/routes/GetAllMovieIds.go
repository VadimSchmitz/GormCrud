package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
)

func (h handler) GetAllMovieIds(w http.ResponseWriter, r *http.Request) {
	var movie *models.Movie
	var IMDb_id []models.IMDb_id

	if result := h.DB.Model(&movie).Find(&IMDb_id); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(IMDb_id)
}
	
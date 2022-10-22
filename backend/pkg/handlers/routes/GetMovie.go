package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var movie models.Movie

	if result := h.DB.Where("imdb_id = ?", id).First(&movie); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(http.StatusNotFound)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movie)
	}
}

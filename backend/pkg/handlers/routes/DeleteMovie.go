package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var movie models.Movie
	
	if result := h.DB.Where("imdb_id = ?", id).Find(&movie); result.Error == nil {
		h.DB.Delete(&movie)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(http.StatusOK)
	}
}

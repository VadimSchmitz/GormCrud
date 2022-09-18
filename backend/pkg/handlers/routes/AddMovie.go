package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
	"github.com/VadimSchmitz/GormCrud/pkg/util"
)

func (h handler) AddMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	util.CheckErr(err)
	
	var movie models.Movie
	json.Unmarshal(body, &movie)

	if result := h.DB.Create(&movie); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(movie)
}

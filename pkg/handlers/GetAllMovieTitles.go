package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
)

func (h handler) GetAllMovieTitles(w http.ResponseWriter, r *http.Request) {
	var movie *models.Movie
	var title []models.Title

	if result := h.DB.Model(&movie).Find(&title); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(title)
}
	
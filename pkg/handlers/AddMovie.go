package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
)

func (h handler) AddMovie(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	
	var movie models.Movie
	json.Unmarshal(body, &movie)

	// Append to the movie table
	if result := h.DB.Create(&movie); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// {"id":"tt0368226","name":"The Room","year":"2003","score":"3.7"}
	json.NewEncoder(w).Encode(movie)
}

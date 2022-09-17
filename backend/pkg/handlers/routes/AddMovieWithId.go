package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	p "github.com/VadimSchmitz/GormCrud/pkg/handlers/externalApi"
	m "github.com/VadimSchmitz/GormCrud/pkg/models"
)

func (h handler) AddMovieWithId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var imdb_id m.IMDb_id
	json.Unmarshal(body, &imdb_id)
		
	idUrl := `http://www.omdbapi.com/?i=` + imdb_id.IMDb_id + `&apikey=183b9a26`

	var movie m.Movie
	p.RetrieveMovieWithId(idUrl, &movie)

	if result := h.DB.Create(&movie); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

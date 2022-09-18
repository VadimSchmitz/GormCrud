package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"

	p "github.com/VadimSchmitz/GormCrud/pkg/handlers/externalApi"
	m "github.com/VadimSchmitz/GormCrud/pkg/models"
	"github.com/VadimSchmitz/GormCrud/pkg/util"
)

func (h handler) AddMovieWithId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	util.CheckErr(err)

	var imdb_id m.IMDb_id
	json.Unmarshal(body, &imdb_id)
	
	e, _:= regexp.MatchString(`(tt[0-9]{7})`, imdb_id.IMDb_id)

	if !e {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid IMDb ID")
		return
	}
	
	idUrl := `http://www.omdbapi.com/?i=` + imdb_id.IMDb_id + `&apikey=f7e4139c`

	var movie m.Movie
	p.RetrieveMovieWithId(idUrl, &movie)

	if result := h.DB.Create(&movie); result.Error != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Movie already exists")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

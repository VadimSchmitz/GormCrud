package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
	"github.com/VadimSchmitz/GormCrud/pkg/util"
)

func RetrieveMovieWithId(url string, movie *models.Movie) {
	response, err := http.Get(url)
	util.CheckErr(err)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	util.CheckErr(err)

	var omdb_body models.Omdb_body
	err = json.Unmarshal(body, &omdb_body)

	movie.IMDb_id = omdb_body.IMDb_id
	movie.Title = omdb_body.Title
	movie.Rating = omdb_body.Rating
	movie.Year = omdb_body.Year
	movie.Plot_summary = omdb_body.Plot_summary

	util.CheckErr(err)
}

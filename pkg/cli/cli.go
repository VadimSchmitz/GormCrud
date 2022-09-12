package cli

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	h "github.com/VadimSchmitz/GormCrud/pkg/handlers/externalApi"
	"github.com/VadimSchmitz/GormCrud/pkg/models"
)

func CliHandler(done chan bool) {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		return
	}

	if arguments[0] == "summaries" {
		
		var IMDb_id []*models.IMDb_id		
		idUrl := "http://localhost:8090/ids"
		h.ProcessRequest(idUrl, &IMDb_id)

		var Plot_summary *models.Plot		

		for _, id := range IMDb_id {
			omdbUrl := "http://www.omdbapi.com/?i=" + id.IMDb_id + "&apikey=183b9a26"
			h.ProcessRequest(omdbUrl, &Plot_summary)

			var jsonStr = []byte(`{"plot": "` + Plot_summary.Plot_summary + `"}`)
			req, _ := http.NewRequest("PATCH", "http://localhost:8090/movies/" + id.IMDb_id, bytes.NewBuffer(jsonStr))
			
			client := &http.Client{}
			client.Do(req)
		}

		fmt.Println("Summaries added")
		done <- true
	}
}

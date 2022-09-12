package cli

import (
	"fmt"
	"os"

	handlers "github.com/VadimSchmitz/GormCrud/pkg/handlers/externalApi"
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
		handlers.ProcessRequest(idUrl, &IMDb_id)

		var Plot_summary *models.Plot		
		omdbUrl := "http://www.omdbapi.com/?i=tt3896198&apikey=183b9a26"
		handlers.ProcessRequest(omdbUrl, &Plot_summary)


		//loop over all the slice of titles
		for _, id := range IMDb_id {
			fmt.Println(id.IMDb_id)
		}

		done <- true
	}
}

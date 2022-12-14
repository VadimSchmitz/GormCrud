package cli

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	h "github.com/VadimSchmitz/GormCrud/pkg/handlers/externalApi"
	"github.com/VadimSchmitz/GormCrud/pkg/models"
	e "github.com/VadimSchmitz/GormCrud/pkg/models/external"
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

		numJobs := len(IMDb_id)

		jobs := make(chan *models.IMDb_id, numJobs)
		results := make(chan string, numJobs)

		for w := 1; w <= 5; w++ {
			go FetchAndPatch(w, jobs, results)
		}

		for _, imdb_id := range IMDb_id {
			jobs <- imdb_id
		}

		close(jobs)

		for a := 1; a <= numJobs; a++ {
			// fmt.Printf("Wrote summary for %s to database\n", <-results)
			<-results
		}

		fmt.Println("Summaries added")
		done <- true
	}
}

func FetchAndPatch(wId int, jobs <-chan *models.IMDb_id, results chan<- string) {
	for imdb_id := range jobs {
		// fmt.Printf("Worker %d getting summary for %s \n", wId, imdb_id.IMDb_id)

		omdbUrl := "http://www.omdbapi.com/?i=" + imdb_id.IMDb_id + "&apikey=f7e4139c"
		var plot *e.Plot

		// fmt.Printf("Worker %d received summary for %s \n", wId, imdb_id.IMDb_id)

		h.ProcessRequest(omdbUrl, &plot)

		var jsonStr = []byte(`{"summary": "` + plot.Plot_summary + `"}`)
		req, _ := http.NewRequest("PATCH", "http://localhost:8090/movies/"+imdb_id.IMDb_id, bytes.NewBuffer(jsonStr))

		client := &http.Client{}
		client.Do(req)
		
		results <- imdb_id.IMDb_id
	}
}

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	// "github.com/VadimSchmitz/GormCrud/pkg/handlers"
)

func CliHandler(done chan bool) {
	fmt.Println("Welcome to the CLI")
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		return
	}

	if arguments[0] == "summaries" {
		//slice of strings to hold the movie titles returned from GetAllMovieTitles()
		// allTitles := GetAllMovieTitles(h)
		
		url := "http://www.omdbapi.com/?i=tt3896198&apikey=183b9a26"

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		spaceClient := http.Client{
			Timeout: time.Second * 2, // Timeout after 2 seconds
		}

		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		type title struct{
			Title string `json:"title"`
		}
		
		title1 := title{}
		jsonErr := json.Unmarshal(body, &title1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		//loop over all the entries of imdb_id
		
		fmt.Println(title1)
		fmt.Println("Summaries added")

		done <- true
	}
}

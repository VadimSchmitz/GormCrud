package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/VadimSchmitz/GormCrud/pkg/models"
)

func CliHandler(done chan bool) {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		return
	}

	if arguments[0] == "summaries" {
		
		response, apiError := http.Get("http://localhost:8090/titles")

		if apiError != nil {
			log.Fatal(apiError)
		}

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}

		var titles []models.Title		

		err = json.Unmarshal(body, &titles)
		if err != nil {
			panic(err.Error())
		}

		//loop over all the slice of titles
		for _, title := range titles {
			fmt.Println(title.Title)
			
		}
			
			done <- true
	}
}

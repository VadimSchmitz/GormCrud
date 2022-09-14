package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/VadimSchmitz/GormCrud/pkg/util"
)

func ProcessRequest(url string, arr interface{}) {
	response, err := http.Get(url)
	util.CheckErr(err)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	util.CheckErr(err)

	err = json.Unmarshal(body, &arr)
	util.CheckErr(err)
}

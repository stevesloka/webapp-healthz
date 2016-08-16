package healthz

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type APIChecker struct {
	APIUrl     string
	MinVersion string
}

type Appversion struct {
	Appver string `json:"version"`
}

func NewAPIChecker(APIUrl, MinVersion string) (*APIChecker, error) {
	return &APIChecker{APIUrl, MinVersion}, nil
}

func (api *APIChecker) CheckVersion() error {

	req, err := http.Get(api.APIUrl)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	data := Appversion{}
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)

	json.Unmarshal([]byte(string(jsonDataFromHttp[:])), &data)

	if err != nil {
		return err
	}

	log.Println("ver1: ", data.Appver)
	log.Println("minv: ", api.MinVersion)

	if data.Appver != api.MinVersion {
		return errors.New("versionMismatch")
	}

	return nil
}

package healthz

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APIChecker struct {
	APIUrl     string
	MinVersion string
}

type Version struct {
	AppVersion string `json:appVersion`
}

func NewAPIChecker(APIUrl, MinVersion string) (*APIChecker, error) {
	return &APIChecker{APIUrl, MinVersion}, nil
}

func (api *APIChecker) CheckVersion() error {

	req, err := http.Get(api.APIUrl)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(req.Body)

	data := Version{}
	err = decoder.Decode(&data)

	if err != nil {
		return err
	}

	if data.AppVersion != api.MinVersion {
		return errors.New("boo")
	}

	return nil
}

package healthz

import (
	"encoding/json"
	"log"
	"net/http"
)

type Config struct {
	Hostname string
	API      APIConfig
}

type APIConfig struct {
	APIUrl     string
	MinVersion string
}

type handler struct {
	api      *APIChecker
	hostname string
	metadata map[string]string
}

func Handler(hc *Config) (http.Handler, error) {
	api, err := NewAPIChecker(hc.API.APIUrl, hc.API.MinVersion)
	if err != nil {
		return nil, err
	}

	metadata := make(map[string]string)
	metadata["api_url"] = hc.API.APIUrl
	metadata["min_version"] = hc.API.MinVersion

	h := &handler{api, hc.Hostname, metadata}
	return h, nil
}

type Response struct {
	Hostname string            `json:"hostname"`
	Metadata map[string]string `json:"metadata"`
	Errors   []Error           `json:"errors"`
}

type Error struct {
	Description string            `json:"description"`
	Error       string            `json:"error"`
	Metadata    map[string]string `json:"metadata"`
	Type        string            `json:"type"`
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Hostname: h.hostname,
		Metadata: h.metadata,
	}

	log.Println("|| request made |||")

	statusCode := http.StatusOK

	errors := make([]Error, 0)

	err := h.api.CheckVersion()
	if err != nil {
		errors = append(errors, Error{
			Type:        "APIVersion",
			Description: "API Version is out of date.",
			Error:       err.Error(),
		})
	}

	response.Errors = errors
	if len(response.Errors) > 0 {
		statusCode = http.StatusInternalServerError
		for _, e := range response.Errors {
			log.Println(e.Error)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	log.Println("response: ", response)

	data, err := json.MarshalIndent(&response, "", "  ")
	if err != nil {
		log.Println(err)
	}

	w.Write(data)
}

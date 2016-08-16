package main

import (
	"log"
	"net/http"
	"os"

	"steve.io/abstractions/webapp-healthz/healthz"
)

func main() {
	log.Println("Starting webapp-healthz...")

	httpAddr := os.Getenv("HTTP_ADDR")
	apiURL := os.Getenv("APIURL")
	minVersion := os.Getenv("MINVERSION")

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing api ...")

	hc := &healthz.Config{
		Hostname: hostname,
		API: healthz.APIConfig{
			APIUrl:     apiURL,
			MinVersion: minVersion,
		},
	}

	healthzHandler, err := healthz.Handler(hc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("HTTP service listening on %s", httpAddr)
	http.Handle("/healthz", healthzHandler)
	http.ListenAndServe(httpAddr, nil)
}

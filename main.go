package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stevesloka/webapp-healthz/healthz"
)

func main() {
	log.Println("Starting webapp-healthz...")

	httpAddr := os.Getenv("HTTP_ADDR")
	apiURL := os.Getenv("APIURL")
	minVersion := os.Getenv("MINVERSION")

	log.Println("MINVERSION: ", minVersion)
	log.Println("APIURL: ", apiURL)

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

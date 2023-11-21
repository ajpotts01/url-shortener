package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Failed to load dotenv: %v\n", err)
		log.Printf("This is not a fatal error.\n")
	}

	apiConfig, err := getApiConfig("stub")

	if err != nil {
		log.Fatalf("Failed to initialise API config: %v", err)
	}

	log.Printf("API config: %v", apiConfig)

	router := getAppRouter()
	router.Mount("/v1", getApiRouterV1(apiConfig))

	server := &http.Server{
		Addr:              ":6666",
		Handler:           router,
		ReadHeaderTimeout: 15 * time.Second,
	}

	log.Printf("Now listening on port 6666")
	log.Fatal(server.ListenAndServe())
}

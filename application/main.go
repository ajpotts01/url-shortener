package main

import (
	"log"
	"net/http"
	"os"
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

	// Normally for an API, would do a V1, V2 mount etc.
	// But this is a URL shortener, and expected experience is not to access www.url.com/v1/{key}...
	router.Mount("/", getApiRouter(apiConfig))

	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 15 * time.Second,
	}

	log.Printf("Now listening on port %v", port)
	log.Fatal(server.ListenAndServe())
}

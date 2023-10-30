package main

import (
	"log"
	"net/http"
)

func main() {
	apiConfig, err := getApiConfig("stub")

	if err != nil {
		log.Fatalf("Failed to initialise API config: %v", err)
	}

	log.Printf("API config: %v", apiConfig)

	router := getAppRouter()
	router.Mount("/v1", getApiRouterV1(apiConfig))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Now listening on port 8080")
	log.Fatal(server.ListenAndServe())
}

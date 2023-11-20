package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiConfig, err := getApiConfig("stub")

	if err != nil {
		log.Fatalf("Failed to initialise API config: %v", err)
	}

	log.Printf("API config: %v", apiConfig)

	router := getAppRouter()
	router.Mount("/v1", getApiRouterV1(apiConfig))

	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	server := &http.Server{
		Addr:              ":6666",
		Handler:           router,
		ReadHeaderTimeout: 15 * time.Second,
	}

	log.Printf("Now listening on port 6666")
	log.Fatal(server.ListenAndServe())
}

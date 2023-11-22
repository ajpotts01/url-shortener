package main

import (
	"context"
	"log"
	"os"

	"github.com/ajpotts01/url-shortener/application/api"
	"github.com/ajpotts01/url-shortener/application/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func getApiConfig(dbConnStr string) (*api.ApiConfig, error) {
	projectId := os.Getenv("PROJECT_ID")
	databaseId := os.Getenv("DATABASE_ID")

	dbConn := database.CreateConnection(context.Background(), projectId, databaseId)

	return &api.ApiConfig{
		DbConn: &database.Database{
			DbClient: dbConn,
		},
	}, nil
}

func getApiRouterV1(config *api.ApiConfig) *chi.Mux {
	const healthEndpoint = "/healthz"
	const createEndpoint = "/create"

	apiRouter := chi.NewRouter()
	apiRouter.Get(healthEndpoint, api.Health)
	apiRouter.Post(createEndpoint, config.CreateShortenedUrl)

	log.Printf("API router init: %v", apiRouter)

	return apiRouter
}

func getAppRouter() *chi.Mux {
	appRouter := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET, POST, OPTIONS, PUT, DELETE"},
		AllowedHeaders: []string{"*"},
	}
	appRouter.Use(cors.Handler(corsOptions))

	return appRouter
}

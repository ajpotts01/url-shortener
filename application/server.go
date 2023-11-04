package main

import (
	"github.com/ajpotts01/url-shortener/application/api"
	"github.com/ajpotts01/url-shortener/application/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func getApiConfig(dbConnStr string) (*api.ApiConfig, error) {
	dbConn := database.Database{
		ConnStr: dbConnStr,
	}

	return &api.ApiConfig{
		DbConn: dbConn,
	}, nil
}

func getApiRouterV1(config *api.ApiConfig) *chi.Mux {
	const healthEndpoint = "/healthz"

	apiRouter := chi.NewRouter()
	apiRouter.Get(healthEndpoint, api.Health)

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

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ajpotts01/url-shortener/application/internal/database"
	"github.com/ajpotts01/url-shortener/application/internal/urls"
	"github.com/go-chi/chi/v5"
)

type urlRequest struct {
	Url string `json:"url"`
}

// POST /v1/create
func (config *ApiConfig) CreateShortenedUrl(w http.ResponseWriter, r *http.Request) {
	log.Printf("Creating shortened URL")
	decoder := json.NewDecoder(r.Body)
	requestParams := urlRequest{}
	err := decoder.Decode(&requestParams)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	urlShortKey, err := urls.GetShortenedUrl(requestParams.Url)

	if err != nil {
		log.Printf("Error creating new shortened URL: %v", err)
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	domain := os.Getenv("DOMAIN_NAME")
	port := os.Getenv("PORT")
	protocol := os.Getenv("PROTOCOL")

	urlShort := constructShortUrl(protocol, domain, port, urlShortKey)

	if err != nil {
		log.Printf("Error getting hostname: %v", err)
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	dbParams := database.CreateUrlParams{
		Key:      urlShortKey,
		UrlLong:  requestParams.Url,
		UrlShort: urlShort,
	}

	err = config.DbConn.CreateUrl(context.Background(), dbParams)

	if err != nil {
		log.Printf("Error saving to database: %v", err)
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	validResponse(w, http.StatusCreated, Url{
		Key:      urlShortKey,
		UrlLong:  requestParams.Url,
		UrlShort: urlShort,
	})
}

// GET /
func (config *ApiConfig) FetchShortenedUrl(w http.ResponseWriter, r *http.Request) {
	urlKey := chi.URLParam(r, "key")

	w.Header().Set("Content-Type", "application/json")

	params := database.FetchUrlParams{
		Key: urlKey,
	}

	if urlKey == "" {
		errorResponse(w, http.StatusBadRequest, "Please specify a URL key")
		return
	}

	url, err := config.DbConn.FetchUrl(context.Background(), params)

	if err != nil {
		log.Printf("Error fetching URL: %v", err)
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	validResponse(w, http.StatusOK, Url{
		Key:      urlKey,
		UrlLong:  url.UrlLong,
		UrlShort: url.UrlShort,
	})
}

func constructShortUrl(protocol string, domain string, port string, shortKey string) string {
	urlShort := fmt.Sprintf("%s://%s", protocol, domain)

	if port != "" {
		urlShort = fmt.Sprintf("%s:%s", urlShort, port)
	}

	urlShort = fmt.Sprintf("%s/%s", urlShort, shortKey)

	return urlShort
}

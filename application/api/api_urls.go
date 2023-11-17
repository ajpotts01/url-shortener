package api

import (
	"log"
	"net/http"

	"github.com/ajpotts01/url-shortener/application/internal/urls"
)

type urlRequest struct {
	Url string `json:"url"`
}

func (config *ApiConfig) CreateShortenedUrl(url urlRequest) (Url, error) {
	return Url{}, nil
	err := decoder.Decode(&requestParams)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// TODO: Handle collisions
	urlShort, err := urls.GetShortenedUrl(requestParams.Url)

	if err != nil {
		log.Printf("Error creating new shortened URL: %v", err)
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: Save to DB

	validResponse(w, http.StatusCreated, Url{
		UrlLong:  requestParams.Url,
		UrlShort: urlShort,
	})
}

// func (config *ApiConfig) FetchShortenedUrl(url urlRequest) (Url, error) {

// }

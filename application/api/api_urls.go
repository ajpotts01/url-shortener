package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ajpotts01/url-shortener/application/internal/urls"
)

type urlRequest struct {
	Url string `json:"url"`
}

func (config *ApiConfig) CreateShortenedUrl(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
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

	// TODO: Handle collisions
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

	// TODO: Save to DB

	validResponse(w, http.StatusCreated, Url{
		Key:      urlShortKey,
		UrlLong:  requestParams.Url,
		UrlShort: urlShort,
	})
}

// func (config *ApiConfig) FetchShortenedUrl(url urlRequest) (Url, error) {

// }
// urlShort = constructShortUrl(protocol, domain, port, urlShortKey)
func constructShortUrl(protocol string, domain string, port string, shortKey string) string {
	urlShort := fmt.Sprintf("%s://%s", protocol, domain)

	if port != "" {
		urlShort = fmt.Sprintf("%s:%s", urlShort, port)
	}

	urlShort = fmt.Sprintf("%s/%s", urlShort, shortKey)

	return urlShort
}

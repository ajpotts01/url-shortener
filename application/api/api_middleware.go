package api

import (
	"log"
	"net/http"

	"github.com/ajpotts01/url-shortener/application/internal/auth"
)

type authorisedMethod func(http.ResponseWriter, *http.Request)

func (config *ApiConfig) AuthMiddleware(method authorisedMethod) http.HandlerFunc {
	// No body expected: just API key header
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("Now in auth middleware")

		key, err := auth.GetAuthFromHeader(r, "ApiKey")
		if err != nil {
			errorResponse(w, http.StatusUnauthorized, "Bad authorization header")
			return
		}

		log.Printf("Received API key: %v", key)
		//usr, err := config.DbConn.GetUserByApiKey(context.TODO(), key)

		if err != nil {
			errorResponse(w, http.StatusUnauthorized, "Bad API key")
			return
		}

		method(w, r)
	})
}

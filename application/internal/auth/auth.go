package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAuthFromHeader(r *http.Request, tokenType string) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("must supply authorization header")
	}

	suppliedItem := strings.Replace(authHeader, tokenType, "", 1)
	suppliedItem = strings.Trim(suppliedItem, " ")
	return suppliedItem, nil
}

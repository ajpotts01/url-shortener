package api

import "github.com/ajpotts01/url-shortener/application/internal/database"

type ApiConfig struct {
	DbConn *database.Database
}

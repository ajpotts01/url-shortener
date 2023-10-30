package api

import "github.com/ajpotts01/url-shortener/internal/database"

type ApiConfig struct {
	DbConn database.Database
}

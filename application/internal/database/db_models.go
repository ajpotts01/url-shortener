package database

import "time"

type Url struct {
	Key       string
	UrlLong   string
	UrlShort  string
	CreatedAt time.Time
}

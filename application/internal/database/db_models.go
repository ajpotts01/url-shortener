package database

import "time"

type Url struct {
	UrlLong   string
	UrlShort  string
	CreatedAt time.Time
}

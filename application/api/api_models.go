package api

// import "github.com/ajpotts01/url-shortener/application/internal/database"

type Url struct {
	Key      string `json:"key"`
	UrlLong  string `json:"long_url"`
	UrlShort string `json:"short_url"`
}

// func DatabaseUrlToApiUrl(url database.Url) Url {
// 	return Url{
// 		UrlLong:  url.UrlLong,
// 		UrlShort: url.UrlShort,
// 	}
// }

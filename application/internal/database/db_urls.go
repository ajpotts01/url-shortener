package database

import (
	"context"
	"log"
	"time"
)

type CreateUrlParams struct {
	Key      string
	UrlLong  string
	UrlShort string
}

type FetchUrlParams struct {
	Key string
}

func (database *Database) CreateUrl(ctx context.Context, urlParams CreateUrlParams) error {
	newUrl := Url{
		UrlLong:   urlParams.UrlLong,
		UrlShort:  urlParams.UrlShort,
		CreatedAt: time.Now(),
	}

	docMap := urlToFirestoreDoc(newUrl)
	_, err := database.DbClient.Collection("urls").Doc(urlParams.Key).Set(ctx, docMap)

	if err != nil {
		log.Printf("Error creating new URL: %v", err)
		return err
	}

	return nil
}

func urlToFirestoreDoc(url Url) map[string]interface{} {
	return map[string]interface{}{
		"url_short":  url.UrlShort,
		"url_long":   url.UrlLong,
		"created_at": url.CreatedAt,
	}
}

func firestoreDocToUrl(docMap map[string]interface{}) Url {
	return Url{
		UrlLong:   docMap["url_long"].(string),
		UrlShort:  docMap["url_short"].(string),
		CreatedAt: docMap["created_at"].(time.Time),
	}
}

func (database *Database) FetchUrl(ctx context.Context, urlParams FetchUrlParams) (Url, error) {
	docSnap, err := database.DbClient.Collection("urls").Doc(urlParams.Key).Get(ctx)

	if err != nil {
		log.Printf("Error fetching URL: %v", err)
		return Url{}, err
	}

	docMap := docSnap.Data()
	url := firestoreDocToUrl(docMap)

	return url, nil
}

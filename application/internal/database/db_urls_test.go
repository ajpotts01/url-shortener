package database

import (
	"reflect"
	"testing"
	"time"
)

func TestUrlToDoc(t *testing.T) {
	input := Url{
		UrlLong:   "https://ajpcloudblog.com",
		UrlShort:  "https://ajpurlshortener.com:6666/HEqHPdx",
		CreatedAt: time.Date(2023, 11, 22, 0, 0, 0, 0, time.UTC),
	}

	want := map[string]interface{}{
		"url_short":  "https://ajpurlshortener.com:6666/HEqHPdx",
		"url_long":   "https://ajpcloudblog.com",
		"created_at": "2023-11-22 00:00:00.000000 +0000 UTC",
	}

	got := urlToFirestoreDoc(input)

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

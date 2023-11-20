package api

import (
	"testing"
)

func TestConstructShortUrlHttp(t *testing.T) {
	domain := "shorturl.com"
	port := "6666"
	protocol := "http"
	shortKey := "HEqHPdx"

	want := "http://shorturl.com:6666/HEqHPdx"
	got := constructShortUrl(protocol, domain, port, shortKey)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConstructShortUrlHttps(t *testing.T) {
	domain := "shorturl.com"
	port := "6666"
	protocol := "https"
	shortKey := "HEqHPdx"

	want := "https://shorturl.com:6666/HEqHPdx"
	got := constructShortUrl(protocol, domain, port, shortKey)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConstructShortUrlHttpNoPort(t *testing.T) {
	domain := "shorturl.com"
	port := ""
	protocol := "http"
	shortKey := "HEqHPdx"

	want := "http://shorturl.com/HEqHPdx"
	got := constructShortUrl(protocol, domain, port, shortKey)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerTokenFromHeader(t *testing.T) {
	testToken := "Bearer da39a3ee5e6b4b0d3255bfef95601890afd80709"
	want := "da39a3ee5e6b4b0d3255bfef95601890afd80709"

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", testToken)

	got, _ := GetAuthFromHeader(request, "Bearer")

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetApiKeyTokenFromHeader(t *testing.T) {
	testToken := "ApiKey 38468bdaaca66caa43308042322ef0cf407cf912"
	want := "38468bdaaca66caa43308042322ef0cf407cf912"

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", testToken)

	got, _ := GetAuthFromHeader(request, "ApiKey")

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

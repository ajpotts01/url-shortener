package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ajpotts01/url-shortener/application/api"
)

func TestHealth(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	api.Health(response, request)

	t.Run("Gets valid OK status from health check", func(t *testing.T) {
		got := response.Result().StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

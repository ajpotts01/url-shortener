package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidResponseBase(t *testing.T) {
	response := httptest.NewRecorder()

	url := Url{
		Key:      "HEqHPdx",
		UrlLong:  "https://ajpcloudblog.com",
		UrlShort: "https://www.shorturl.com/HEqHPdx",
	}

	validResponse(response, 200, url)

	t.Run("Object is as expected", func(t *testing.T) {
		gotStatus := response.Result().StatusCode
		gotJson := response.Result().Body

		wantStatus := http.StatusOK

		// Decode the response into a new object.
		// This should match the original one passed to validResponse.
		decoder := json.NewDecoder(gotJson)
		gotObj := Url{}
		err := decoder.Decode(&gotObj)

		if err != nil {
			t.Errorf("Failed to get valid response: %v", err)
		}

		if gotStatus != wantStatus {
			t.Errorf("got %d, want %d", gotStatus, wantStatus)
		}

		if gotObj != url {
			t.Errorf("got %q, want %q", gotObj, url)
		}
	})
}

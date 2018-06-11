package http_test

import (
	. "github.com/bmuschko/link-verifier/http"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestGetValidUrl(t *testing.T) {
	url := "http://www.google.com/"
	result := Get(url)

	Equal(t, url, result.Url)
	True(t, result.Success)
	Equal(t, 200, result.StatusCode)
}

func TestGetNonExistentUrl(t *testing.T) {
	url := "http://www.unknown1x.com/"
	Panics(t, func() {
		Get(url)
	})
}

func TestGetUrlForBadRequest(t *testing.T) {
	url := "https://www.googleapis.com/urlshortener/v1/url"
	result := Get(url)

	Equal(t, url, result.Url)
	False(t, result.Success)
	Equal(t, 400, result.StatusCode)
}

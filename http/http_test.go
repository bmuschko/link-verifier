package http_test

import (
	. "github.com/bmuschko/link-verifier/http"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestSetTimeout(t *testing.T) {
	SetTimeout(20)
}

func TestGetValidUrl(t *testing.T) {
	url := "http://www.google.com/"
	result := Get(url)

	Equal(t, url, result.Url)
	True(t, result.Success)
	Nil(t, result.Error)
	Equal(t, 200, result.StatusCode)
}

func TestGetUrlForBadRequest(t *testing.T) {
	url := "https://www.googleapis.com/urlshortener/v1/url"
	result := Get(url)

	Equal(t, url, result.Url)
	False(t, result.Success)
	Nil(t, result.Error)
	Equal(t, 400, result.StatusCode)
}

func TestGetNonExistentUrl(t *testing.T) {
	url := "http://www.unknown1x.com/"
	result := Get(url)

	Equal(t, url, result.Url)
	False(t, result.Success)
	NotNil(t, result.Error)
}

func TestGetInvalidUrl(t *testing.T) {
	url := "123://www.invalid.com/"
	result := Get(url)

	Equal(t, url, result.Url)
	False(t, result.Success)
	NotNil(t, result.Error)
}

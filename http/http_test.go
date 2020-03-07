package http_test

import (
	. "github.com/bmuschko/link-verifier/http"
	. "github.com/stretchr/testify/assert"
	"testing"
)

const (
	validUrl      = "http://www.google.com/"
	incompleteUrl = "https://www.googleapis.com/urlshortener/v1/url"
	unknownUrl    = "http://www.unknown1x.com/"
	invalidUrl    = "123://www.invalid.com/"
)

func TestSetTimeout(t *testing.T) {
	SetTimeout(20)
}

func TestHeadValidUrl(t *testing.T) {
	result := Head(validUrl)

	Equal(t, validUrl, result.Url)
	True(t, result.Success)
	Nil(t, result.Error)
	Equal(t, 200, result.StatusCode)
}

func TestHeadUrlForBadRequest(t *testing.T) {
	t.Skip("Needs to emulate an HTTP server to reproducible results")
	result := Head(incompleteUrl)

	Equal(t, incompleteUrl, result.Url)
	False(t, result.Success)
	Nil(t, result.Error)
	Equal(t, 400, result.StatusCode)
}

func TestHeadNonExistentUrl(t *testing.T) {
	result := Head(unknownUrl)

	Equal(t, unknownUrl, result.Url)
	False(t, result.Success)
	NotNil(t, result.Error)
}

func TestHeadInvalidUrl(t *testing.T) {
	result := Head(invalidUrl)

	Equal(t, invalidUrl, result.Url)
	False(t, result.Success)
	NotNil(t, result.Error)
}

func TestGetValidUrl(t *testing.T) {
	result := Get(validUrl)

	Equal(t, validUrl, result.Url)
	True(t, result.Success)
	Nil(t, result.Error)
	Equal(t, 200, result.StatusCode)
}

func TestGetUrlForBadRequest(t *testing.T) {
	t.Skip("Needs to emulate an HTTP server to reproducible results")
	result := Get(incompleteUrl)

	Equal(t, incompleteUrl, result.Url)
	False(t, result.Success)
	Nil(t, result.Error)
	Equal(t, 400, result.StatusCode)
}

func TestGetNonExistentUrl(t *testing.T) {
	result := Get(unknownUrl)

	Equal(t, unknownUrl, result.Url)
	False(t, result.Success)
	NotNil(t, result.Error)
}

func TestGetInvalidUrl(t *testing.T) {
	result := Get(invalidUrl)

	Equal(t, invalidUrl, result.Url)
	False(t, result.Success)
	NotNil(t, result.Error)
}

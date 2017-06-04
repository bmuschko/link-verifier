package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetValidUrl(t *testing.T) {
	url := "http://www.google.com/"
	result := Get(url)

	assert.Equal(t, url, result.Url)
	assert.True(t, result.Success)
	assert.Equal(t, 200, result.StatusCode)
}

func TestGetNonExistentUrl(t *testing.T) {
	url := "http://www.unknown1x.com/"
	assert.Panics(t, func() {
		Get(url)
	})
}

func TestGetUrlForBadRequest(t *testing.T) {
	url := "https://www.googleapis.com/urlshortener/v1/url"
	result := Get(url)

	assert.Equal(t, url, result.Url)
	assert.False(t, result.Success)
	assert.Equal(t, 400, result.StatusCode)
}

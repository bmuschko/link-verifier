package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/bmuschko/link-verifier/http"
	. "github.com/stretchr/testify/assert"
)

const (
	invalidUrl = "123://www.invalid.com/"
)

func TestSetTimeout(t *testing.T) {
	h := NewHTTP()
	h.SetTimeout(20)

	Equal(t, h.GetTimeout(), time.Duration(int(time.Second)*20))
}

func TestHeadValidUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	h := NewHTTP()
	result := h.Head(server.URL)

	Equal(t, server.URL, result.Url)
	True(t, result.Success)
	Nil(t, result.Error)
	Equal(t, http.StatusOK, result.StatusCode)
}

func TestGetValidUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	h := NewHTTP()
	result := h.Get(server.URL)

	Equal(t, server.URL, result.Url)
	True(t, result.Success)
	Nil(t, result.Error)
	Equal(t, http.StatusOK, result.StatusCode)
}

func TestHeadUrlForBadRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	h := NewHTTP()
	result := h.Head(server.URL)

	Equal(t, server.URL, result.Url)
	False(t, result.Success)
	Nil(t, result.Error)
	Equal(t, http.StatusBadRequest, result.StatusCode)
}

func TestGetUrlForBadRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	h := NewHTTP()
	result := h.Get(server.URL)

	Equal(t, server.URL, result.Url)
	False(t, result.Success)
	Nil(t, result.Error)
	Equal(t, http.StatusBadRequest, result.StatusCode)
}

func TestHeadNotFoundUrl(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()

	h := NewHTTP()
	result := h.Head(server.URL)

	Equal(t, server.URL, result.Url)
	False(t, result.Success)
	Nil(t, result.Error)
	Equal(t, http.StatusNotFound, result.StatusCode)
}

func TestGetNotFoundUrl(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()

	h := NewHTTP()
	result := h.Get(server.URL)

	Equal(t, server.URL, result.Url)
	False(t, result.Success)
	Nil(t, result.Error)
	Equal(t, http.StatusNotFound, result.StatusCode)
}

func TestHeadInvalidUrl(t *testing.T) {
	h := NewHTTP()
	result := h.Head(invalidUrl)

	Equal(t, invalidUrl, result.Url)
	False(t, result.Success)
	NotNil(t, result.Error)
}

func TestGetInvalidUrl(t *testing.T) {
	h := NewHTTP()
	result := h.Get(invalidUrl)

	Equal(t, invalidUrl, result.Url)
	False(t, result.Success)
	NotNil(t, result.Error)
}

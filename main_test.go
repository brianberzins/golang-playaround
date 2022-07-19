package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequests(t *testing.T) {
	tests := []struct {
		endpoint     string
		handler      func(responseWriter http.ResponseWriter, _ *http.Request)
		expectedCode int
		expectedBody string
	}{
		{endpoint: "/hello", expectedCode: http.StatusOK, expectedBody: "hello"},
		{endpoint: "/goodbye", expectedCode: http.StatusOK, expectedBody: "goodbye"},
		{endpoint: "/hello/murasaki", expectedCode: http.StatusOK, expectedBody: "hello murasaki"},
	}

	ro := createRouter()
	for _, tc := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, tc.endpoint, nil)
		ro.ServeHTTP(w, r)
		assert.Equal(t, tc.expectedCode, w.Code)
		assert.Equal(t, tc.expectedBody, w.Body.String())
	}
}

func TestPostRequest(t *testing.T) {
	ro := createRouter()
	w := httptest.NewRecorder()
	b, _ := json.Marshal(incrementer{Start: 3, Increment: 2})
	r := httptest.NewRequest(http.MethodPost, "/incrementer", bytes.NewReader(b))
	ro.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "5", w.Body.String())
}

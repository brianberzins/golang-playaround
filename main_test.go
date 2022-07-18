package main

import (
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
		{endpoint: "/hello", handler: greeting, expectedCode: http.StatusOK, expectedBody: "hello"},
		{endpoint: "/goodbye", handler: goodbye, expectedCode: http.StatusOK, expectedBody: "goodbye"},
	}

	for _, tc := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, tc.endpoint, nil)
		tc.handler(w, r)
		assert.Equal(t, tc.expectedCode, w.Code)
		assert.Equal(t, tc.expectedBody, w.Body.String())
	}
}

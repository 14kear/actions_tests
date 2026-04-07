package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloEndpointReturnsStatusOK(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	recorder := httptest.NewRecorder()

	newMux().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
}

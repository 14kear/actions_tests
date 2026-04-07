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

func TestHelloEndpointReturnsMethodNotAllowedForPost(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/api/hello", nil)
	recorder := httptest.NewRecorder()

	newMux().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}

func TestServerAddressUsesDefaultPort(t *testing.T) {
	if got := serverAddress(""); got != ":8080" {
		t.Fatalf("expected default address :8080, got %s", got)
	}
}

func TestServerAddressUsesProvidedPort(t *testing.T) {
	if got := serverAddress("9090"); got != ":9090" {
		t.Fatalf("expected address :9090, got %s", got)
	}
}

func TestRunUsesConfiguredPortAndHandler(t *testing.T) {
	var receivedAddress string
	var receivedHandler http.Handler

	err := run(
		func(key string) string {
			if key != "PORT" {
				t.Fatalf("expected PORT env lookup, got %s", key)
			}

			return "9090"
		},
		func(address string, handler http.Handler) error {
			receivedAddress = address
			receivedHandler = handler
			return nil
		},
	)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if receivedAddress != ":9090" {
		t.Fatalf("expected address :9090, got %s", receivedAddress)
	}

	request := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	recorder := httptest.NewRecorder()

	receivedHandler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
}

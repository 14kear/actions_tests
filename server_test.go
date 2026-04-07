package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloEndpoint(t *testing.T) {
	handler := newMux()
	testCases := []struct {
		name       string
		method     string
		wantStatus int
	}{
		{name: "get returns ok", method: http.MethodGet, wantStatus: http.StatusOK},
		{name: "post returns method not allowed", method: http.MethodPost, wantStatus: http.StatusMethodNotAllowed},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := httptest.NewRequest(testCase.method, "/api/hello", nil)
			recorder := httptest.NewRecorder()

			handler.ServeHTTP(recorder, request)

			if recorder.Code != testCase.wantStatus {
				t.Fatalf("expected status %d, got %d", testCase.wantStatus, recorder.Code)
			}
		})
	}
}

func TestRunUsesDefaultPortAndHandler(t *testing.T) {
	var receivedAddress string
	var receivedHandler http.Handler

	err := run(
		func(key string) string {
			if key != "PORT" {
				t.Fatalf("expected PORT env lookup, got %s", key)
			}

			return ""
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

	if receivedAddress != ":8080" {
		t.Fatalf("expected address :8080, got %s", receivedAddress)
	}

	request := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	recorder := httptest.NewRecorder()

	receivedHandler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a new request
	req, err := http.NewRequest("GET", "", nil)

	// In case of error we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	// recorder acts as the target of the http request
	recorder := httptest.NewRecorder()

	// handler is the function we want to test from main
	hf := http.HandlerFunc(handler)

	hf.ServeHTTP(recorder, req)

	// Check if the status code is what we want
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

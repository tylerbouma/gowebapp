package main

import (
	"encoding/json"
	"io/ioutil"
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

func TestRouter(t *testing.T) {
	// Instantiate the router using the constructor function
	r := newRouter()

	mockServer := httptest.NewServer(r)

	// Test the call to the mock server
	resp, err := http.Get(mockServer.URL + "/hello")

	// Handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want a 200 status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	// Read the body into a bunch of bytes
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// Convert the bytes to a string
	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	// Expect an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// We want to hit the GET /assets/ route to get the index.html file response
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	// We want our status code to be 200
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// Only test the header
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}

func TestGetBirdHandler(t *testing.T) {

	birds = []Bird{
		{"pidgeon", "A flying rat"},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getBirdHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected OK code, received code: %d", status)
	}

	expected := Bird{"pidgeon", "A flying rat"}
	b := []Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", actual, expected)
	}
}

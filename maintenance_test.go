package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a test server with your handler
	testServer := httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()

	// Make a request to the test server
	response, err := http.Get(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Check response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}

	// Read and compare the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "My Modified Title"
	if !strings.Contains(string(body), expectedContent) {
		t.Errorf("Expected response body to contain: %s", expectedContent)
	}
}

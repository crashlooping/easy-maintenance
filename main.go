package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func processHTML(htmlPath string, replacements map[string]string) (string, error) {
	// Read the HTML content from the file
	content, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		return "", err
	}

	// Replace the placeholders with actual values
	modifiedContent := string(content)
	for placeholder, replacement := range replacements {
		modifiedContent = strings.Replace(modifiedContent, placeholder, replacement, -1)
	}

	return modifiedContent, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Define your replacement values
	replacements := map[string]string{
		"{{TITLE}}":   "My Modified Title",
		"{{CONTENT}}": "Hello, World!",
	}

	// Process the HTML content
	modifiedHTML, err := processHTML("html/index.html", replacements)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers and send the modified HTML response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(modifiedHTML))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

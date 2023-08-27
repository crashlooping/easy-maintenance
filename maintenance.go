package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func processHTML(htmlPath string, replacements map[string]string) (string, error) {
	// Open the HTML file
	file, err := os.Open(htmlPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the HTML content from the file
	contentBuilder := strings.Builder{}
	_, err = io.Copy(&contentBuilder, file)
	if err != nil {
		return "", err
	}
	content := contentBuilder.String()

	// Replace the placeholders with actual values
	modifiedContent := content
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

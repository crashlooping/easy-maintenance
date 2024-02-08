package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var BuildTimestamp string

func main() {
	fmt.Printf("Build time: %s\n", BuildTimestamp)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", getIndex)
	e.GET("/ip", getRemoteIP)
	e.GET("/ip/json", getRemoteIPJSON)
	e.GET("/headers", getHeadersJSON)

	e.Logger.Fatal(e.Start(":8080"))
}

func getIndex(c echo.Context) error {
	req := c.Request()

	// Read the index.html file content
	file, err := os.Open("html/index.html")
	if err != nil {
		return err
	}
	defer file.Close()

	contentBuilder := strings.Builder{}
	_, err = io.Copy(&contentBuilder, file)
	if err != nil {
		return err
	}
	htmlContent := contentBuilder.String()

	userAgent := req.Header.Get("User-Agent")
	randomUUID, _ := uuid.NewRandom()
	host := req.Host

	// Replace placeholders in HTML content
	modifiedContent := string(htmlContent)
	modifiedContent = strings.Replace(modifiedContent, "{userAgent}", userAgent, -1)
	modifiedContent = strings.Replace(modifiedContent, "{uuid}", randomUUID.String(), -1)
	modifiedContent = strings.Replace(modifiedContent, "{host}", host, -1)
	modifiedContent = strings.Replace(modifiedContent, "{buildTimestamp}", BuildTimestamp, -1)

	return c.HTMLBlob(http.StatusOK, []byte(modifiedContent))
}

func getRemoteIP(c echo.Context) error {
	req := c.Request()
	xForwardedFor := req.Header.Get("X-Forwarded-For")
	if xForwardedFor == "" {
		return c.String(http.StatusOK, "127.0.0.1")
	}
	return c.String(http.StatusOK, xForwardedFor)
}

func getRemoteIPJSON(c echo.Context) error {
	headers := getIPHeaders(c.Request())
	jsonData, err := json.MarshalIndent(headers, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return c.JSONBlob(http.StatusOK, jsonData)
}

func getHeadersJSON(c echo.Context) error {
	headers := getHeaders(c.Request())
	jsonData, err := json.MarshalIndent(headers, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return c.JSONBlob(http.StatusOK, jsonData)
}

func getIPHeaders(req *http.Request) map[string]string {
	headers := make(map[string]string)
	headers["Timestamp"] = time.Now().Format(time.RFC3339)
	addToMapIfPresent(headers, req, "User-Agent")
	addToMapIfPresent(headers, req, "X-Forwarded-For")
	addToMapIfPresent(headers, req, "X-Real-Ip")
	return headers
}

func getHeaders(req *http.Request) map[string]string {
	headers := make(map[string]string)
	for key, value := range req.Header {
		headers[key] = strings.Join(value, ", ")
	}
	return headers
}

func addToMapIfPresent(m map[string]string, req *http.Request, key string) {
	value := req.Header.Get(key)
	if value != "" {
		m[key] = value
	}
}

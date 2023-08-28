package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getIndex(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check response body content
		expectedContent := "Under Maintenance"

		if !strings.Contains(string(rec.Body.String()), expectedContent) {
			t.Errorf("Expected response body to contain: %s", expectedContent)
		}

	}
}

func TestGetRemoteIP(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/ip", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getRemoteIP(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "127.0.0.1", rec.Body.String()) // Modify this as per your expectations
	}
}

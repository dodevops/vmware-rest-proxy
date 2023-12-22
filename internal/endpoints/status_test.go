package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestStatusEndpoint tests the status endpoint
func TestStatusEndpoint(t *testing.T) {
	r := gin.Default()
	s := StatusEndpoint{}
	s.Register(r)
	req, _ := http.NewRequest("GET", "/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, `{"status":"running"}`, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

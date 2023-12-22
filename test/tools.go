package test

import (
	"fmt"
	"github.com/dodevops/golang-handlerinspector/pkg/builder"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/http/httptest"
	"vmware-rest-proxy/internal/api"
	"vmware-rest-proxy/internal/endpoints"
)

// AUTHTOKEN holds a test token that should be issued and used in all tests
const AUTHTOKEN = "testtoken"

// SessionRule holds a builder Rule for the session api
var SessionRule = builder.NewRule("session").
	WithCondition(builder.HasPath("/api/session")).
	ReturnBodyFromFunction(func(r *http.Request) string {
		if r.Method == "POST" {
			return fmt.Sprintf(`"%s"`, AUTHTOKEN)
		} else {
			return `{"user": "test"}`
		}
	}).
	ReturnHeader("Content-Type", "application/json").
	Build()

// TestRequests is a short helper function to call requests on the build-up endpoints and mock server
// requires a http.Handler and a list of http.Request objects
func TestRequests(handler http.Handler, requests []*http.Request) *httptest.ResponseRecorder {
	s := httptest.NewServer(handler)
	defer s.Close()

	r := resty.New().SetBaseURL(s.URL).SetBasicAuth("test", "test")
	a := api.DefaultVSphereProxyApi{Resty: r}
	g := gin.Default()
	for _, endpoint := range []endpoints.Endpoint{&endpoints.VMSEndpoint{API: a}, &endpoints.HostsEndpoint{API: a}, &endpoints.StatusEndpoint{}, &endpoints.DataStoreEndpoint{API: a}} {
		endpoint.Register(g)
	}

	w := httptest.NewRecorder()

	for _, request := range requests {
		g.ServeHTTP(w, request)
	}
	return w
}

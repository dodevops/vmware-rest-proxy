package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"vmware-rest-proxy/internal/api"
	"vmware-rest-proxy/pkg/handlerinspector"
)

// AUTHTOKEN holds a test token that should be issued and used in all tests
const AUTHTOKEN = "testtoken"

// sessionRule holds a handlerinspector Rule for the session api
var sessionRule = handlerinspector.NewRule("session").
	WithCondition(handlerinspector.HasPath("/api/session")).
	ReturnBodyFromFunction(func(r *http.Request) string {
		if r.Method == "POST" {
			return fmt.Sprintf(`"%s"`, AUTHTOKEN)
		} else {
			return `{"user": "test"}`
		}
	}).
	ReturnHeader("Content-Type", "application/json").
	Build()

// testRequests is a short helper function to call requests on the build-up endpoints and mock server
// requires a http.Handler and a list of http.Request objects
func testRequests(handler http.Handler, requests []*http.Request) *httptest.ResponseRecorder {
	s := httptest.NewServer(handler)
	defer s.Close()

	r := resty.New().SetBaseURL(s.URL).SetBasicAuth("test", "test")
	a := api.DefaultVSphereProxyApi{Resty: r}
	sub := VMSEndpoint{API: a}
	g := gin.Default()
	sub.Register(g)

	w := httptest.NewRecorder()

	for _, request := range requests {
		g.ServeHTTP(w, request)

	}
	return w
}

// TestVMSEndpoint_GetSession checks if the session endpoint is called
func TestVMSEndpoint_GetSession(t *testing.T) {
	b := handlerinspector.NewBuilder().
		WithRule(sessionRule).
		WithRule(
			handlerinspector.NewRule("vms").
				WithCondition(handlerinspector.HasPath("/api/vcenter/vm")).
				WithCondition(handlerinspector.HasMethod("GET")).
				WithCondition(handlerinspector.HasHeader("Vmware-Api-Session-Id", AUTHTOKEN)).
				ReturnBody("[]").
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)

	req, _ := http.NewRequest("GET", "/vms", nil)
	req.SetBasicAuth("test", "test")
	w := testRequests(b.Build(), []*http.Request{req})

	i := handlerinspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestVMSEndpoint_GetVMS checks the vms endpoint
func TestVMSEndpoint_GetVMS(t *testing.T) {
	b := handlerinspector.NewBuilder().
		WithRule(sessionRule).
		WithRule(
			handlerinspector.NewRule("vms").
				WithCondition(handlerinspector.HasPath("/api/vcenter/vm")).
				WithCondition(handlerinspector.HasMethod("GET")).
				WithCondition(handlerinspector.HasHeader("Vmware-Api-Session-Id", AUTHTOKEN)).
				ReturnBody(`[{"VM": "1", "Name": "test1"}, {"VM": "2", "Name": "test2"}]`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)
	req, _ := http.NewRequest("GET", "/vms", nil)
	req.SetBasicAuth("test", "test")
	w := testRequests(b.Build(), []*http.Request{req})

	type resp struct {
		VMS struct {
			Count int              `json:"count"`
			VMS   []api.VMResponse `json:"vms"`
		} `json:"vms"`
	}
	var r resp
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := handlerinspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, r.VMS.Count, 2)
	assert.Equal(t, len(r.VMS.VMS), 2)
	assert.Equal(t, r.VMS.VMS[0].VM, "1")
	assert.Equal(t, r.VMS.VMS[0].Name, "test1")
	assert.Equal(t, r.VMS.VMS[1].VM, "2")
	assert.Equal(t, r.VMS.VMS[1].Name, "test2")
}

// TestVMSEndpoint_GetVMTags checks the /vms/tags endpoint
func TestVMSEndpoint_GetVMTags(t *testing.T) {
	b := handlerinspector.NewBuilder().
		WithRule(sessionRule).
		WithRule(
			handlerinspector.NewRule("list-associated-tags").
				WithCondition(handlerinspector.HasPath("/api/cis/tagging/tag-association")).
				WithCondition(handlerinspector.HasMethod("POST")).
				WithCondition(handlerinspector.HasHeader("Vmware-Api-Session-Id", AUTHTOKEN)).
				WithCondition(handlerinspector.HasQueryParam("action", "list-attached-tags")).
				WithCondition(handlerinspector.HasBody(`{"object_id":{"id":"1","type":"VirtualMachine"}}`)).
				ReturnBody(`["1", "2"]`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			handlerinspector.NewRule("tag-data-1").
				WithCondition(handlerinspector.HasPath("/api/cis/tagging/tag/1")).
				WithCondition(handlerinspector.HasMethod("GET")).
				WithCondition(handlerinspector.HasHeader("Vmware-Api-Session-Id", AUTHTOKEN)).
				ReturnBody(`{"category_id": "1", "name": "testtag1"}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			handlerinspector.NewRule("tag-data-2").
				WithCondition(handlerinspector.HasPath("/api/cis/tagging/tag/2")).
				WithCondition(handlerinspector.HasMethod("GET")).
				WithCondition(handlerinspector.HasHeader("Vmware-Api-Session-Id", AUTHTOKEN)).
				ReturnBody(`{"category_id": "2", "name": "testtag2"}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			handlerinspector.NewRule("tag-category-1").
				WithCondition(handlerinspector.HasPath("/api/cis/tagging/category/1")).
				WithCondition(handlerinspector.HasMethod("GET")).
				WithCondition(handlerinspector.HasHeader("Vmware-Api-Session-Id", AUTHTOKEN)).
				ReturnBody(`{"name": "testcategory1"}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			handlerinspector.NewRule("tag-category-2").
				WithCondition(handlerinspector.HasPath("/api/cis/tagging/category/2")).
				WithCondition(handlerinspector.HasMethod("GET")).
				WithCondition(handlerinspector.HasHeader("Vmware-Api-Session-Id", AUTHTOKEN)).
				ReturnBody(`{"name": "testcategory2"}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)
	req, _ := http.NewRequest("GET", "/vms/1/tags", nil)
	req.SetBasicAuth("test", "test")
	w := testRequests(b.Build(), []*http.Request{req})

	type resp struct {
		Tags struct {
			Count int         `json:"count"`
			Tags  []api.VMTag `json:"tags"`
		} `json:"tags"`
	}
	var r resp
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := handlerinspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, r.Tags.Count, 2)
	assert.Equal(t, len(r.Tags.Tags), 2)
	assert.Equal(t, r.Tags.Tags[0].Category, "testcategory1")
	assert.Equal(t, r.Tags.Tags[0].Value, "testtag1")
	assert.Equal(t, r.Tags.Tags[1].Category, "testcategory2")
	assert.Equal(t, r.Tags.Tags[1].Value, "testtag2")
}

// TestVMSEndpoint_GetFQDN checks the vm/fqdn endpoint
func TestVMSEndpoint_GetFQDN(t *testing.T) {
	b := handlerinspector.NewBuilder().
		WithRule(sessionRule).
		WithRule(
			handlerinspector.NewRule("get-fqdm").
				WithCondition(handlerinspector.HasPath("/api/vcenter/vm/1/guest/networking")).
				WithCondition(handlerinspector.HasMethod("GET")).
				WithCondition(handlerinspector.HasHeader("Vmware-Api-Session-Id", AUTHTOKEN)).
				ReturnBody(`{"dns_values":{"domain_name":"example.com","host_name":"test"}}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)
	req, _ := http.NewRequest("GET", "/vms/1/fqdn", nil)
	req.SetBasicAuth("test", "test")
	w := testRequests(b.Build(), []*http.Request{req})

	type resp struct {
		FQDN string
	}
	var r resp
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := handlerinspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, r.FQDN, "test.example.com")
}

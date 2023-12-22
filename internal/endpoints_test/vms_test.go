package endpoints

import (
	"encoding/json"
	"github.com/dodevops/golang-handlerinspector/pkg/builder"
	"github.com/dodevops/golang-handlerinspector/pkg/inspector"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
	"vmware-rest-proxy/internal/api"
	"vmware-rest-proxy/test"
)

// TestVMSEndpoint_GetSession checks if the session endpoint is called
func TestVMSEndpoint_GetSession(t *testing.T) {
	b := builder.NewBuilder().
		WithRule(test.SessionRule).
		WithRule(
			builder.NewRule("vms").
				WithCondition(builder.HasPath("/api/vcenter/vm")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody("[]").
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)

	req, _ := http.NewRequest("GET", "/vms", nil)
	req.SetBasicAuth("test", "test")
	w := test.TestRequests(b.Build(), []*http.Request{req})

	i := inspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestVMSEndpoint_GetVMS checks the vms endpoint
func TestVMSEndpoint_GetVMS(t *testing.T) {
	b := builder.NewBuilder().
		WithRule(test.SessionRule).
		WithRule(
			builder.NewRule("vms").
				WithCondition(builder.HasPath("/api/vcenter/vm")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`[{"vm": "1", "name": "test1", "power_state": "POWERED_OFF"}, {"vm": "2", "name": "test2", "power_state": "POWERED_ON"}]`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)
	req, _ := http.NewRequest("GET", "/vms", nil)
	req.SetBasicAuth("test", "test")
	w := test.TestRequests(b.Build(), []*http.Request{req})

	type resp struct {
		VMS struct {
			Count int      `json:"count"`
			VMS   []api.VM `json:"vms"`
		} `json:"vms"`
	}
	var r resp
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := inspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, r.VMS.Count, 2)
	assert.Equal(t, len(r.VMS.VMS), 2)
	assert.Equal(t, r.VMS.VMS[0].VM, "1")
	assert.Equal(t, r.VMS.VMS[0].Name, "test1")
	assert.Equal(t, r.VMS.VMS[0].PowerState, "POWERED_OFF")
	assert.Equal(t, r.VMS.VMS[1].VM, "2")
	assert.Equal(t, r.VMS.VMS[1].Name, "test2")
	assert.Equal(t, r.VMS.VMS[1].PowerState, "POWERED_ON")
}

// TestVMSEndpoint_GetVMTags checks the /vms/tags endpoint
func TestVMSEndpoint_GetVMTags(t *testing.T) {
	b := builder.NewBuilder().
		WithRule(test.SessionRule).
		WithRule(
			builder.NewRule("list-associated-tags").
				WithCondition(builder.HasPath("/api/cis/tagging/tag-association")).
				WithCondition(builder.HasMethod("POST")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				WithCondition(builder.HasQueryParam("action", "list-attached-tags")).
				WithCondition(builder.HasBody(`{"object_id":{"id":"1","type":"VirtualMachine"}}`)).
				ReturnBody(`["1", "2"]`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			builder.NewRule("tag-data-1").
				WithCondition(builder.HasPath("/api/cis/tagging/tag/1")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`{"category_id": "1", "name": "testtag1"}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			builder.NewRule("tag-data-2").
				WithCondition(builder.HasPath("/api/cis/tagging/tag/2")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`{"category_id": "2", "name": "testtag2"}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			builder.NewRule("tag-category-1").
				WithCondition(builder.HasPath("/api/cis/tagging/category/1")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`{"name": "testcategory1"}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			builder.NewRule("tag-category-2").
				WithCondition(builder.HasPath("/api/cis/tagging/category/2")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`{"name": "testcategory2"}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)
	req, _ := http.NewRequest("GET", "/vms/1/tags", nil)
	req.SetBasicAuth("test", "test")
	w := test.TestRequests(b.Build(), []*http.Request{req})

	type resp struct {
		Tags struct {
			Count int         `json:"count"`
			Tags  []api.VMTag `json:"tags"`
		} `json:"tags"`
	}
	var r resp
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := inspector.NewInspector(b)
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
	b := builder.NewBuilder().
		WithRule(test.SessionRule).
		WithRule(
			builder.NewRule("get-fqdn").
				WithCondition(builder.HasPath("/api/vcenter/vm/1/guest/networking")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`{"dns_values":{"domain_name":"example.com","host_name":"test"}}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)
	req, _ := http.NewRequest("GET", "/vms/1/fqdn", nil)
	req.SetBasicAuth("test", "test")
	w := test.TestRequests(b.Build(), []*http.Request{req})

	type resp struct {
		FQDN string
	}
	var r resp
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := inspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, r.FQDN, "test.example.com")
}

func TestVMSEndpoint_GetVMInfo(t *testing.T) {
	b := builder.NewBuilder().
		WithRule(test.SessionRule).
		WithRule(
			builder.NewRule("get-vm").
				WithCondition(builder.HasPath("/api/vcenter/vm/1")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`{"name":"test", "cpu": {"cores_per_socket": 2, "count": 2}, "memory": {"size_MiB": 200}}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		).
		WithRule(
			builder.NewRule("get-vm").
				WithCondition(builder.HasPath("/api/vcenter/vm/1/guest/local-filesystem")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`{"/": {"capacity": 100, "free_space": 0}, "/opt": {"capacity": 2000, "free_space": 20}}`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)

	req, _ := http.NewRequest("GET", "/vms/1/info", nil)
	req.SetBasicAuth("test", "test")
	w := test.TestRequests(b.Build(), []*http.Request{req})

	var r api.VMInfo
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := inspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, r.Name, "test")
	assert.Equal(t, r.CPUCores, 4)
	assert.Equal(t, r.ProvisionedRAM, 200)
	assert.Equal(t, r.ProvisionedStorage, 100+2000)
	assert.Equal(t, r.UsedStorage, 100+2000-20)
}

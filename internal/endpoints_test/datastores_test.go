package endpoints_test

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

func TestDatastoresEndpoint_GetDatastores(t *testing.T) {
	b := builder.NewBuilder().
		WithRule(test.SessionRule).
		WithRule(
			builder.NewRule("datastores").
				WithCondition(builder.HasPath("/api/vcenter/datastore")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`[{"datastore": "1", "name": "test1", "type": "VMFS", "capacity": 100, "free_space": 0}, {"datastore": "2", "name": "test2", "type": "NFS", "capacity": 1000, "free_space": 10}]`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)
	req, _ := http.NewRequest("GET", "/datastores", nil)
	req.SetBasicAuth("test", "test")
	w := test.TestRequests(b.Build(), []*http.Request{req})

	type resp struct {
		Datastores struct {
			Count      int             `json:"count"`
			Datastores []api.Datastore `json:"datastores"`
		} `json:"datastores"`
	}
	var r resp
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := inspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, r.Datastores.Count, 2)
	assert.Equal(t, len(r.Datastores.Datastores), 2)
	assert.Equal(t, r.Datastores.Datastores[0].Datastore, "1")
	assert.Equal(t, r.Datastores.Datastores[0].Name, "test1")
	assert.Equal(t, r.Datastores.Datastores[0].Type, "VMFS")
	assert.Equal(t, r.Datastores.Datastores[0].Capacity, 100)
	assert.Equal(t, r.Datastores.Datastores[0].FreeSpace, 0)
	assert.Equal(t, r.Datastores.Datastores[1].Datastore, "2")
	assert.Equal(t, r.Datastores.Datastores[1].Name, "test2")
	assert.Equal(t, r.Datastores.Datastores[1].Type, "NFS")
	assert.Equal(t, r.Datastores.Datastores[1].Capacity, 1000)
	assert.Equal(t, r.Datastores.Datastores[1].FreeSpace, 10)
}

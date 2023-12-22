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

// TestHostsEndpoint_GetVMS checks the hosts endpoint
func TestHostsEndpoint_GetHosts(t *testing.T) {
	b := builder.NewBuilder().
		WithRule(test.SessionRule).
		WithRule(
			builder.NewRule("hosts").
				WithCondition(builder.HasPath("/api/vcenter/host")).
				WithCondition(builder.HasMethod("GET")).
				WithCondition(builder.HasHeader("Vmware-Api-Session-Id", test.AUTHTOKEN)).
				ReturnBody(`[{"host": "1", "name": "test1", "power_state": "POWERED_OFF", "connection_state": "CONNECTED"}, {"host": "2", "name": "test2", "power_state": "POWERED_ON", "connection_state": "DISCONNECTED"}]`).
				ReturnHeader("Content-Type", "application/json").
				Build(),
		)
	req, _ := http.NewRequest("GET", "/hosts", nil)
	req.SetBasicAuth("test", "test")
	w := test.TestRequests(b.Build(), []*http.Request{req})

	type resp struct {
		Hosts struct {
			Count int        `json:"count"`
			Hosts []api.Host `json:"hosts"`
		} `json:"hosts"`
	}
	var r resp
	err := json.NewDecoder(w.Body).Decode(&r)
	assert.Equal(t, err, nil)

	i := inspector.NewInspector(b)
	assert.Equal(t, i.Failed(), false)
	assert.Equal(t, i.AllWereCalled(), true)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, r.Hosts.Count, 2)
	assert.Equal(t, len(r.Hosts.Hosts), 2)
	assert.Equal(t, r.Hosts.Hosts[0].Host, "1")
	assert.Equal(t, r.Hosts.Hosts[0].Name, "test1")
	assert.Equal(t, r.Hosts.Hosts[0].PowerState, "POWERED_OFF")
	assert.Equal(t, r.Hosts.Hosts[0].ConnectionState, "CONNECTED")
	assert.Equal(t, r.Hosts.Hosts[1].Host, "2")
	assert.Equal(t, r.Hosts.Hosts[1].Name, "test2")
	assert.Equal(t, r.Hosts.Hosts[1].PowerState, "POWERED_ON")
	assert.Equal(t, r.Hosts.Hosts[1].ConnectionState, "DISCONNECTED")
}

package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vmware-rest-proxy/internal/api"
)

type HostsEndpoint struct {
	API api.VSphereProxyApi
}

func (H HostsEndpoint) Register(engine *gin.Engine) {
	engine.GET("/hosts", H.getHosts)
}

type HostsResult struct {
	Count int        `json:"count"`
	Hosts []api.Host `json:"hosts"`
}

type Hosts struct {
	Hosts HostsResult `json:"hosts"`
}

// @Summary Retrieve a list of ESXi hosts
// @Description Fetches a list of registered ESXi hosts in the vCenter
// @Tags host
// @Produce json
// @Security BasicAuth
// @Success 200 {object} Hosts
// @Failure 401 "Authorization is required"
// @Failure 400 "Invalid request"
// @Router /hosts [get]
func (H HostsEndpoint) getHosts(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		if hosts, err := H.API.GetHosts(r.Username, r.Password); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting hosts: %s", err),
			})
		} else {
			context.JSON(200, gin.H{
				"hosts": Hosts{Hosts: HostsResult{
					Count: len(hosts),
					Hosts: hosts,
				}},
			})
		}
	}
}

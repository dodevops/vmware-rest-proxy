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

func (H HostsEndpoint) getHosts(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		if hosts, err := H.API.GetHosts(r.Username, r.Password); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting hosts: %s", err),
			})
		} else {
			context.JSON(200, gin.H{
				"hosts": gin.H{
					"count": len(hosts),
					"hosts": hosts,
				},
			})
		}
	}
}

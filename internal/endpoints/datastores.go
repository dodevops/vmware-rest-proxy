package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vmware-rest-proxy/internal/api"
)

type DataStoreEndpoint struct {
	API api.VSphereProxyApi
}

func (d DataStoreEndpoint) Register(engine *gin.Engine) {
	engine.GET("/datastores", d.getDatastores)
}

func (d DataStoreEndpoint) getDatastores(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		if datastores, err := d.API.GetDatastores(r.Username, r.Password); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting datastores: %s", err),
			})
		} else {
			context.JSON(200, gin.H{
				"datastores": gin.H{
					"count":      len(datastores),
					"datastores": datastores,
				},
			})
		}
	}
}

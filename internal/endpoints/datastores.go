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

type DatastoresResult struct {
	Count      int             `json:"count"`
	Datastores []api.Datastore `json:"datastores"`
}

type Datastores struct {
	Datastores DatastoresResult `json:"datastores"`
}

// @Summary Retrieve a list of datastores
// @Description Fetches a list of registered datastores in the vCenter
// @Tags datastore
// @Produce json
// @Security BasicAuth
// @Success 200 {object} Datastores
// @Failure 401 "Authorization is required"
// @Failure 400 "Invalid request"
// @Router /datastores [get]
func (d DataStoreEndpoint) getDatastores(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		if datastores, err := d.API.GetDatastores(r.Username, r.Password); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting datastores: %s", err),
			})
		} else {
			context.JSON(200, Datastores{
				Datastores: DatastoresResult{
					Count:      len(datastores),
					Datastores: datastores,
				},
			})
		}
	}
}

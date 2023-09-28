package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vmware-rest-proxy/internal"
	"vmware-rest-proxy/internal/api"
)

// The VMSEndpoint returns information about VMs in vSphere
type VMSEndpoint struct {
	config internal.Config
}

var _ Endpoint = &VMSEndpoint{}

// VMBinding is used for the path binding of the :vm param
type VMBinding struct {
	ID string `uri:"vm" binding:"required"`
}

func (V *VMSEndpoint) Register(engine *gin.Engine, config internal.Config) {
	V.config = config
	engine.GET("/vms", V.getVMS)
	engine.GET("/vms/:vm/tags", V.getVMTags)
}

// getVMS exposes all vms of the vCenter at /vms
func (V *VMSEndpoint) getVMS(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		if vms, err := api.GetVMs(V.config, r.Username, r.Password); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting VMs: %s", err),
			})
		} else {
			context.JSON(200, gin.H{
				"vms": gin.H{
					"count": len(vms),
					"vms":   vms,
				},
			})
		}
	}
}

// getVMTags exposes a list of tags associated with a vm at /vms/:vm/tags
func (V *VMSEndpoint) getVMTags(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		var vm VMBinding
		if err := context.ShouldBindUri(&vm); err != nil {
			context.AbortWithStatusJSON(400, gin.H{
				"error": fmt.Sprintf("Missing VM id in path: %s", err),
			})
			return
		}
		if tags, err := api.GetVMTags(V.config, r.Username, r.Password, vm.ID); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting tags: %s", err),
			})
		} else {
			context.JSON(200, gin.H{
				"tags": gin.H{
					"count": len(tags),
					"tags":  tags,
				},
			})
		}
	}
}

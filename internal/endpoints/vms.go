package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vmware-rest-proxy/internal/api"
)

// The VMSEndpoint returns information about VMs in vSphere
type VMSEndpoint struct {
	API api.VSphereProxyApi
}

var _ Endpoint = &VMSEndpoint{}

// VMBinding is used for the path binding of the :vm param
type VMBinding struct {
	ID string `uri:"vm" binding:"required"`
}

func (V *VMSEndpoint) Register(engine *gin.Engine) {
	engine.GET("/vms", V.getVMS)
	engine.GET("/vms/:vm/tags", V.getVMTags)
	engine.GET("/vms/:vm/fqdn", V.getFQDN)
	engine.GET("/vms/:vm/info", V.getVMInfo)
}

// getVMS exposes all VMS of the vCenter at /VMS
func (V *VMSEndpoint) getVMS(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		if vms, err := V.API.GetVMs(r.Username, r.Password); err != nil {
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

// getVMTags exposes a list of tags associated with a vm at /VMS/:vm/tags
func (V *VMSEndpoint) getVMTags(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		var vm VMBinding
		if err := context.ShouldBindUri(&vm); err != nil {
			context.AbortWithStatusJSON(400, gin.H{
				"error": fmt.Sprintf("Missing VM id in path: %s", err),
			})
			return
		}
		if tags, err := V.API.GetVMTags(r.Username, r.Password, vm.ID); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting tags: %s", err),
			})
		} else {
			context.JSON(200, gin.H{
				"tags": gin.H{
					"Count": len(tags),
					"tags":  tags,
				},
			})
		}
	}
}

func (V *VMSEndpoint) getFQDN(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		var vm VMBinding
		if err := context.ShouldBindUri(&vm); err != nil {
			context.AbortWithStatusJSON(400, gin.H{
				"error": fmt.Sprintf("Missing VM id in path: %s", err),
			})
			return
		}
		if fqdn, err := V.API.GetFQDN(r.Username, r.Password, vm.ID); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting tags: %s", err),
			})
		} else {
			context.JSON(200, gin.H{
				"fqdn": fqdn,
			})
		}
	}
}

func (V *VMSEndpoint) getVMInfo(context *gin.Context) {
	if r, ok := HandleRequest(context); ok {
		var vm VMBinding
		if err := context.ShouldBindUri(&vm); err != nil {
			context.AbortWithStatusJSON(400, gin.H{
				"error": fmt.Sprintf("Missing VM id in path: %s", err),
			})
			return
		}
		if vmInfo, err := V.API.GetVMInfo(r.Username, r.Password, vm.ID); err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": fmt.Sprintf("Error getting vm info: %s", err),
			})
		} else {
			context.JSON(200, vmInfo)
		}
	}
}

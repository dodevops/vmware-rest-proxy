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

type VMSResult struct {
	Count int      `json:"count"`
	VMS   []api.VM `json:"vms"`
}
type VMS struct {
	VMS VMSResult `json:"vms"`
}

// @Summary Retrieve a list of all vms
// @Description Fetches a list of vms from the vCenter
// @Tags vm
// @Produce json
// @Security BasicAuth
// @Success 200 {object} VMS
// @Failure 401 "Authorization is required"
// @Failure 400 "Invalid request"
// @Router /vms [get]
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

type TagsResult struct {
	Count int         `json:"count"`
	Tags  []api.VMTag `json:"tags"`
}

type Tags struct {
	Tags TagsResult `json:"tags"`
}

// @Summary Retrieve tags
// @Description Retrieve tags  and their categories for a vm
// @Param id path string true "ID of VM"
// @Tags vm
// @Produce json
// @Security BasicAuth
// @Success 200 {object} Tags
// @Failure 401 "Authorization is required"
// @Failure 400 "Invalid request"
// @Router /vms/{id}/tags [get]
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

type FQDN struct {
	FQDN string `json:"fqdn"`
}

// @Summary Get fqdn of VM
// @Description Try to find out the fqdn of the given VM using the guest tools
// @Param id path string true "ID of VM"
// @Tags vm
// @Produce json
// @Security BasicAuth
// @Success 200 {object} FQDN
// @Failure 401 "Authorization is required"
// @Failure 400 "Invalid request"
// @Router /vms/{id}/fqdn [get]
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

// @Summary Get informational data about a VM
// @Description Find out some information about a VM and return them
// @Param id path string true "ID of VM"
// @Tags vm
// @Produce json
// @Security BasicAuth
// @Success 200 {object} api.VMInfo
// @Failure 401 "Authorization is required"
// @Failure 400 "Invalid request"
// @Router /vms/{id}/info [get]
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

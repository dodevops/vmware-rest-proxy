package endpoints

import (
	"github.com/gin-gonic/gin"
)

// The StatusEndpoint returns information about VMs in vSphere
type StatusEndpoint struct{}

var _ Endpoint = &StatusEndpoint{}

type Status struct {
	Status string `json:"status"`
}

// Register registers the status endpoint
// @Summary Checks whether the service is running
// @Description Just responses with a 200 to signal that the service is running
// @Produce json
// @Success 200 {object} Status
// @Router /status [get]
func (StatusEndpoint) Register(engine *gin.Engine) {
	engine.GET("/status", func(context *gin.Context) {
		context.JSON(200, Status{Status: "running"})
	})
}

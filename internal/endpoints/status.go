package endpoints

import (
	"github.com/gin-gonic/gin"
)

// The StatusEndpoint returns information about VMs in vSphere
type StatusEndpoint struct{}

var _ Endpoint = &StatusEndpoint{}

func (StatusEndpoint) Register(engine *gin.Engine) {
	engine.GET("/status", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "running",
		})
	})
}

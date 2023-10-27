package endpoints

import (
	"github.com/gin-gonic/gin"
	"vmware-rest-proxy/internal"
)

// The StatusEndpoint returns information about VMs in vSphere
type StatusEndpoint struct {
	config internal.Config
}

var _ Endpoint = &StatusEndpoint{}

func (StatusEndpoint) Register(engine *gin.Engine, _ internal.Config) {
	engine.GET("/status", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "running",
		})
	})
}

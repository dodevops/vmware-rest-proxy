package api

//go:generate mockgen -source api.go -package test -destination ../../test/api.go

import (
	"github.com/go-resty/resty/v2"
)

type VSphereProxyApi interface {

	// GetSession returns the vmware session id to be used by other requests
	GetSession(username string, password string) (string, error)

	// GetVMs returns all VMs from the VM endpoint
	GetVMs(username string, password string) ([]VMResponse, error)

	// GetVMTags retrieves a list of tags associated with the given vm
	GetVMTags(username string, password string, VMID string) ([]VMTag, error)

	// GetFQDN uses the VMware guest tools to get the fqdn of a VM (if possible)
	GetFQDN(username string, password string, VMID string) (string, error)
}

type DefaultVSphereProxyApi struct {
	Resty *resty.Client
}

var _ VSphereProxyApi = DefaultVSphereProxyApi{}

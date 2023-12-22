package api

import (
	"github.com/go-resty/resty/v2"
)

// The VSphereProxyApi interface describes the API available to the endpoints
type VSphereProxyApi interface {

	// GetSession returns the vmware session id to be used by other requests
	GetSession(username string, password string) (string, error)

	// GetVMs returns all VMs from the VM endpoint
	GetVMs(username string, password string) ([]VM, error)

	// GetVMTags retrieves a list of tags associated with the given vm
	GetVMTags(username string, password string, VMID string) ([]VMTag, error)

	// GetFQDN uses the VMware guest tools to get the fqdn of a VM (if possible)
	GetFQDN(username string, password string, VMID string) (string, error)

	// GetHosts retrieves a list of ESXi hosts from the vCenter
	GetHosts(username string, password string) ([]Host, error)

	// GetDatastores retrieves a list of datastores from the vCenter
	GetDatastores(username string, password string) ([]Datastore, error)
}

// DefaultVSphereProxyApi is the default API implementation
type DefaultVSphereProxyApi struct {
	Resty *resty.Client
}

var _ VSphereProxyApi = DefaultVSphereProxyApi{}

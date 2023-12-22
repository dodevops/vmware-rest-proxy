package api

// VM represents a virtual machine in the vCenter as described in https://developer.vmware.com/apis/vsphere-automation/v8.0U1/vcenter/data-structures/VM/Summary/
type VM struct {
	VM         string `json:"vm"`
	Name       string `json:"name"`
	PowerState string `json:"power_state"`
}

// VMTag holds a tag from vSphere
type VMTag struct {
	// Value holds the value of the tag
	Value string `json:"value"`
	// Category holds the tag category
	Category string `json:"category"`
}

// Host represents a host in the vCenter as described in https://developer.vmware.com/apis/vsphere-automation/v8.0U1/vcenter/data-structures/Host/Summary/
type Host struct {
	Host            string `json:"host"`
	Name            string `json:"name"`
	PowerState      string `json:"power_state"`
	ConnectionState string `json:"connection_state"`
}

// Datastore represents a host in the vCenter as described in https://developer.vmware.com/apis/vsphere-automation/v8.0U1/vcenter/data-structures/Datastore/Summary/
type Datastore struct {
	Datastore string `json:"datastore"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Capacity  int    `json:"capacity"`
	FreeSpace int    `json:"free_space"`
}

package api

// VM is the value in response from the VM endpoint
type VM struct {
	VM         string `json:"vm"`
	Name       string `json:"name"`
	PowerState string `json:"power_state"`
}

// VMTag holds a tag from vSphere
type VMTag struct {
	Value    string `json:"value"`
	Category string `json:"category"`
}

type Host struct {
	Host            string `json:"host"`
	Name            string `json:"name"`
	PowerState      string `json:"power_state"`
	ConnectionState string `json:"connection_state"`
}

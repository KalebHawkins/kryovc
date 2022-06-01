package kvclient

import "fmt"

type VmwareError struct {
	ErrorType string `json:"error_type"`
	Messages  []struct {
		Args           []string `json:"args"`
		DefaultMessage string   `json:"default_message"`
		Id             string   `json:"id"`
	} `json:"messages"`
}

func (vme *VmwareError) String() string {
	return fmt.Sprintf("ErrorType: %s\nDefault Message: %s\nId: %s", vme.ErrorType, vme.Messages[0].DefaultMessage, vme.Messages[0].Id)
}

func (vme VmwareError) Error() string {
	return vme.String()
}

const (
	PoweredOff = "POWERED_OFF"
	PoweredOn  = "POWERED_ON"
	Suspended  = "SUSPENDED"
)

type VMSummary struct {
	CPUCount      int    `json:"cpu_count,omitempty"`
	MemorySizeMiB int    `json:"memory_size_MiB,omitempty"`
	Name          string `json:"name"`
	PowerState    string `json:"power_state"`
	VM            string `json:"vm"`
}

func (vms VMSummary) String() string {
	return fmt.Sprintf("Name: %s\nPowerState: %s\nCPUCount: %d\nMemorySizeMib: %d\nVM UUID: %s\n",
		vms.Name,
		vms.PowerState,
		vms.CPUCount,
		vms.MemorySizeMiB,
		vms.VM)
}

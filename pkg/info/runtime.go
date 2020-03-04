package info

import (
	"os"
)

// Runtime ...
type Runtime struct {
	Hostname string              `json:"hostname"`
	CPUInfo  []map[string]string `json:"cpuInfo"`
	MemInfo  map[string]int64    `json:"memInfo"`
}

// GetRuntime ...
func GetRuntime() Runtime {
	i := Runtime{}
	i.Hostname, _ = os.Hostname()

	return i
}

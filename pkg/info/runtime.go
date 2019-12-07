package info

import (
	"os"
)

type Runtime struct {
	Hostname string              `json:"hostname"`
	CpuInfo  []map[string]string `json:"cpuInfo"`
	MemInfo  map[string]int64    `json:"memInfo"`
}

func GetRuntime() Runtime {
	i := Runtime{}
	i.Hostname, _ = os.Hostname()

	return i
}

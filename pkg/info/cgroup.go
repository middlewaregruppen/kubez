package info

// https://kernel.googlesource.com/pub/scm/linux/kernel/git/glommer/memcg/+/cpu_stat/Documentation/cgroups/cpu.txt

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type CGroup struct {
	// The weight of each group living in the same hierarchy, that
	// translates into the amount of CPU it is expected to get
	CpuShares int64 `json:"cpuShares"`
	// The duration in microseconds of each scheduler period, for
	// bandwidth decisions
	CpuCfsPeriod int64 `json:"cpuPeriod"`
	// The maximum time in microseconds during each cfs_period_us
	// in for the current group will be allowed to run.
	CpuCfsQuota int64 `json:"cpuQuota"`
	// Statistics about the bandwidth controls. No data will be presented
	// if cpu.cfs_quota_us is not set. The file presents three numbers:
	// nr_periods: how many full periods have been elapsed.
	// nr_throttled: number of times we exausted the full allowed bandwidth
	// throttled_time: total time the tasks were not run due to being overquota
	CpuStatNrPeriods     int64 `json:"cpuNumberPeriods"`
	CpuStatNrThrottled   int64 `json:"cpuNumberThrottled"`
	CpuStatThrottledTime int64 `json:"cpuTimeThrottled"`
	// The aggregate CPU time, in nanoseconds, consumed by all tasks
	// in this group.
	CpuAcctUsage int64 `json:"cpuAcctUsage"`
	// The CPU time, in nanoseconds, consumed by all tasks in
	// this group, separated by CPU.
	CpuAcctUsagePerCpu []int64 `json:"cpuacctUsage_percpu"`

	MemoryStat         map[string]int64 `json:"memoryStat"`
	MemoryLimitInBytes int64            `json:"memoryLimit"`
	MemoryUsageInBytes int64            `json:"memoryUsage"`
}

func GetCgroup() *CGroup {
	c := &CGroup{}

	c.CpuShares, _ = getIntFromFile("/sys/fs/cgroup/cpu/cpu.shares")
	c.CpuCfsPeriod, _ = getIntFromFile("/sys/fs/cgroup/cpu/cpu.cfs_period_us")
	c.CpuCfsQuota, _ = getIntFromFile("/sys/fs/cgroup/cpu/cpu.cfs_quota_us")
	c.CpuAcctUsage, _ = getIntFromFile("/sys/fs/cgroup/cpu/cpuacct.usage")

	cpustat, _ := getIntMapFromFile("/sys/fs/cgroup/cpu/cpu.stat")
	c.CpuStatNrPeriods = cpustat["nr_periods"]
	c.CpuStatNrThrottled = cpustat["nr_throttled"]
	c.CpuStatThrottledTime = cpustat["throttled_time"]

	c.CpuAcctUsagePerCpu, _ = getIntArrayFromFile("/sys/fs/cgroup/cpu/cpuacct.usage_percpu")

	c.MemoryStat, _ = getIntMapFromFile("/sys/fs/cgroup/memory/memory.stat")
	c.MemoryLimitInBytes, _ = getIntFromFile("/sys/fs/cgroup/memory/memory.limit_in_bytes")
	c.MemoryUsageInBytes, _ = getIntFromFile("/sys/fs/cgroup/memory/memory.usage_in_bytes")

	return c
}

// getIntFromFile returns an integer that are stored in a file.
func getIntFromFile(file string) (int64, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(strings.TrimSuffix(string(f), "\n"), 10, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func getIntMapFromFile(file string) (map[string]int64, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lns := strings.Split(string(f), "\n")
	m := make(map[string]int64, len(lns))
	for _, ln := range lns {
		kv := strings.Split(ln, " ")
		if len(kv) != 2 {
			continue
		}
		v, _ := strconv.ParseInt(strings.TrimSuffix(string(kv[1]), "\n"), 10, 64)
		m[kv[0]] = v
	}
	return m, nil
}

func getIntArrayFromFile(file string) ([]int64, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	a := strings.Split(string(f), " ")

	var rv []int64
	for _, ai := range a {
		v, _ := strconv.ParseInt(ai, 10, 64)
		rv = append(rv, v)
	}
	return rv, nil
}

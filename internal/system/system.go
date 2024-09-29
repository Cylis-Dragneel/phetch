package system

import (
	"os"
	"runtime"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemInfo struct {
	OS            string
	Distribution  string
	Version       string
	Hostname      string
	Kernel        string
	Uptime        uint64
	Architecture  string
	TotalMemory   uint64
	UsedMemoryPct float64
}

func GetSystemInfo() SystemInfo {
	memory, _ := mem.VirtualMemory()
	hostname, _ := os.Hostname()
	kernel, _ := host.KernelVersion()
	distro, _, version, _ := host.PlatformInformation()
	uptime, _ := host.Uptime()
	architecture, _ := host.KernelArch()

	return SystemInfo{
		OS:            runtime.GOOS,
		Distribution:  distro,
		Version:       version,
		Hostname:      hostname,
		Kernel:        kernel,
		Uptime:        uptime,
		Architecture:  architecture,
		TotalMemory:   memory.Total,
		UsedMemoryPct: memory.UsedPercent,
	}
}

package main
import (
	"fmt"
  "runtime"
  "github.com/shirou/gopsutil/v3/mem"
  "github.com/shirou/gopsutil/host"
  "os"
  "strings"
)
func main() {
  memory, _ := mem.VirtualMemory()
  hostname, _ := os.Hostname()
  kernel, _ := host.KernelVersion()
  distro, _, version, _ := host.PlatformInformation()
  uptime, _ := host.Uptime()
  uptime = uptime / 60 / 60
  architecture, _ := host.KernelArch()
  
	fmt.Println("System Information:")
	fmt.Printf("OS: %s\n", runtime.GOOS)
  fmt.Printf("Distribution: %s %s\n", strings.ToUpper(distro), version)
  fmt.Printf("Hostname: %s\n", hostname)
  fmt.Printf("Kernel: %s\n", kernel)
  fmt.Printf("Uptime: %dH\n", uptime)
	fmt.Printf("Architecture: %s\n", architecture)
  fmt.Printf("Total memory: %dGB\n", (memory.Total/1024/1024/1024)+1)
  fmt.Printf("Used memory percentage: %s%\n", fmt.Sprintf("%.0f", memory.UsedPercent))
}

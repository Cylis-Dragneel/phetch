package display

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cylis-Dragneel/phetch/internal/system"
)

func ShowArt(path string) error {
	if path == "" {
		return nil // No art to display
	}

	ext := filepath.Ext(path)
	switch ext {
	case ".txt":
		return displayASCII(path)
	case ".png", ".jpg", ".jpeg":
		return displayImage(path)
	default:
		return fmt.Errorf("unsupported file type: %s", ext)
	}
}

func displayASCII(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}

func displayImage(path string) error {
	// Implement image display logic here
	// You may need to use a library like github.com/eliukblau/pixterm
	// to display images in the terminal
	return fmt.Errorf("image display not implemented yet")
}

func ShowSystemInfo(info system.SystemInfo) {
	fmt.Println("System Information:")
	fmt.Printf("OS: %s\n", info.OS)
	fmt.Printf("Distribution: %s %s\n", info.Distribution, info.Version)
	fmt.Printf("Hostname: %s\n", info.Hostname)
	fmt.Printf("Kernel: %s\n", info.Kernel)
	fmt.Printf("Uptime: %dH\n", info.Uptime/3600)
	fmt.Printf("Architecture: %s\n", info.Architecture)
	fmt.Printf("Total memory: %dGB\n", info.TotalMemory/1024/1024/1024+1)
	fmt.Printf("Used memory percentage: %.0f%%\n", info.UsedMemoryPct)
}

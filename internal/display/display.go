package display

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/Cylis-Dragneel/phetch/internal/config"
	"github.com/Cylis-Dragneel/phetch/internal/system"
	"github.com/eliukblau/pixterm/pkg/ansimage"
	"golang.org/x/term"
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
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	width, height, err := getTerminalSize()
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %v", err)
	}

	// Adjust the image size to fit the terminal
	scaledHeight := height / 2 // Using half of the terminal height for the image
	scaledWidth := width / 2   // Using half of the terminal width for the image

	ansImg, err := ansimage.NewScaledFromImage(
		img,
		scaledHeight,
		scaledWidth,
		color.Black,
		ansimage.ScaleModeFit,
		ansimage.NoDithering,
	)
	if err != nil {
		return fmt.Errorf("failed to create ANSImage: %v", err)
	}

	ansImg.Draw()
	return nil
}

func getTerminalSize() (width, height int, err error) {
	fd := int(os.Stdout.Fd())
	if term.IsTerminal(fd) {
		width, height, err = term.GetSize(fd)
		if err != nil {
			return 0, 0, err
		}
	} else {
		// Fallback to a default size if not a terminal
		width, height = 80, 24
	}
	return width, height, nil
}

func ShowSystemInfo(info system.SystemInfo, cfg *config.Config) {
	fmt.Println("System Information:")
	if cfg.ShowOS {
		fmt.Printf("OS: %s\n", info.OS)
	}
	if cfg.ShowDistribution {
		fmt.Printf("Distribution: %s %s\n", info.Distribution, info.Version)
	}
	if cfg.ShowHostname {
		fmt.Printf("Hostname: %s\n", info.Hostname)
	}
	if cfg.ShowKernel {
		fmt.Printf("Kernel: %s\n", info.Kernel)
	}
	if cfg.ShowUptime {
		fmt.Printf("Uptime: %dH\n", info.Uptime/3600)
	}
	if cfg.ShowArchitecture {
		fmt.Printf("Architecture: %s\n", info.Architecture)
	}
	if cfg.ShowMemory {
		fmt.Printf("Total memory: %dGB\n", info.TotalMemory/1024/1024/1024+1)
		fmt.Printf("Used memory percentage: %.0f%%\n", info.UsedMemoryPct)
	}
}

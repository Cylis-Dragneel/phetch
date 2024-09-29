package main

import (
	"fmt"
	"os"

	"github.com/Cylis-Dragneel/phetch/internal/config"
	"github.com/Cylis-Dragneel/phetch/internal/display"
	"github.com/Cylis-Dragneel/phetch/internal/system"
)

func main() {
	// Load or create config file
	cfg, err := config.LoadOrCreate()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Get system information
	sysInfo := system.GetSystemInfo()

	// Override system info with hardcoded values from config
	sysInfo = config.OverrideSystemInfo(cfg, sysInfo)

	// Display ASCII art or image
	err = display.ShowArt(cfg.ArtPath)
	if err != nil {
		fmt.Printf("Error displaying art: %v\n", err)
	}

	// Display system information
	display.ShowSystemInfo(sysInfo, cfg)
}

package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cylis-Dragneel/phetch/internal/system"
	lua "github.com/yuin/gopher-lua"
)

type Config struct {
	ArtPath string
	// Add other config fields as needed
	ShowOS           bool
	ShowDistribution bool
	ShowHostname     bool
	ShowKernel       bool
	ShowUptime       bool
	ShowArchitecture bool
	ShowMemory       bool
	UseKittyProtocol bool // Add this new field
}

func LoadOrCreate() (*Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "phetch", "config.lua")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return createDefaultConfig(configPath)
	}

	return loadConfig(configPath)
}

func createDefaultConfig(path string) (*Config, error) {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create config directory: %v", err)
	}

	defaultConfig := `
-- Phetch Configuration File

-- Path to ASCII art or image file
art_path = ""

-- System information display options
show_os = true
show_distribution = true
show_hostname = true
show_kernel = true
show_uptime = true
show_architecture = true
show_memory = true
use_kitty_protocol = true

-- Add other configuration options here
`

	err = os.WriteFile(path, []byte(defaultConfig), 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write default config: %v", err)
	}

	fmt.Printf("Created default configuration file at %s\n", path)
	return loadConfig(path)
}

func loadConfig(path string) (*Config, error) {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(path); err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}

	cfg := &Config{
		ShowOS:           true,
		ShowDistribution: true,
		ShowHostname:     true,
		ShowKernel:       true,
		ShowUptime:       true,
		ShowArchitecture: true,
		ShowMemory:       true,
	}

	if lv := L.GetGlobal("art_path"); lv.Type() == lua.LTString {
		cfg.ArtPath = L.ToString(-1)
	}

	getBoolOption := func(key string) bool {
		if lv := L.GetGlobal(key); lv.Type() == lua.LTBool {
			return lua.LVAsBool(lv)
		}
		return true // Default to true if not specified
	}

	cfg.ShowOS = getBoolOption("show_os")
	cfg.ShowDistribution = getBoolOption("show_distribution")
	cfg.ShowHostname = getBoolOption("show_hostname")
	cfg.ShowKernel = getBoolOption("show_kernel")
	cfg.ShowUptime = getBoolOption("show_uptime")
	cfg.ShowArchitecture = getBoolOption("show_architecture")
	cfg.ShowMemory = getBoolOption("show_memory")
	cfg.UseKittyProtocol = getBoolOption("use_kitty_protocol")

	return cfg, nil
}

func OverrideSystemInfo(cfg *Config, sysInfo system.SystemInfo) system.SystemInfo {
	// This function remains unchanged for now
	// You can implement logic to override system info with hardcoded values from config if needed
	return sysInfo
}

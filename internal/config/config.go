package config

import (
	"os"
	"path/filepath"

	"github.com/Cylis-Dragneel/phetch/internal/system"
)

type Config struct {
	ArtPath string
	// Add other config fields as needed
}

func LoadOrCreate() (*Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "phetch", "config.lua")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return createDefaultConfig(configPath)
	}

	return loadConfig(configPath)
}

func createDefaultConfig(path string) (*Config, error) {
	// Create default config file
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, err
	}

	defaultConfig := `
artPath = ""
-- Add other default config values here
`

	err = os.WriteFile(path, []byte(defaultConfig), 0644)
	if err != nil {
		return nil, err
	}

	return &Config{ArtPath: ""}, nil
}

func loadConfig(path string) (*Config, error) {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(path); err != nil {
		return nil, err
	}

	cfg := &Config{}

	if lv := L.GetGlobal("artPath"); lv.Type() == lua.LTString {
		cfg.ArtPath = L.ToString(-1)
	}

	// Load other config values here

	return cfg, nil
}

func OverrideSystemInfo(cfg *Config, sysInfo system.SystemInfo) system.SystemInfo {
	// Implement logic to override system info with hardcoded values from config
	return sysInfo
}

package device

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/anothermeer/sambal/internal/core/protocol"
)

func GetID() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return uuid.NewString()
	}

	sambalDir := filepath.Join(configDir, "sambal")
	configFile := filepath.Join(sambalDir, "device.json")

	// already exists
	if data, err := os.ReadFile(configFile); err == nil {
		var cfg protocol.Config
		if json.Unmarshal(data, &cfg) == nil && cfg.ID != "" {
			return cfg.ID
		}
	}

	// create new
	os.MkdirAll(sambalDir, 0755)

	cfg := protocol.Config{
		ID: uuid.NewString(),
	}

	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(configFile, data, 0644)

	return cfg.ID
}

func GetName() string {
	name, err := os.Hostname()
	if err != nil {
		return "Unknown Device"
	}
	return name
}

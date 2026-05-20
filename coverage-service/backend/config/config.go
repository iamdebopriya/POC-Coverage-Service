package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type ServiceConfig struct {
	Name         string `json:"name"`
	DisplayName  string `json:"display_name"`
	BackendPath  string `json:"backend_path"`
	FrontendPath string `json:"frontend_path"`
	BackendType  string `json:"backend_type"`
	FrontendType string `json:"frontend_type"`
}

func LoadServices() ([]ServiceConfig, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "services.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var services []ServiceConfig
	if err := json.Unmarshal(data, &services); err != nil {
		return nil, err
	}
	return services, nil
}

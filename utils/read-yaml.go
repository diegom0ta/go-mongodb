// utils/fileutil.go

package utils

import (
	"os"
	"path/filepath"
)

func ReadConfigFile() ([]byte, error) {

	// Construct the path to config.yaml (assuming it's in the root directory)
	configPath := filepath.Join("config", "config.yml")

	// Read the contents of config.yaml
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

package utils

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"db"`
	Secret struct {
		SecretKey string `yaml:"secretKey"`
	} `yaml:"secret"`
}

// Parse YAML data into config struct
func ParseYaml() (Config, error) {
	var config Config

	data, err := ReadConfigFile()
	if err != nil {
		return config, fmt.Errorf("error reading YAML: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return config, nil
}

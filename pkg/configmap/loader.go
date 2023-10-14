package configmap

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ReadConfigFile(filePath string) (Config, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

package microservice

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	CONFIG_RUN_ADDRESS = "RunAddress"
	CONFIG_GATEWAY     = "Gateway"
)

// Reads ServiceConfig.json from the microservice directory
func (service *wdtkService) readConfig() (map[string]interface{}, error) {
	configContent, err := os.ReadFile("ServiceConfig.json")
	if err != nil {
		return nil, err
	}

	var config map[string]interface{}
	err = json.Unmarshal(configContent, &config)

	if err != nil {
		return nil, err
	}

	// Verify
	if _, exists := config[CONFIG_RUN_ADDRESS]; !exists {
		return nil, errors.New("key " + CONFIG_RUN_ADDRESS + " doesn't exist in the config file")
	}

	if _, exists := config[CONFIG_GATEWAY]; !exists {
		return nil, errors.New("key  " + CONFIG_GATEWAY + " doesn't exist in the config file")
	}

	return config, nil
}

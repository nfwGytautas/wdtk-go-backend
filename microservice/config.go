package microservice

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	CONFIG_RUN_ADDRESS = "runAddress"
	CONFIG_GATEWAY     = "gatewayIp"
	CONFIG_API_KEY     = "apiKey"
)

// Config struct that is read from the ServiceConfig.json
type MicroserviceConfig struct {
	RunAddress  string                 `json:"runAddress"`
	GatewayIp   string                 `json:"gatewayIp"`
	ApiKey      string                 `json:"apiKey"`
	UserDefines map[string]interface{} `json:"-"`
}

func ReadConfig() (*MicroserviceConfig, error) {
	result := MicroserviceConfig{}

	configContent, err := os.ReadFile("ServiceConfig.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configContent, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

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

	if _, exists := config[CONFIG_API_KEY]; !exists {
		return nil, errors.New("key  " + CONFIG_API_KEY + " doesn't exist in the config file")
	}

	return config, nil
}

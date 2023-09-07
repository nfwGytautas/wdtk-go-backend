package microservice

import (
	"encoding/json"
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
	result.UserDefines = make(map[string]interface{})

	configContent, err := os.ReadFile("WdtkConfig.json")
	if err != nil {
		return nil, err
	}

	var entries map[string]*json.RawMessage
	err = json.Unmarshal(configContent, &entries)
	if err != nil {
		return nil, err
	}

	for key, element := range entries {
		var object interface{}
		err = json.Unmarshal(*element, &object)
		if err != nil {
			return nil, err
		}

		if key == CONFIG_RUN_ADDRESS {
			result.RunAddress = object.(string)
		} else if key == CONFIG_GATEWAY {
			result.GatewayIp = object.(string)
		} else if key == CONFIG_API_KEY {
			result.ApiKey = object.(string)
		} else {
			result.UserDefines[key] = object
		}
	}

	return &result, nil
}

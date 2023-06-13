package microservice

import (
	"encoding/json"
	"os"
)

// The config for a service
type ServiceConfig struct {
	Gateway string `json:"Gateway"`
}

// Reads ServiceConfig.json from the microservice directory
func (service *wdtkService) readConfig(config *interface{}) error {
	configContent, err := os.ReadFile("ServiceConfig.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(configContent, &config)
	return err
}

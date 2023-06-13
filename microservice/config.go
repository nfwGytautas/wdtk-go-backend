package microservice

import (
	"encoding/json"
	"os"
)

// Reads ServiceConfig.json from the microservice directory
func (service *wdtkService) readConfig() (map[string]interface{}, error) {
	configContent, err := os.ReadFile("ServiceConfig.json")
	if err != nil {
		return nil, err
	}

	var config map[string]interface{}
	err = json.Unmarshal(configContent, &config)
	return config, err
}

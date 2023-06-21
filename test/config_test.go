package test

import (
	"fmt"
	"testing"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

func TestReadConfig(t *testing.T) {
	config, err := microservice.ReadConfig()
	if err != nil {
		t.Fatal(err)
	}

	if config.ApiKey != "TEST_SECRET" {
		fmt.Printf("API key expected: 'TEST_SECRET' got '%s'", config.ApiKey)
		t.Fail()
	}

	if config.RunAddress != ":8080" {
		fmt.Printf("RunAddress key expected: ':8080' got '%s'", config.ApiKey)
		t.Fail()
	}

	if config.GatewayIp != "..." {
		fmt.Printf("GatewayIp key expected: '...' got '%s'", config.ApiKey)
		t.Fail()
	}

	if config.UserDefines["userValue"].(string) != "VALUE" {
		fmt.Printf("UserValue key expected: 'VALUE' got '%s'", config.ApiKey)
		t.Fail()
	}
}

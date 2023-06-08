package test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

type ServiceData struct {
}

func exampleEndpoint(executor *microservice.EndpointExecutor) {
	log.Println("Running example endpoint")
	log.Println(executor.Params)

	executor.Return(http.StatusOK, nil)
}

func createAndRunService() {
	if microservice.RegisterService(microservice.ServiceDescription{
		ServiceContext: &ServiceData{},
	}, []microservice.ServiceEndpoint{
		{
			Type:            microservice.ENDPOINT_TYPE_GET,
			Name:            "TestEndpoint/:id",
			Fn:              exampleEndpoint,
			EndpointContext: nil,
		},
	}) != nil {
		panic("Failed to run")
	}
}

func TestSimpleService(t *testing.T) {
	go createAndRunService()

	time.Sleep(time.Second * 3)

	resp, err := http.Get("http://localhost:8080/TestEndpoint/3")
	if err != nil {
		t.Error(err)
		return
	}

	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))

	if resp.StatusCode != http.StatusOK {
		t.Fail()
		return
	}
}

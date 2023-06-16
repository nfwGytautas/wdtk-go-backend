package test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/nfwGytautas/gdev/jwt"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

type ServiceData struct {
}

type EndpointContext struct {
	Int int
}

func noAuthEndpoint(executor *microservice.EndpointExecutor) {
	log.Println("Running noAuth endpoint")
	log.Println(executor.Params)
	executor.EndpointContext.(*EndpointContext).Int = 5
	executor.Return(http.StatusOK, nil)
}

func authEndpoint(executor *microservice.EndpointExecutor) {
	log.Println("Running auth endpoint")
	log.Println(executor.Params)
	log.Println(executor.RequesterInfo)
	executor.Return(http.StatusOK, nil)
}

var eContext EndpointContext

func createAndRunService() {
	if err := microservice.RegisterService[ServiceData](microservice.ServiceDescription[ServiceData]{}, []microservice.ServiceEndpoint{
		{
			Type:            microservice.ENDPOINT_TYPE_GET,
			Name:            "noAuth/",
			Fn:              noAuthEndpoint,
			EndpointContext: &eContext,
			AuthRequired:    false,
		},
		{
			Type:            microservice.ENDPOINT_TYPE_GET,
			Name:            "auth/",
			Fn:              authEndpoint,
			EndpointContext: &eContext,
			AuthRequired:    true,
		},
	}); err != nil {
		log.Println(err)
		panic("Failed to register service")
	}
}

func TestNoAuth(t *testing.T) {
	go createAndRunService()

	time.Sleep(time.Second * 3)

	resp, err := http.Get("http://localhost:8080/noAuth")
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

	if eContext.Int != 5 {
		fmt.Printf("Endpoint context value %v expected 5", eContext.Int)
		t.Fail()
	}
}

func TestAuth(t *testing.T) {
	go createAndRunService()
	time.Sleep(time.Second * 3)

	tokenString, err := jwt.GenerateToken(123, "Role")
	if err != nil {
		t.Error(err)
		return
	}

	request, err := http.NewRequest("GET", "http://localhost:8080/auth/", nil)
	if err != nil {
		t.Error(err)
		return
	}

	request.Header.Add("Authorization", "bearer "+tokenString)

	resp, err := http.DefaultClient.Do(request)
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

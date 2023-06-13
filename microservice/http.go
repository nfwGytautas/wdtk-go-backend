package microservice

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Run HTTP communication service (uses gin)
func (service *wdtkService) runHTTP() error {
	r := gin.Default()
	gs := r.Group("/")

	for _, endpoint := range service.endpoints {
		gs.Handle(endpoint.Type, endpoint.Name, service.createEndpointHandler(&endpoint))
	}

	return r.Run(":8080")
}

func (service *wdtkService) createEndpointHandler(sp *ServiceEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse request body"})
			return
		}

		executor := EndpointExecutor{
			ServiceContext:  service.context,
			EndpointContext: sp.EndpointContext,
			Body:            body,
			Params:          map[string]string{},
		}

		for _, p := range c.Params {
			executor.Params[p.Key] = p.Value
		}

		sp.Fn(&executor)

		if executor.statusCode == 0 {
			panic("Endpoint returned a invalid status code")
		}

		if executor.body == nil {
			c.JSON(executor.statusCode, gin.H{})
			return
		}
		c.JSON(executor.statusCode, executor.body)
	}
}

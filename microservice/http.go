package microservice

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/gdev/jwt"
)

// Run HTTP communication service (uses gin)
func (service *wdtkService) runHTTP() error {
	r := gin.Default()
	gs := r.Group("/")

	if len(service.config[CONFIG_API_KEY].(string)) == 0 {
		return errors.New("api key is empty")
	}

	jwt.APISecret = service.config[CONFIG_API_KEY].(string)

	for _, endpoint := range service.endpoints {
		var handlers []gin.HandlerFunc

		if len(endpoint.Roles) > 0 {
			// Requires role authorization
			handlers = append(handlers, jwt.AuthorizationMiddleware(endpoint.Roles))
		} else if endpoint.AuthRequired {
			handlers = append(handlers, jwt.AuthenticationMiddleware())
		}

		handlers = append(handlers, service.createEndpointHandler(&endpoint))
		gs.Handle(endpoint.Type, endpoint.Name, handlers...)
	}

	return r.Run(service.config[CONFIG_RUN_ADDRESS].(string))
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

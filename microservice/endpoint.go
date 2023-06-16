package microservice

import "github.com/nfwGytautas/gdev/jwt"

// Endpoint type enum
const (
	ENDPOINT_TYPE_GET    = "GET"
	ENDPOINT_TYPE_POST   = "POST"
	ENDPOINT_TYPE_EDIT   = "PUT"
	ENDPOINT_TYPE_DELETE = "DELETE"
)

// Endpoint function type
type EndpointFn func(*EndpointExecutor)

/*
 * A single endpoint
 */
type ServiceEndpoint struct {
	Type            string
	Name            string
	EndpointContext interface{}
	Fn              EndpointFn
	AuthRequired    bool
	Roles           []string
}

/*
 * Executor for endpoints
 */
type EndpointExecutor struct {
	// The service context
	ServiceContext interface{}

	// Endpoint context
	EndpointContext interface{}

	// Body for endpoint to execute
	Body []byte

	// Request parameters
	Params map[string]string

	// The token info for the requester, nil for endpoints that have authorization disabled
	RequesterInfo *jwt.TokenInfo

	statusCode int
	body       interface{}
}

// Finish endpoint execution
func (ee *EndpointExecutor) Return(statusCode int, body interface{}) {
	ee.statusCode = statusCode
	ee.body = body
}

package microservice

/*
 * A struct for describing endpoint errors
 */
type EndpointError struct {
	Description string `json:"description"`
	Error       error  `json:"error"`
}

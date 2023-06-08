package microservice

/*
 * Description of the service
 */
type ServiceDescription struct {
	// The service context
	ServiceContext interface{}
}

/*
 * The returned object from RegisterService method, this is basically the WDTK controller for a service
 */
type wdtkService struct {
	endpoints []ServiceEndpoint
	desc      ServiceDescription
}

// Register any interface type as service
func RegisterService(description ServiceDescription, endpoints []ServiceEndpoint) error {
	service := wdtkService{}

	service.desc = description
	service.endpoints = endpoints

	return service.runHTTP()
}

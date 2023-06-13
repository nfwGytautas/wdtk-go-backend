package microservice

// Setup function type, first argument is the service context, the second argument is the config
type SetupFn func(*interface{}, interface{}) error

/*
 * Description of the service
 */
type ServiceDescription struct {
	// The service context
	ServiceContext interface{}

	// The callback for setup if needed
	SetupFn SetupFn

	// The config structure, must inherit ServiceConfig struct
	Config interface{}
}

/*
 * The returned object from RegisterService method, this is basically the WDTK controller for a service
 */
type wdtkService struct {
	endpoints []ServiceEndpoint
	context   interface{}
}

// Register any interface type as service
func RegisterService(description ServiceDescription, endpoints []ServiceEndpoint) error {
	service := wdtkService{}

	service.context = description.ServiceContext
	service.endpoints = endpoints

	// Read config
	err := service.readConfig(&description.Config)
	if err != nil {
		return err
	}

	if description.SetupFn != nil {
		err := description.SetupFn(&service.context, description.Config)
		if err != nil {
			return err
		}
	}

	return service.runHTTP()
}

package microservice

// Setup function type, first argument is the service context, the second argument is the config
type SetupFn[Context any] func(*Context, map[string]interface{}) error

/*
 * Description of the service
 */
type ServiceDescription[Context any] struct {
	// The callback for setup if needed
	SetupFn SetupFn[Context]
}

/*
 * The returned object from RegisterService method, this is basically the WDTK controller for a service
 */
type wdtkService struct {
	endpoints []ServiceEndpoint
	context   any
	config    map[string]interface{}
}

// Register any interface type as service
func RegisterService[Context any](description ServiceDescription[Context], endpoints []ServiceEndpoint) error {
	service := wdtkService{}

	var context Context
	service.context = &context
	service.endpoints = endpoints

	// Read config
	config, err := service.readConfig()
	if err != nil {
		return err
	}

	if description.SetupFn != nil {
		err := description.SetupFn(service.context.(*Context), config)
		if err != nil {
			return err
		}
	}

	service.config = config
	return service.runHTTP()
}

package microservice

import "errors"

const (
	COMM_TYPE_HTTP int = 0
)

/*
 * The returned object from RegisterService method, this is basically the WDTK controller for a service
 */
type WDTKService struct {
	CommunicationType int
	impl              interface{}
}

// Register any interface type as service
func RegisterService(implementation interface{}) (*WDTKService, error) {
	service := WDTKService{}

	if implementation == nil {
		return nil, errors.New("implementation cannot be nil")
	}
	service.impl = implementation

	return &service, nil
}

// Run the service
func (service *WDTKService) Run() error {
	return nil
}

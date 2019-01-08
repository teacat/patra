package balancer

import (
	"github.com/hashicorp/consul/api"
	"github.com/teacat/patra/model"
)

// NewConsul
func NewConsul(conf *api.Config) (Balancer, error) {
	client, err := api.NewClient(conf)
	if err != nil {
		return err
	}
	return &consulBalancer{
		client: client,
	}, nil
}

type consulBalancer struct {
	client *api.Client
}

// RegisterService
func (c *consulBalancer) RegisterService(req *model.RegisterServiceRequest) error {
	return nil
}

// GetService
func (c *consulBalancer) GetService(typ BalanceType, req *model.GetServiceRequest) (*model.ServiceInformation, error) {

}

// ListServices
func (c *consulBalancer) ListServices(req *model.ListServicesRequest) ([]*model.ServiceInformation, error) {
	return nil, nil
}

// DeregisterService
func (c *consulBalancer) DeregisterService(req *model.DeregisterServiceRequest) error {
	return nil
}

// CheckServices
func (c *consulBalancer) CheckServices() error {
	return nil
}

// PutValue
func (c *consulBalancer) PutValue(req *model.PutValueRequest) error {
	return nil
}

// GetValue
func (c *consulBalancer) GetValue(req *model.GetValueRequest) (string, error) {

}

// ListKeys
func (c *consulBalancer) ListKeys(req *model.ListKeysRequest) ([]string, error) {

}

// ListValues
func (c *consulBalancer) ListValues(req *model.ListValuesRequest) (map[string]string, error) {

}

// DeleteValue
func (c *consulBalancer) DeleteValue(req *model.DeleteValueRequest) error {
	return nil
}

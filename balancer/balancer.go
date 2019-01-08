package balancer

import (
	"github.com/teacat/patra/model"
)

type BalanceType int

const (
	TypeRoundRobin BalanceType = iota
	TypeLeastConnection
	TypeLeastUsage
	TypeRandom
)

// Balancer
type Balancer interface {
	// RegisterService
	RegisterService(*model.RegisterServiceRequest) error
	// GetService
	GetService(BalanceType, *model.GetServiceRequest) (*model.ServiceInformation, error)
	// ListServices
	ListServices(*model.ListServicesRequest) ([]*model.ServiceInformation, error)
	// DeregisterService
	DeregisterService(*model.DeregisterServiceRequest) error
	// CheckServices
	CheckServices() error

	// PutValue
	PutValue(*model.PutValueRequest) error
	// GetValue
	GetValue(*model.GetValueRequest) (string, error)
	// ListKeys
	ListKeys(*model.ListKeysRequest) ([]string, error)
	// ListValues
	ListValues(*model.ListValuesRequest) (map[string]string, error)
	// DeleteValue
	DeleteValue(*model.DeleteValueRequest) error
}

// NewDefault
func NewDefault() (Balancer, error) {
	return &defaultBalancer{}, nil
}

// service
type service struct {
	// id
	id string
	// name
	name string
	// port
	port int
	// address
	address string
	// tags
	tags []string
	// metadata
	metadata string
	// version
	version string

	// cpuUsage
	cpuUsage int
	// ramUsage
	ramUsage int
	// diskUsage
	diskUsage int
	// isServing
	isServing bool
}

// defaultBalancer
type defaultBalancer struct {
	kvstore  map[string]string
	services map[string][]*service

	lastRoundRobinSvc *service
}

// RegisterService
func (d *defaultBalancer) RegisterService(req *model.RegisterServiceRequest) error {
	return nil
}

// GetService
func (d *defaultBalancer) GetService(typ BalanceType, req *model.GetServiceRequest) (*model.ServiceInformation, error) {
	switch typ {
	case TypeLeastConnection:
		break
	case TypeLeastUsage:
		break
	case TypeRandom:
		break
	default:
		break
	}
}

// ListServices
func (d *defaultBalancer) ListServices(req *model.ListServicesRequest) ([]*model.ServiceInformation, error) {
	var services []*model.ServiceInformation

	for k, v := range d.services {
		//
		if req.Name != "" && req.Name != k {
			continue
		}
		for _, v := range v {
			//
			if req.Version != "" && req.Version != v.version {
				continue
			}
			//
			if req.Tag != "" {
				var has bool
				for _, v := range v.tags {
					if req.Tag == v {
						has = true
						break
					}
				}
				if !has {
					continue
				}
			}
			services = append(services, &model.ServiceInformation{
				Id:       v.id,
				Name:     v.name,
				Port:     uint32(v.port),
				Address:  v.address,
				Tags:     v.tags,
				Metadata: v.metadata,
				Version:  v.version,
			})
		}
	}
	return services, nil
}

// DeregisterService
func (d *defaultBalancer) DeregisterService(req *model.DeregisterServiceRequest) error {
	for k, svcs := range d.services {
		for i, v := range svcs {
			if v.id == req.Id {
				d.services[k] = append(svcs[:i], svcs[i+1:]...)
			}
		}
	}
	return nil
}

// CheckServices
func (d *defaultBalancer) CheckServices() error {
	return nil
}

// PutValue
func (d *defaultBalancer) PutValue(req *model.PutValueRequest) error {
	return nil
}

// GetValue
func (d *defaultBalancer) GetValue(req *model.GetValueRequest) (string, error) {

}

// ListKeys
func (d *defaultBalancer) ListKeys(req *model.ListKeysRequest) ([]string, error) {

}

// ListValues
func (d *defaultBalancer) ListValues(req *model.ListValuesRequest) (map[string]string, error) {

}

// DeleteValue
func (d *defaultBalancer) DeleteValue(req *model.DeleteValueRequest) error {
	return nil
}

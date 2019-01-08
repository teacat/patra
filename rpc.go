package patra

import (
	"context"

	"github.com/teacat/patra/balancer"
	"github.com/teacat/patra/model"
)

type server struct {
	balancer balancer.Balancer
}

func (s *server) RegisterService(ctx context.Context, in *model.RegisterServiceRequest) (*model.RegisterServiceResponse, error) {

}

func (s *server) GetService(ctx context.Context, in *model.GetServiceRequest) (*model.GetServiceResponse, error) {

}

func (s *server) ListServices(ctx context.Context, in *model.ListServicesRequest) (*model.ListServicesResponse, error) {

}

func (s *server) DeregisterService(ctx context.Context, in *model.DeregisterServiceRequest) (*model.DeregisterServiceResponse, error) {

}

func (s *server) PutValue(ctx context.Context, in *model.PutValueRequest) (*model.PutValueResponse, error) {

}

func (s *server) GetValue(ctx context.Context, in *model.GetValueRequest) (*model.GetValueResponse, error) {

}

func (s *server) ListKeys(ctx context.Context, in *model.ListKeysRequest) (*model.ListKeysResponse, error) {

}

func (s *server) ListValues(ctx context.Context, in *model.ListValuesRequest) (*model.ListValuesResponse, error) {

}

func (s *server) DeleteValue(ctx context.Context, in *model.DeleteValueRequest) (*model.DeleteValueResponse, error) {

}

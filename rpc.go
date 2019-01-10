package patra

import (
	"context"

	"github.com/teacat/patra/balancer"
	"github.com/teacat/patra/model"
)

type server struct {
	balancer balancer.Balancer
}

func (s *server) GetService(ctx context.Context, in *model.GetServiceRequest) (*model.GetServiceResponse, error) {

}

func (s *server) ListServices(ctx context.Context, in *model.ListServicesRequest) (*model.ListServicesResponse, error) {

}

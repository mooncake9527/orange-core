package health

import (
	"context"
)

type HealthServerImpl struct {
	*UnimplementedHealthServer
}

func (s *HealthServerImpl) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return &HealthCheckResponse{Status: HealthCheckResponse_SERVING}, nil
}

func (s *HealthServerImpl) Watch(*HealthCheckRequest, Health_WatchServer) error {
	return nil
}

//func (HealthServerImpl) mustEmbedUnimplementedHealthServer() {}

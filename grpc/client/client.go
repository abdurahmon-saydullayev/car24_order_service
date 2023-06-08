package client

import (
	"Projects/Car24/car24_order_service/config"
	"Projects/Car24/car24_order_service/genproto/client_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	UserService() client_service.ClientServiceClient
}

type grpcClients struct {
	userService client_service.ClientServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService: client_service.NewClientServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) UserService() client_service.ClientServiceClient {
	return g.userService
}

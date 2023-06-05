package client

import "Projects/Car24/car24_order_service/config"

type ServiceManagerI interface{}

type grpcClients struct{}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	return &grpcClients{}, nil
}

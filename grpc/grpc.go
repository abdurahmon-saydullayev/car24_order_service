package grpc

import (
	"Projects/Car24/car24_order_service/config"
	"Projects/Car24/car24_order_service/genproto/order_service"
	"Projects/Car24/car24_order_service/grpc/client"
	"Projects/Car24/car24_order_service/grpc/service"
	"Projects/Car24/car24_order_service/pkg/logger"
	"Projects/Car24/car24_order_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	order_service.RegisterOrderServiceServer(grpcServer, service.NewOrderService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}

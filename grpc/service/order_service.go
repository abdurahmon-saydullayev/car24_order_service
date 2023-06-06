package service

import (
	"Projects/Car24/car24_order_service/config"
	"Projects/Car24/car24_order_service/genproto/order_service"
	"Projects/Car24/car24_order_service/grpc/client"
	"Projects/Car24/car24_order_service/pkg/logger"
	"Projects/Car24/car24_order_service/storage"
)

type OrderService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*order_service.UnimplementedOrderServiceServer
}

func NewOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *OrderService {
	return &OrderService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

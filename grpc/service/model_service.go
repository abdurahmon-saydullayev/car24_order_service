package service

import (
	"Projects/Car24/car24_order_service/config"
	"Projects/Car24/car24_order_service/genproto/order_service"
	"Projects/Car24/car24_order_service/grpc/client"
	"Projects/Car24/car24_order_service/pkg/logger"
	"Projects/Car24/car24_order_service/storage"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ModelService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*order_service.UnimplementedModelServiceServer
}

func NewModelService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ModelService {
	return &ModelService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ModelService) Create(ctx context.Context, req *order_service.CreateModel) (resp *order_service.Model, err error) {

	i.log.Info("---CreateOrder------>", logger.Any("req", req))

	pKey, err := i.strg.Model().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateOrder->Order->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Model().GetByID(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ModelService) GetByID(ctx context.Context, req *order_service.ModelPK) (resp *order_service.Model, err error) {

	i.log.Info("---GetOrderByID------>", logger.Any("req", req))

	resp, err = i.strg.Model().GetByID(ctx, req)
	if err != nil {
		i.log.Error("!!!GetOrderByID->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

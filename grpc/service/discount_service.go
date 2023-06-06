package service

import (
	"Projects/Car24/car24_order_service/config"
	"Projects/Car24/car24_order_service/genproto/order_service"
	"Projects/Car24/car24_order_service/grpc/client"
	"Projects/Car24/car24_order_service/pkg/logger"
	"Projects/Car24/car24_order_service/storage"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DiscountService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*order_service.UnimplementedDiscountServiceServer
}

func NewDiscountService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *DiscountService {
	return &DiscountService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *DiscountService) Create(ctx context.Context, req *order_service.CreateDiscount) (resp *order_service.Discount, err error) {

	i.log.Info("---CreateOrder------>", logger.Any("req", req))

	pKey, err := i.strg.Discount().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateOrder->Order->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Discount().GetByID(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *DiscountService) GetByID(ctx context.Context, req *order_service.DiscountPK) (resp *order_service.Discount, err error) {

	i.log.Info("---GetOrderByID------>", logger.Any("req", req))

	resp, err = i.strg.Discount().GetByID(ctx, req)
	if err != nil {
		i.log.Error("!!!GetOrderByID->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *DiscountService) Delete(ctx context.Context, req *order_service.DiscountPK) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteOrder------>", logger.Any("req", req))

	err = i.strg.Discount().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}

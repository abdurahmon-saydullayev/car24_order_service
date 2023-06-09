package service

import (
	"Projects/Car24/car24_order_service/config"
	"Projects/Car24/car24_order_service/genproto/order_service"
	"Projects/Car24/car24_order_service/grpc/client"
	"Projects/Car24/car24_order_service/models"
	"Projects/Car24/car24_order_service/pkg/logger"
	"Projects/Car24/car24_order_service/storage"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (i *OrderService) Create(ctx context.Context, req *order_service.CreateOrder) (resp *order_service.Order, err error) {

	i.log.Info("---CreateOrder------>", logger.Any("req", req))

	pKey, err := i.strg.Order().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateOrder->Order->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Order().GetByID(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *OrderService) GetByID(ctx context.Context, req *order_service.OrderPrimaryKey) (resp *order_service.Order, err error) {

	i.log.Info("---GetOrderByID------>", logger.Any("req", req))

	resp, err = i.strg.Order().GetByID(ctx, req)
	if err != nil {
		i.log.Error("!!!GetOrderByID->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *OrderService) GetList(ctx context.Context, req *order_service.GetListOrderRequest) (resp *order_service.GetListOrderResponse, err error) {

	i.log.Info("---GetOrders------>", logger.Any("req", req))

	resp, err = i.strg.Order().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetOrders->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *OrderService) Update(ctx context.Context, req *order_service.UpdateOrder) (resp *order_service.Order, err error) {

	i.log.Info("---UpdateOrder------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Order().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Order().GetByID(ctx, &order_service.OrderPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *OrderService) UpdatePatch(ctx context.Context, req *order_service.UpdatePatchOrder) (resp *order_service.Order, err error) {

	i.log.Info("---UpdatePatchOrder------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Order().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Order().GetByID(ctx, &order_service.OrderPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetOrder->Order->Get--->", logger.Error(err))

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *OrderService) Delete(ctx context.Context, req *order_service.OrderPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteOrder------>", logger.Any("req", req))

	err = i.strg.Order().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}

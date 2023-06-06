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

type CarService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*order_service.UnimplementedCarServiceServer
}

func NewCarService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CarService {
	return &CarService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CarService) Create(ctx context.Context, req *order_service.CreateCar) (resp *order_service.Car, err error) {

	i.log.Info("---CreateOrder------>", logger.Any("req", req))

	pKey, err := i.strg.Car().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateOrder->Order->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Car().GetByID(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CarService) GetByID(ctx context.Context, req *order_service.CarPrimaryKey) (resp *order_service.Car, err error) {

	i.log.Info("---GetOrderByID------>", logger.Any("req", req))

	resp, err = i.strg.Car().GetByID(ctx, req)
	if err != nil {
		i.log.Error("!!!GetOrderByID->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CarService) GetList(ctx context.Context, req *order_service.GetListCarRequest) (resp *order_service.GetListCarResponse, err error) {

	i.log.Info("---GetOrders------>", logger.Any("req", req))

	resp, err = i.strg.Car().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetOrders->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CarService) Update(ctx context.Context, req *order_service.UpdateCar) (resp *order_service.Car, err error) {

	i.log.Info("---UpdateOrder------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Car().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Car().GetByID(ctx, &order_service.CarPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CarService) UpdatePatch(ctx context.Context, req *order_service.UpdatePathCar) (resp *order_service.Car, err error) {

	i.log.Info("---UpdatePatchOrder------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Car().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Car().GetByID(ctx, &order_service.CarPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetOrder->Order->Get--->", logger.Error(err))

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CarService) Delete(ctx context.Context, req *order_service.CarPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteOrder------>", logger.Any("req", req))

	err = i.strg.Car().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteOrder->Order->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}

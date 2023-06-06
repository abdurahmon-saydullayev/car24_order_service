package storage

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"Projects/Car24/car24_order_service/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Discount() DiscountRepoI
	Mechanic() MechnicRepoI
	Model() ModelRepoI
	Tarif() TarifRepoI
	Car() CarRepoI
	Order() OrderRepoI
}

type DiscountRepoI interface {
	Create(context.Context, *order_service.CreateDiscount) (*order_service.DiscountPK, error)
	GetByID(context.Context, *order_service.DiscountPK) (*order_service.Discount, error)
	Delete(context.Context, *order_service.DiscountPK) error
}

type MechnicRepoI interface {
	Create(context.Context, *order_service.CreateMechanic) (*order_service.MechanicPK, error)
	GetByID(context.Context, *order_service.MechanicPK) (*order_service.Mechanic, error)
}

type ModelRepoI interface {
	Create(context.Context, *order_service.CreateModel) (*order_service.ModelPK, error)
	GetByID(context.Context, *order_service.ModelPK) (*order_service.Model, error)
}

type TarifRepoI interface {
	Create(context.Context, *order_service.CreateTarif) (*order_service.TarifPK, error)
	GetByID(context.Context, *order_service.TarifPK) (*order_service.Tarif, error)
}

type CarRepoI interface {
	Create(context.Context, *order_service.CreateCar) (*order_service.Car, error)
	GetByID(context.Context, *order_service.CarPrimaryKey) (*order_service.Car, error)
	GetList(context.Context, *order_service.GetListCarRequest) (*order_service.GetListCarResponse, error)
	Update(context.Context, *order_service.UpdateCar) (int64, error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (resp int64, err error)
	Delete(context.Context, *order_service.CarPrimaryKey) error
}

type OrderRepoI interface {
	Create(context.Context, *order_service.CreateOrder) (*order_service.OrderPrimaryKey, error)
	GetByID(context.Context, *order_service.OrderPrimaryKey) (*order_service.Order, error)
	GetList(context.Context, *order_service.GetListOrderRequest) (*order_service.GetListOrderResponse, error)
	Update(context.Context, *order_service.UpdateOrder) (int64, error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (resp int64, err error)
	Delete(context.Context, *order_service.OrderPrimaryKey) error
}

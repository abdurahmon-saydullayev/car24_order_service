package storage

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"
)

type StorageI interface {
	CloseDB()
	Order() OrderRepoI
}

type OrderRepoI interface {
	Create(context.Context, *order_service.CreateOrder) (*order_service.OrderPrimaryKey, error)
	GetByID(context.Context, *order_service.OrderPrimaryKey) (*order_service.Order, error)
	GetList(context.Context, *order_service.GetListOrdertRequest) (*order_service.GetListOrdertResponse, error)
	Update(context.Context, *order_service.UpdateOrder) (int64, error)
	UpdatePatch(context.Context, *order_service.UpdatePatchOrder) (int64, error)
	Delete(context.Context, *order_service.OrderPrimaryKey) error
}

type MechnicServiceI interface {
	Create(context.Context, *order_service.CreateMechanic) (*order_service.MechanicPK, error)
	GetByID(context.Context, *order_service.MechanicPK) (*order_service.Mechanic,error)
}

type ModelService interface {
	Create(context.Context, *order_service.CreateModel) (*order_service.ModelPK, error)
	GetByID(context.Context, *order_service.ModelPK) (*order_service.Mechanic, error)
}

type TarifService interface{
	Create(context.Context, *order_service.CreateTarif) (*order_service.TarifPK, error)
	GetByID(context.Context, *order_service.TarifPK) (*order_service.Tarif, error)
}
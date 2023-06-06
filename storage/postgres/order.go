package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (c *orderRepo) Create(context.Context, *order_service.CreateOrder) (resp *order_service.OrderPrimaryKey, err error) {
	return
}
func (c *orderRepo) GetByID(context.Context, *order_service.OrderPrimaryKey) (resp *order_service.Order, err error) {
	return
}
func (c *orderRepo) GetList(context.Context, *order_service.GetListOrderRequest) (resp *order_service.GetListOrderResponse, err error) {
	return
}
func (c *orderRepo) Update(context.Context, *order_service.UpdateOrder) (resp int64, err error) {
	return
}
func (c *orderRepo) UpdatePatch(context.Context, *order_service.UpdatePatchOrder) (resp int64, err error) {
	return
}
func (c *orderRepo) Delete(context.Context, *order_service.OrderPrimaryKey) error {
	return nil
}

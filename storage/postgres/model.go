package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type modelRepo struct {
	db *pgxpool.Pool
}

func NewModelRepo(db *pgxpool.Pool) *modelRepo {
	return &modelRepo{
		db: db,
	}
}

func (s *modelRepo) Create(ctx context.Context, req *order_service.CreateModel) (resp *order_service.ModelPK, err error) {
	return
}
func (s *modelRepo) GetByID(ctx context.Context, req *order_service.ModelPK) (resp *order_service.Mechanic, err error) {
	return
}

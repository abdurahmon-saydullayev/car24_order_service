package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type tarifRepo struct {
	db *pgxpool.Pool
}

func NewTarifRepo(db *pgxpool.Pool) *tarifRepo {
	return &tarifRepo{
		db: db,
	}
}

func (s *tarifRepo) Create(ctx context.Context, req *order_service.CreateTarif) (resp *order_service.TarifPK, err error) {
	return
}
func (s *tarifRepo) GetByID(ctx context.Context, req *order_service.TarifPK) (resp *order_service.Tarif, err error) {
	return
}

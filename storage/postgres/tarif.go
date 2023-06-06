package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"
	"database/sql"

	"github.com/google/uuid"
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
	id := uuid.New().String()

	query := `
		INSERT INTO tarif (
			id,
			name,
			model_id,
			price_per_day
		) VALUES ( $1, $2, $3, $4)
	`

	_, err = s.db.Exec(
		ctx,
		query,
		id,
		req.Name,
		req.ModelId,
		req.PricePerDay,
	)
	if err != nil {
		return nil, err
	}

	return &order_service.TarifPK{Id: id}, nil
}
func (s *tarifRepo) GetByID(ctx context.Context, req *order_service.TarifPK) (resp *order_service.Tarif, err error) {
	query := `
		SELECT id,name, model_id,price_per_day
		 FROM "tarif"
		 WHERE id = $1
	`
	resp = &order_service.Tarif{}
	var (
		id            sql.NullString
		name          sql.NullString
		model         sql.NullString
		price_per_day sql.NullString
	)
	err = s.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&model,
		&price_per_day,
	)
	if err != nil {
		return nil, err
	}

	resp = &order_service.Tarif{
		Id:          id.String,
		Name:        name.String,
		ModelId:     model.String,
		PricePerDay: price_per_day.String,
	}

	return
}

func (c *tarifRepo) Delete(ctx context.Context, req *order_service.TarifPK) error {
	query := `DELETE FROM "tarif" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
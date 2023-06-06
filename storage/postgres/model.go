package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"
	"database/sql"

	"github.com/google/uuid"
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
	id := uuid.New().String()

	query := `
		INSERT INTO "model"(
			id,
			name,
			created_at,
			updated_at
		) VALUES($1, $2, NOW(), NOW())
		`

	_, err = s.db.Exec(ctx, query,
		id,
		req.Name,
	)
	if err != nil {
		return nil, err
	}

	return &order_service.ModelPK{Id: id}, nil
}

func (s *modelRepo) GetByID(ctx context.Context, req *order_service.ModelPK) (resp *order_service.Model, err error) {
	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM "model" WHERE id = $1
	`
	resp = &order_service.Model{}
	var (
		id         sql.NullString
		name       sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	err = s.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	resp = &order_service.Model{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}

	return
}

func (c *modelRepo) Delete(ctx context.Context, req *order_service.ModelPK) error {
	query := `DELETE FROM "model" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
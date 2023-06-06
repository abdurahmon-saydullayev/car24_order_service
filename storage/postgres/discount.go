package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type discountRepo struct {
	db *pgxpool.Pool
}

func NewDiscountRepo(db *pgxpool.Pool) *discountRepo {
	return &discountRepo{
		db: db,
	}
}

func (c *discountRepo) Create(ctx context.Context, req *order_service.CreateDiscount) (resp *order_service.DiscountPK, err error) {
	id := uuid.New().String()

	query := `
		INSERT INTO  "discount"(
			id,
			name,
			discount_type,
			created_at,
			updated_at
	) VALUES($1, $2, $3, NOW(), NOW())
	`
	_, err = c.db.Exec(
		ctx,
		query,
		id,
		req.Name,
		req.DiscountType,
	)
	if err != nil {
		return nil, err
	}

	return &order_service.DiscountPK{Id: id}, nil
}
func (c *discountRepo) GetByID(ctx context.Context, req *order_service.DiscountPK) (resp *order_service.Discount, err error) {
	query := `
		SELECT 
			id,
			name,
			discount_type,
			created_at,
			updated_at
		FROM "discount"
		WHERE id = $1
	`
	resp = &order_service.Discount{}
	var (
		id            sql.NullString
		name          sql.NullString
		discount_type sql.NullString
		created_at    sql.NullString
		updated_at    sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&discount_type,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err

	}

	resp = &order_service.Discount{
		Id:           id.String,
		Name:         name.String,
		DiscountType: discount_type.String,
		CreateAt:     discount_type.String,
		UpdateAt:     updated_at.String,
	}

	return
}
func (c *discountRepo) Delete(ctx context.Context, req *order_service.DiscountPK) error {
	query := `
		DELETE FROM "discount" WHERE id = $1
	`
	_, err := c.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}

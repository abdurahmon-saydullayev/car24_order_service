package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type mechanicRepo struct {
	db *pgxpool.Pool
}

func NewMechanicRepo(db *pgxpool.Pool) *mechanicRepo {
	return &mechanicRepo{
		db: db,
	}
}

func (c *mechanicRepo) Create(ctx context.Context, req *order_service.CreateMechanic) (resp *order_service.MechanicPK, err error) {
	id := uuid.New().String()

	query := `
		INSERT INTO  "mechanic"(
			id,
			fullname,
			phone_number,
			photo,
			price_per_hour,
	) VALUES($1, $2, $3, $4, $5)
	`
	_, err = c.db.Exec(
		ctx,
		query,
		id,
		req.Fullname,
		req.PhoneNumber,
		req.Photo,
		req.PricePerHour,
	)
	if err != nil {
		return nil, err
	}

	return &order_service.MechanicPK{Id: id}, nil
}
func (c *mechanicRepo) GetByID(ctx context.Context, req *order_service.MechanicPK) (resp *order_service.Mechanic, err error) {
	query := `
		SELECT 
			id,
			fullname,
			phone_number,
			photo,
			price_per_hour,
		FROM "discount"
		WHERE id = $1 AND status = false
	`
	resp = &order_service.Mechanic{}
	var (
		id             sql.NullString
		fullname       sql.NullString
		phone_number   sql.NullString
		photo          sql.NullString
		price_per_hour sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&fullname,
		&phone_number,
		&photo,
		&price_per_hour,
	)
	if err != nil {
		return nil, err

	}

	resp = &order_service.Mechanic{
		Id:           id.String,
		Fullname:     fullname.String,
		PhoneNumber:  phone_number.String,
		Photo:        photo.String,
		PricePerHour: price_per_hour.String,
	}

	return
}

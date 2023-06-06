package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"Projects/Car24/car24_order_service/models"
	"Projects/Car24/car24_order_service/pkg/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
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

func (c *orderRepo) Create(ctx context.Context, req *order_service.CreateOrder) (resp *order_service.OrderPrimaryKey, err error) {
	id := uuid.New().String()

	order_num, err := helper.GenerateOrderNomer()
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO "order" (
			id,
			car_id,
			client_id,
			tarif_id,
			total_price,
			paid_price,
			day_count,
			start_date,
			discount,
			order_number,
			status,
			mileage,
			is_paid_date,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6,$7, $8, $9, $10,$11, $12, $13, NOW(),NOW())
	`

	_, err = c.db.Exec(
		ctx,
		query,
		id,
		req.CarId,
		req.ClientId,
		req.TarifId,
		req.TotalPrice,
		req.PaidPrice,
		req.DayCount,
		req.StartDate,
		req.Discount,
		order_num,
		req.Status,
		req.Miliage,
		req.IsPaidDate,
	)
	if err != nil {
		return nil, err
	}

	return &order_service.OrderPrimaryKey{Id: id}, nil
}

func (c *orderRepo) GetByID(ctx context.Context, req *order_service.OrderPrimaryKey) (order *order_service.Order, err error) {
	query := `
		SELECT 
			id,
			car_id,
			client_id,
			tarif_id,
			total_price,
			paid_price,
			day_count,
			start_date,
			discount,
			order_number,
			status,
			miliage,
			is_paid_date,
			created_at,
			updated_at
		FROM "order"
		WHERE id = $1
	`

	var (
		id           sql.NullString
		car_id       sql.NullString
		client_id    sql.NullString
		tarif_id     sql.NullString
		total_price  sql.NullFloat64
		paid_price   sql.NullFloat64
		day_count    sql.NullInt32
		start_date   sql.NullString
		discount     sql.NullString
		order_num    sql.NullString
		status       sql.NullBool
		miliage      sql.NullInt32
		is_paid_date sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&car_id,
		&client_id,
		&tarif_id,
		&total_price,
		&paid_price,
		&day_count,
		&start_date,
		&discount,
		&order_num,
		&status,
		&miliage,
		&is_paid_date,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return order, err
	}

	order = &order_service.Order{
		Id:          id.String,
		CarId:       car_id.String,
		TarifId:     tarif_id.String,
		TotalPrice:  total_price.Float64,
		PaidPrice:   paid_price.Float64,
		DayCount:    day_count.Int32,
		StartDate:   start_date.String,
		Discount:    discount.String,
		OrderNumber: order_num.String,
		Status:      status.Bool,
		Miliage:     miliage.Int32,
		IsPaidDate:  is_paid_date.String,
		CreatedAt:   created_at.String,
		UpdatedAt:   updated_at.String,
	}

	return
}

func (c *orderRepo) GetList(ctx context.Context, req *order_service.GetListOrderRequest) (resp *order_service.GetListOrderResponse, err error) {
	resp = &order_service.GetListOrderResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE "
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			car_id,
			client_id,
			tarif_id,
			total_price,
			paid_price,
			day_count,
			start_date,
			discount,
			order_number,
			status,
			miliage,
			is_paid_date,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "order"
	`
	if len(req.GetSearch()) > 0 {
		filter += " AND order_num ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}
	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}
	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)

	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id           sql.NullString
			car_id       sql.NullString
			client_id    sql.NullString
			tarif_id     sql.NullString
			total_price  sql.NullFloat64
			paid_price   sql.NullFloat64
			day_count    sql.NullInt32
			start_date   sql.NullString
			discount     sql.NullString
			order_num    sql.NullString
			status       sql.NullBool
			miliage      sql.NullInt32
			is_paid_date sql.NullString
			created_at   sql.NullString
			updated_at   sql.NullString
		)

		err := rows.Scan(
			&id,
			&car_id,
			&client_id,
			&tarif_id,
			&total_price,
			&paid_price,
			&day_count,
			&start_date,
			&discount,
			&order_num,
			&status,
			&miliage,
			&is_paid_date,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return resp, err
		}

		resp.Orders = append(resp.Orders, &order_service.Order{
			Id:          id.String,
			CarId:       car_id.String,
			TarifId:     tarif_id.String,
			TotalPrice:  total_price.Float64,
			PaidPrice:   paid_price.Float64,
			DayCount:    day_count.Int32,
			StartDate:   start_date.String,
			Discount:    discount.String,
			OrderNumber: order_num.String,
			Status:      status.Bool,
			Miliage:     miliage.Int32,
			IsPaidDate:  is_paid_date.String,
			CreatedAt:   created_at.String,
			UpdatedAt:   updated_at.String,
		})
	}

	return
}

func (c *orderRepo) Update(ctx context.Context, req *order_service.UpdateOrder) (resp int64, err error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"order"
		SET
			car_id = :car_id,
			client_id = :client_id,
			tarif_id = :tarif_id,
			total_price = :total_price,
			paid_price = :paid_price,
			day_count = :day_count,
			start_date = :start_date,
			discount = :discount,
			order_number = :order_number,
			status = :status,
			miliage = :miliage,
			is_paid_date = :is_paid_date,
			updated_at = now()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":           req.GetId(),
		"car_id":       req.GetCarId(),
		"client_id":    req.GetClientId(),
		"tarif_id":     req.GetTarifId(),
		"total_price":  req.GetTotalPrice(),
		"paid_price":   req.GetPaidPrice(),
		"day_count":    req.GetDayCount(),
		"start_date":   req.GetStartDate(),
		"discount":     req.GetDiscount(),
		"order_number": req.GetOrderNumber(),
		"status":       req.GetStatus(),
		"miliage":      req.GetMiliage(),
		"is_paid_date": req.GetIsPaidDate(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *orderRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (resp int64, err error) {
	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Fields) == 0 {
		err = errors.New("no updates provided")
		return
	}

	req.Fields["id"] = req.Id

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Fields)-1 {
			set += ", "
		}
		ind++
	}

	query = `
		UPDATE
			"order"
	` + set + ` , updated_at = now()
		WHERE
			id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), err
}

func (c *orderRepo) Delete(ctx context.Context, req *order_service.OrderPrimaryKey) error {
	query := `DELETE FROM "order" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}

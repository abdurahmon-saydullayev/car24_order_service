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

type carRepo struct {
	db *pgxpool.Pool
}

func NewCarRepo(db *pgxpool.Pool) *carRepo {
	return &carRepo{
		db: db,
	}
}

func (s *carRepo) Create(ctx context.Context, req *order_service.CreateCar) (resp *order_service.CarPrimaryKey, err error) {
	id := uuid.New().String()

	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO "car"(
			id,
			state_number,
			tarif_id,
			model_id,
			created_at,
			updated_at
		) VALUES($1, $2, $3, $4, NOW(),NOW())
	`
	_, err = s.db.Exec(ctx, query,
		id,
		req.StateNumber,
		req.TarifId,
		req.ModelId,
	)
	if err != nil {
		return nil, err
	}

	return &order_service.CarPrimaryKey{Id: id}, nil
}
func (s *carRepo) GetByID(ctx context.Context, req *order_service.CarPrimaryKey) (resp *order_service.Car, err error) {
	query := `
		SELECT id, state_number, tarif_id, model_id, created_at, updated_at
		FROM "car"
		WHERE id = $1
	`
	resp = &order_service.Car{}
	var (
		id           sql.NullString
		state_number sql.NullString
		tarif_id     sql.NullString
		model_id     sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
	)

	err = s.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&state_number,
		&tarif_id,
		&model_id,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	resp = &order_service.Car{
		Id:          id.String,
		StateNumber: state_number.String,
		TarifId:     tarif_id.String,
		ModelId:     model_id.String,
		CreatedAt:   created_at.String,
		UpdatedAt:   updated_at.String,
	}

	return
}

func (s *carRepo) GetList(ctx context.Context, req *order_service.GetListCarRequest) (resp *order_service.GetListCarResponse, err error) {
	resp = &order_service.GetListCarResponse{}
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
			state_number,
			tarif_id,
			model_id,
			created_at,
			updated_at
		FROM "car"
	`
	if len(req.GetSearch()) > 0 {
		filter += " AND state_number ILIKE '%' || '" + req.Search + "' || '%' "
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
	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id           sql.NullString
			state_number sql.NullString
			tarif_id     sql.NullString
			model_id     sql.NullString
			created_at   sql.NullString
			updated_at   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&state_number,
			&model_id,
			&tarif_id,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}

		resp.Cars = append(resp.Cars, &order_service.Car{
			Id:          id.String,
			StateNumber: state_number.String,
			ModelId:     model_id.String,
			TarifId:     tarif_id.String,
			CreatedAt:   created_at.String,
			UpdatedAt:   updated_at.String,
		})
	}

	return
}

func (s *carRepo) Update(ctx context.Context, req *order_service.UpdateCar) (resp int64, err error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"car"
		SET
			state_number = :state_number,
			model_id = :model_id,
			tarif_id = :tarif_id,
			updated_at = now()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":           req.GetId(),
		"state_number": req.GetStateNumber(),
		"model_id":     req.GetModelId(),
		"tarif_id":     req.GetTarifId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (s *carRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (resp int64, err error) {
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
			"car"
	` + set + ` , updated_at = now()
		WHERE
			id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), err
}

func (s *carRepo) Delete(ctx context.Context, req *order_service.CarPrimaryKey) error {
	query := `DELETE FROM "car" WHERE id = $1`

	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}
	return nil
}

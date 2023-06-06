package postgres

import (
	"Projects/Car24/car24_order_service/config"
	"Projects/Car24/car24_order_service/storage"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	order    storage.OrderRepoI
	discount storage.DiscountRepoI
	mechanic storage.MechnicRepoI
	model    storage.ModelRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pool,
		order:    NewOrderRepo(pool),
		discount: NewDiscountRepo(pool),
		mechanic: NewMechanicRepo(pool),
		model:    NewModelRepo(pool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Order() storage.OrderRepoI {
	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}
	return s.order
}

func (s *Store) Discount() storage.DiscountRepoI {
	if s.discount == nil {
		s.discount = NewDiscountRepo(s.db)
	}
	return s.discount
}

func (s *Store) Mechanic() storage.MechnicRepoI {
	if s.mechanic == nil {
		s.mechanic = NewMechanicRepo(s.db)
	}
	return s.mechanic
}

func (s *Store) Model() storage.ModelRepoI {
	if s.model == nil {
		s.model = NewModelRepo(s.db)
	}
	return s.model
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

package postgres

import (
	"Projects/Car24/car24_order_service/genproto/order_service"
	"context"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Test_orderRepo_Create(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		req *order_service.CreateOrder
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *order_service.OrderPrimaryKey
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &orderRepo{
				db: tt.fields.db,
			}
			gotResp, err := c.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("orderRepo.Create() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

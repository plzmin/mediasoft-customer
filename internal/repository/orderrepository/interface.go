package orderrepository

import (
	"context"
	"mediasoft-customer/internal/model"
)

//go:generate mockery --all

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
}

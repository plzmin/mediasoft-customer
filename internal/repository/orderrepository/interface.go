package orderrepository

import (
	"context"
	"mediasoft-customer/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
}

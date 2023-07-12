package officerepository

import (
	"context"
	"mediasoft-customer/internal/model"
)

//go:generate mockery --all

type OfficeRepository interface {
	Create(ctx context.Context, office *model.Office) error
	List(ctx context.Context) ([]*model.Office, error)
}

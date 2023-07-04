package officerepository

import (
	"context"
	"mediasoft-customer/internal/model"
)

type OfficeRepository interface {
	Create(ctx context.Context, o *model.Office) error
	List(ctx context.Context) ([]*model.Office, error)
}

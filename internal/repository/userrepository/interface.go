package userrepository

import (
	"context"
	"github.com/google/uuid"
	"mediasoft-customer/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	ListByOfficeUuid(ctx context.Context, uuid uuid.UUID) ([]*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
}

package userrepository

import (
	"context"
	"mediasoft-customer/internal/model"
)

//go:generate mockery --all

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	List(ctx context.Context, uuid string) ([]*model.User, error)
}

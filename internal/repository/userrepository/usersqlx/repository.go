package usersqlx

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"mediasoft-customer/internal/model"
)

type UserSQLx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *UserSQLx {
	return &UserSQLx{db: db}
}

func (r *UserSQLx) Create(ctx context.Context, u *model.User) error {
	const q = `insert into users (uuid, name, office_uuid, created_at) values (:uuid, :name,:office_uuid,:created_at)`
	_, err := r.db.NamedExecContext(ctx, q, u)
	return err
}

func (r *UserSQLx) ListByOfficeUuid(ctx context.Context, uuid uuid.UUID) ([]*model.User, error) {
	const q = `select users.*, o.name AS office_name from users join offices o on o.uuid = users.office_uuid where users.office_uuid = $1`
	var list []*model.User
	err := r.db.SelectContext(ctx, &list, q, uuid)
	return list, err
}

func (r *UserSQLx) List(ctx context.Context) ([]*model.User, error) {
	const q = `select users.*, o.name AS office_name from users join offices o on o.uuid = users.office_uuid`
	var list []*model.User
	err := r.db.SelectContext(ctx, &list, q)
	return list, err
}

package officesqlx

import (
	"context"
	"github.com/jmoiron/sqlx"
	"mediasoft-customer/internal/model"
)

type OfficeSqlx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *OfficeSqlx {
	return &OfficeSqlx{db: db}
}

func (r *OfficeSqlx) Create(ctx context.Context, office *model.Office) error {
	const q = `insert into offices (uuid, name, address) values (:uuid,:name,:address)`
	_, err := r.db.NamedExecContext(ctx, q, office)
	return err
}

func (r *OfficeSqlx) List(ctx context.Context) ([]*model.Office, error) {
	const q = `select * from offices`
	var list []*model.Office
	err := r.db.SelectContext(ctx, &list, q)
	return list, err
}

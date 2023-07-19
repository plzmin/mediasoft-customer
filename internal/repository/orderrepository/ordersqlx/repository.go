package ordersqlx

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"mediasoft-customer/internal/model"
)

type OrderSqlx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *OrderSqlx {
	return &OrderSqlx{db: db}
}

func (r *OrderSqlx) Create(ctx context.Context, order *model.Order) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	const q = `insert into orders (uuid, user_uuid) values (:uuid, :user_uuid)`
	_, err = tx.NamedExecContext(ctx, q, order)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errors.Join(err, errRollback)
		}
		return err
	}

	oiq := `insert into order_item(order_uuid, count, product_uuid) values ($1,$2,$3)`
	for _, orderItems := range [][]*customer.OrderItem{order.Salads, order.Drinks, order.Meats, order.Desserts, order.Soups} {
		for _, orderItem := range orderItems {
			_, err := tx.ExecContext(ctx, oiq, order.Uuid, orderItem.Count, orderItem.ProductUuid)
			if err != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					return errors.Join(err, errRollback)
				}
				return err
			}
		}
	}

	return tx.Commit()
}

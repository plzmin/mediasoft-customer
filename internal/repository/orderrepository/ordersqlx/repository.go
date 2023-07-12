package ordersqlx

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
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
	const q = `insert into orders (uuid, user_uuid, created_at) values (:uuid, :user_uuid, :created_at)`
	_, err = tx.NamedExec(q, order)
	if err != nil {
		tx.Rollback()
		return err
	}

	var oiq = `insert into order_item(order_uuid, count, product_uuid) values `
	for _, orderItems := range [][]*model.OrderItem{order.Salads, order.Drinks, order.Meats, order.Desserts, order.Soups} {
		for _, orderItem := range orderItems {
			oiq += fmt.Sprintf("($%s,$%d,$%s),", order.Uuid, orderItem.Count, orderItem.ProductUuid)
		}
	}
	_, err = tx.Exec(oiq[:len(oiq)-1])
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return err
}

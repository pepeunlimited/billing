package ordersrepo

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
	"log"
	"time"
)

type OrdersRepository interface {

	CreateOrder(ctx context.Context, userID int64, items []*ent.Item) (*ent.Orders, error)
	CreateOrderStatus(ctx context.Context, orderID int, userID int64)

	GetOrderByUserID(ctx context.Context, orderID int, userID int64)

	Wipe(ctx context.Context)
}

type ordersMySQL struct {
	client *ent.Client
}

func (mysql ordersMySQL) Wipe(ctx context.Context) {
	mysql.client.Txs.Delete().ExecX(ctx)
	mysql.client.Item.Delete().ExecX(ctx)
	mysql.client.Orders.Delete().ExecX(ctx)
}

func (mysql ordersMySQL) GetOrderByUserID(ctx context.Context, orderID int, userID int64) {
	panic("implement me")
}

func (mysql ordersMySQL) CreateOrder(ctx context.Context, userID int64, items []*ent.Item) (*ent.Orders, error) {
	tx, err := mysql.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	order, err := tx.Orders.Create().SetUserID(userID).SetCreatedAt(now).Save(ctx)
	if err != nil {
		mysql.rollback(tx)
		return nil, err
	}
	_, err = tx.
		Txs.
		Create().
		SetCreatedAt(now).
		SetStatus("CREATED").
		SetOrdersID(order.ID).
		Save(ctx)
	if err != nil {
		mysql.rollback(tx)
		return nil, err
	}
	if items == nil {
		if err :=  mysql.commit(tx); err != nil {
			return nil, err
		}
		return order, nil
	}
	for _, item := range items {
		_, err := tx.Item.Create().SetOrdersID(order.ID).SetPriceID(item.PriceID).SetQuantity(item.Quantity).Save(ctx)
		if err != nil {
			mysql.rollback(tx)
			return nil, err
		}
	}
	if err :=  mysql.commit(tx); err != nil {
		return nil, err
	}
	return mysql.client.Orders.Query().Where(orders.ID(order.ID)).WithItems().Only(ctx)
}

func (mysql ordersMySQL) CreateOrderStatus(ctx context.Context, orderID int, userID int64) {
	panic("implement me")
}

func NewOrdersRepository(client *ent.Client) OrdersRepository {
	return ordersMySQL{client:client}
}

func (mysql ordersMySQL) rollback(tx *ent.Tx){
	if err := tx.Rollback(); err != nil {
		log.Print("orders: failed execute rollback.."+err.Error())
	}
}

func (mysql ordersMySQL) commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		log.Print("orders: failed execute commit.."+err.Error())
		return err
	}
	return nil
}
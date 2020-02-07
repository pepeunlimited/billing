package ordersrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
	"log"
	"time"
)

var (
	ErrOrderNotExist = errors.New("orders: not exist")
)

type OrdersRepository interface {

	CreateOrder(ctx context.Context, userID int64, items []*ent.Item) 					  (*ent.Orders, error)

	GetOrderByUserID(ctx context.Context, orderID int, userID int64)					  (*ent.Orders, error)
	GetOrdersByUserID(ctx context.Context, userID int64, pageToken int64, pageSize int32) ([]*ent.Orders, int64, error)

	Wipe(ctx context.Context)
}

type ordersMySQL struct {
	client *ent.Client
}

func (mysql ordersMySQL) GetOrdersByUserID(ctx context.Context, userID int64, pageToken int64, pageSize int32) ([]*ent.Orders, int64, error) {
	order, err := mysql.client.Orders.Query().Where(
		orders.UserID(userID),
		orders.IDGT(int(pageToken))).
		Order(ent.Asc(orders.FieldID)).
		Limit(int(pageSize)).WithItems().WithPayments().WithItems().All(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(order) == 0 {
		return []*ent.Orders{}, pageToken, nil
	}
	return order, int64(order[len(order) - 1].ID), nil
}

func (mysql ordersMySQL) Wipe(ctx context.Context) {
	mysql.client.Txs.Delete().ExecX(ctx)
	mysql.client.Item.Delete().ExecX(ctx)
	mysql.client.Orders.Delete().ExecX(ctx)
}

func (mysql ordersMySQL) GetOrderByUserID(ctx context.Context, orderID int, userID int64) (*ent.Orders, error) {
	order, err := mysql.client.Orders.Query().Where(orders.UserID(userID), orders.ID(orderID)).WithPayments().WithItems().WithTxs().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrOrderNotExist
		}
		return nil, err
	}
	return order, nil
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
		SetStatus(StatusFromString("created").String()).
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
	return mysql.GetOrderByUserID(ctx, order.ID, userID)
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
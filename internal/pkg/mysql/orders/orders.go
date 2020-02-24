package orders

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

	CreateOrderStatus(ctx context.Context, orderID int, userID int64, status Status)	  (*ent.Orders, error)
	CreateOrder(ctx context.Context, userID int64, items []*ent.Item) 					  (*ent.Orders, error)

	GetOrderByUserID(ctx context.Context, orderID int, userID int64, withItems bool, withTXs bool, withPayments bool) (*ent.Orders, error)
	GetOrdersByUserID(ctx context.Context, userID int64, pageToken int64, pageSize int32, withItems bool, withTXs bool, withPayments bool) ([]*ent.Orders, int64, error)

	Wipe(ctx context.Context)
}

type ordersMySQL struct {
	client *ent.Client
}

func (mysql ordersMySQL) CreateOrderStatus(ctx context.Context, orderID int, userID int64, status Status) (*ent.Orders, error) {
	_, err := mysql.GetOrderByUserID(ctx, orderID, userID, false, false, false)
	if err != nil {
		return nil, err
	}
	_, err = mysql.client.Txs.Create().SetOrdersID(orderID).SetCreatedAt(time.Now().UTC()).SetStatus(status.String()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.GetOrderByUserID(ctx, orderID, userID, false, false, false)
}

func (mysql ordersMySQL) GetOrdersByUserID(ctx context.Context, userID int64, pageToken int64, pageSize int32, withItems bool, withTXs bool, withPayments bool) ([]*ent.Orders, int64, error) {
	query := mysql.client.Orders.Query().Where(
		orders.UserID(userID),
		orders.IDGT(int(pageToken))).
		Order(ent.Asc(orders.FieldID)).
		Limit(int(pageSize))
	if withItems {
		query.WithItems()
	}
	if withPayments {
		query.WithPayments()
	}
	if withTXs {
		query.WithTxs()
	}
	order, err := query.All(ctx)
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
	mysql.client.Payment.Delete().ExecX(ctx)
	mysql.client.Orders.Delete().ExecX(ctx)
	mysql.client.Instrument.Delete().ExecX(ctx)
}

func (mysql ordersMySQL) GetOrderByUserID(ctx context.Context, orderID int, userID int64, withItems bool, withTXs bool, withPayments bool) (*ent.Orders, error) {
	query := mysql.client.Orders.Query().Where(orders.UserID(userID), orders.ID(orderID))
	if withItems {
		query.WithItems()
	}
	if withPayments {
		query.WithPayments()
	}
	if withTXs {
		query.WithTxs()
	}
	order, err := query.Only(ctx)
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
	return mysql.GetOrderByUserID(ctx, order.ID, userID, true, true, false)
}

func New(client *ent.Client) OrdersRepository {
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
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleInstrument() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the instrument's edges.
	pa0 := client.Payment.
		Create().
		SaveX(ctx)
	log.Println("payment created:", pa0)

	// create instrument vertex with its edges.
	i := client.Instrument.
		Create().
		SetType("string").
		SetTypeI18nID(1).
		AddPayments(pa0).
		SaveX(ctx)
	log.Println("instrument created:", i)

	// query edges.
	pa0, err = i.QueryPayments().First(ctx)
	if err != nil {
		log.Fatalf("failed querying payments: %v", err)
	}
	log.Println("payments found:", pa0)

	// Output:
}
func ExampleItem() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the item's edges.

	// create item vertex with its edges.
	i := client.Item.
		Create().
		SetPriceID(1).
		SetQuantity(1).
		SaveX(ctx)
	log.Println("item created:", i)

	// query edges.

	// Output:
}
func ExampleOrders() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the orders's edges.
	t0 := client.Txs.
		Create().
		SetStatus("string").
		SetCreatedAt(time.Now()).
		SaveX(ctx)
	log.Println("txs created:", t0)
	i1 := client.Item.
		Create().
		SetPriceID(1).
		SetQuantity(1).
		SaveX(ctx)
	log.Println("item created:", i1)

	// create orders vertex with its edges.
	o := client.Orders.
		Create().
		SetCreatedAt(time.Now()).
		SetUserID(1).
		AddTxs(t0).
		AddItems(i1).
		SaveX(ctx)
	log.Println("orders created:", o)

	// query edges.
	t0, err = o.QueryTxs().First(ctx)
	if err != nil {
		log.Fatalf("failed querying txs: %v", err)
	}
	log.Println("txs found:", t0)

	i1, err = o.QueryItems().First(ctx)
	if err != nil {
		log.Fatalf("failed querying items: %v", err)
	}
	log.Println("items found:", i1)

	// Output:
}
func ExamplePayment() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the payment's edges.
	o0 := client.Orders.
		Create().
		SetCreatedAt(time.Now()).
		SetUserID(1).
		SaveX(ctx)
	log.Println("orders created:", o0)

	// create payment vertex with its edges.
	pa := client.Payment.
		Create().
		SetOrders(o0).
		SaveX(ctx)
	log.Println("payment created:", pa)

	// query edges.
	o0, err = pa.QueryOrders().First(ctx)
	if err != nil {
		log.Fatalf("failed querying orders: %v", err)
	}
	log.Println("orders found:", o0)

	// Output:
}
func ExamplePlan() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the plan's edges.
	s0 := client.Subscription.
		Create().
		SetUserID(1).
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SaveX(ctx)
	log.Println("subscription created:", s0)

	// create plan vertex with its edges.
	pl := client.Plan.
		Create().
		SetTitleI18nID(1).
		SetPriceID(1).
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SetLength(1).
		SetUnit("string").
		AddSubscriptions(s0).
		SaveX(ctx)
	log.Println("plan created:", pl)

	// query edges.
	s0, err = pl.QuerySubscriptions().First(ctx)
	if err != nil {
		log.Fatalf("failed querying subscriptions: %v", err)
	}
	log.Println("subscriptions found:", s0)

	// Output:
}
func ExampleSubscription() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the subscription's edges.

	// create subscription vertex with its edges.
	s := client.Subscription.
		Create().
		SetUserID(1).
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SaveX(ctx)
	log.Println("subscription created:", s)

	// query edges.

	// Output:
}
func ExampleTxs() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the txs's edges.

	// create txs vertex with its edges.
	t := client.Txs.
		Create().
		SetStatus("string").
		SetCreatedAt(time.Now()).
		SaveX(ctx)
	log.Println("txs created:", t)

	// query edges.

	// Output:
}

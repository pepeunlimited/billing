package main

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/server/twirp"
	"github.com/pepeunlimited/billing/pkg/rpc/order"
	"github.com/pepeunlimited/billing/pkg/rpc/payment"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"log"
	"net/http"
)

const (
	Version = "0.0.1"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Printf("Starting the BillingSever... version=[%v]", Version)

	client := ent.NewEntClient()
	os := order.NewOrderServiceServer(twirp.NewOrderServer(client), nil)
	ps := payment.NewPaymentServiceServer(twirp.NewPaymentServer(client), nil)

	mux := http.NewServeMux()
	mux.Handle(os.PathPrefix(), middleware.Adapt(os))
	mux.Handle(ps.PathPrefix(), middleware.Adapt(ps))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}
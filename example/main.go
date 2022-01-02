package main

import (
	"context"
	"log"
	"time"

	"github.com/TomerHeber/avatax-v2-go/client"
	"github.com/TomerHeber/avatax-v2-go/client/transactions"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	basicAuth := httptransport.BasicAuth("username", "password")
	r := httptransport.New("rest.avatax.com", "/api/v2", client.DefaultSchemes)
	r.DefaultAuthentication = basicAuth
	r.Producers["application/*+json"] = runtime.JSONProducer()
	c := client.New(r, strfmt.Default)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := c.Transactions.GetTransactionByID(&transactions.GetTransactionByIDParams{
		ID:      1000,
		Context: ctx,
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", *res)
}

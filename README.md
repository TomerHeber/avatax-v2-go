# avatax-v2-go
Go SDK for Avalara's Avatax V2 Rest API

The code was generated with [go-swagger](https://github.com/go-swagger/go-swagger) (See the script under ./scripts).  

The client should support 3 types of authentication. Check [here](https://goswagger.io/generate/client.html#authentication) for additional details.

## Basic Example

```
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
```
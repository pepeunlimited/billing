# cURL

### Install
```
$ brew install jq > curl ... | jq
```
##### GetOrders
```
$ curl -H "Content-Type: application/json" \
  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.OrderService/GetOrders" \
  -d '{"user_id": 1, "page_size": 20, "page_token": 0}'
```
##### GetOrder
```
curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.OrderService/GetOrder" \
-d '{"user_id": 1, "order_id": 2}'
```
##### CreateOrder
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.OrderService/CreateOrder" \
-d '{"user_id": 1, "order_items": [{"quantity": 1, "price_id":1}]}'
```
##### GetOrderItems
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.OrderService/GetOrderItems" \
-d '{"user_id": 1, "order_id": 1}'
```
##### GetOrderTxs
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.OrderService/GetOrderTxs" \
-d '{"user_id": 1, "order_id": 1}'
```
##### CreatePayment
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.PaymentService/CreatePayment" \
-d '{"user_id": 1, "order_id": 1, "payment_instrument_id": 1}'
```
##### GetPaymentInstrument
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.PaymentService/GetPaymentInstrument" \
-d '{"type":"GOOGLE"}'
```
```
curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.PaymentService/GetPaymentInstrument" \
-d '{"id":1}'©©
```
##### GetPaymentInstruments
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.PaymentService/GetPaymentInstruments" \
-d '{}'
```
##### GetPayments
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.billing.PaymentService/GetPayments" \
-d '{"user_id":1, "page_size": 20, "page_token": 0}'
```

rpc GetPayments(GetPaymentsParams)                         returns (GetPaymentsResponse);
rpc GetPayment(GetPaymentParams)                           returns (Payment);
rpc CreatePaymentInstrument(CreatePaymentInstrumentParams) returns (PaymentInstrument);
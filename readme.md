# Case 3 - simple e-commerce

to start the application

make sure port 6379, 5432, 8080 is free. and copy env.example into docker.env. and run this into terminal

```shell
docker compose up -d --build
```

make sure the application with this ping API

```curl
curl --location 'http://localhost:8080/ping'
```

if te response is `pong`. The Server is running

## create user with email, name. and password

```curl
curl --location 'http://localhost:8080/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"test@test.com",
    "name":"test",
    "password":"test@test.com"
}'
```

and the result will be

```json
{
  "result": "created"
}
```

## login user with password

```plain
curl --location 'http://localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"test@test.com",
    "password":"test@test.com"
}'
```

is login is success, the response will send you the token, and save it into TOKEN variable for before making authenticated api calls

```json
{
  "result": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJuZ3VqaS1hdXRoIiwiZXhwIjoxNzI2MjgzNDIyLCJpYXQiOjE3MjYxOTcwMjIsImp0aSI6IjAxOTFlOTVjLTNlYzYtN2Q3Zi05ZjE1LWU2ZTYyM2JhMGExYiJ9.KIXuX8kvlauziBPWcPbQOKeDMThbLJdZ4maUoI-l5Lo"
}
```

## access the movies

```curl
curl --location 'http://localhost:8080/api/movies'
```

and the response will be like this

```json
{
  "data": [
    {
      "id": 1,
      "name": "Lava A55",
      "manufacturer": "Lava",
      "price": 875837,
      "image": "http://dummyimage.com/164x257.png/cc0000/ffffff",
      "stock": 25
    },
    {
      "id": 2,
      "name": "Nokia 2.2",
      "manufacturer": "Nokia",
      "price": 160441,
      "image": "http://dummyimage.com/246x332.png/ff4444/ffffff",
      "stock": 28
    }
  ]
}
```

## add movie to cart use this api

The process of adding movie into cart is added one by one, so the experience will be smoother

```curl
curl --location 'http://localhost:8080/api/carts' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {{TOKEN}}' \
--data '{
    "movie_id": 3
}'
```

and the result will be

```json
{
  "result": "ok"
}
```

## check the availability of the cart and the movie

```curl
curl --location 'http://localhost:8080/api/carts' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJuZ3VqaS1hdXRoIiwiZXhwIjoxNzI2MjgzNDIyLCJpYXQiOjE3MjYxOTcwMjIsImp0aSI6IjAxOTFlOTVjLTNlYzYtN2Q3Zi05ZjE1LWU2ZTYyM2JhMGExYiJ9.KIXuX8kvlauziBPWcPbQOKeDMThbLJdZ4maUoI-l5Lo'
```

and the cart response is calculating with the actual stock in movie

```json
{
  "data": [
    {
      "id": 1,
      "quantity": 1,
      "message": "stock AVAILABLE for this movie Infinix Hot 10",
      "movie_id": 3,
      "movie_name": "Infinix Hot 10",
      "movie_image": "http://dummyimage.com/202x396.png/5fa2dd/ffffff",
      "movie_price": 504545,
      "actual_stock": 17
    }
  ],
  "result": "ok"
}
```

there is 3 conditions for the cart quantity with the actual stock
| string | send UserID |
|---------|------------------------------------------------------|
| NOT_AVAILABLE | the current stock is 0 |
| NOT_ENOUGH | the cart quantity is more than stock, the order cannot be created |
| AVAILABLE | stock is available |

and the message will be like this

```json
 "message": "stock ${STOCK_CONDITION} for this movie Infinix Hot 10"
```

## check the availability of the cart and the movie

```curl
curl --location 'http://localhost:8080/api/carts' \
--header 'Authorization: Bearer {{TOKEN}}'
```

and the cart response is calculating with the actual stock in movie

```json
{
  "data": [
    {
      "id": 1,
      "quantity": 1,
      "message": "stock AVAILABLE for this movie Infinix Hot 10",
      "movie_id": 3,
      "movie_name": "Infinix Hot 10",
      "movie_image": "http://dummyimage.com/202x396.png/5fa2dd/ffffff",
      "movie_price": 504545,
      "actual_stock": 17
    }
  ],
  "result": "ok"
}
```

there is 3 conditions for the cart quantity with the actual stock
| string | send UserID |
|---------|------------------------------------------------------|
| NOT_AVAILABLE | the current stock is 0 |
| NOT_ENOUGH | the cart quantity is more than stock, the order cannot be created |
| AVAILABLE | stock is available |

and the message will be like this

```json
 "message": "stock ${STOCK_CONDITION} for this movie Infinix Hot 10"
```

## checkout the desired cart

```curl
curl --location 'http://localhost:8080/api/carts/checkout' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {{TOKEN}}' \
--data '{
    "cart_ids": [12]
}'
```

if the cart id is invalid, the response will be

```json
{
  "result": "error",
  "error": "CartIDs not valid"
}
```

and if the cart is valid, the response will be

```json
{
  "result": "ok"
}
```

## Ordering process

here's the ordering process

1. User add movies into cart, and user processing cart after choosing desired cart.
2. User checkout the cart and hitting `/api/cart/checkout` API with sending selected `cart_ids`
3. The server will the `cart_ids`, and The order will be created when stock is available. This process can capture a race condition when making an Order
4. So the Order will be saved with process `PROCESSING`, and the server will create a `OrderPublisher` to process orders one by one
5. So the Order will be consumed by `OrderConsumer` one by one to acquire the real stock
6. If stock is available the Order status will be changed from `PROCESSING` to `WAITING_FOR_PAYMENT`
7. If stock is not available the Order status will be changed from `PROCESSING` to `CANCELED_BECAUSE_LIMITED_STOCK`

## Order Statuses

| name                           | condition                                                          | previous status | next statuses   | implemented | remarks                                             |
| ------------------------------ | ------------------------------------------------------------------ | --------------- | --------------- | ----------- | --------------------------------------------------- |
| PROCESSING                     | cart waiting to be proceed, cart will be processed by Redis PubSub | -               | WAITING_PAYMENT | true        |                                                     |
| WAITING_PAYMENT                | cart can be proceed                                                | PROCESSING      | SUCCESS, FAILED | true        |                                                     |
| CANCELED_BECAUSE_LIMITED_STOCK | stock is less than cart quantity                                   | PROCESSING      | -               | true        |                                                     |
| SUCCESS                        | order success                                                      | WAITING_PAYMENT | -               | false       | can be implemented with creating payment simulation |
| FAILED                         | cancelling order                                                   | WAITING_PAYMENT | -               | false       | can be implemented with creating payment simulation |

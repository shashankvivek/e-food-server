Cmd to generate new changes of swagger :

> swagger generate server -f api/swagger.yaml --default-scheme http --exclude-main

To run the repository locally :

> go run cmd\e-food-server\main.go --scheme http --port=8080

To generate a coupon, make a `POST` call to below endpoint:
> http://127.0.0.1:8080/v1/generateCouponCode
> curl --location --request POST 'http://127.0.0.1:8080/v1/generateCouponCode'


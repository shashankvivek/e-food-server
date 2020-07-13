Cmd to generate new changes of swagger :

> swagger generate server -f api/swagger.yaml --default-scheme http --exclude-main

To run the repository locally :

> go run cmd\e-food-server\main.go --scheme http --port=8080

To generate a coupon, make a `POST` call to below endpoint:
> http://127.0.0.1:8080/v1/generateCouponCode
> curl --location --request POST 'http://127.0.0.1:8080/v1/generateCouponCode' \
  --header 'Cookie: guest_session=0db25c05-f799-4458-b394-3ac48f648609'


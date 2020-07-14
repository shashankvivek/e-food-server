# Intro

e-food is a dummy e-commerce website. The service has been written in **Go** and I have used **go-swagger** framework. *go-swagger* is golang implementation of Swagger 2.0 (aka OpenAPI 2.0), it knows how to serialize and deserialize swagger specifications. 

The go-swagger forces a developer to create documentation while developing an API. The boilerplate code is very helpful in enforcing micro service architecture code style. 

The `swagger.yaml` of this code base can be found at `./api/swagger.yaml`. The documentation can be seen by pasting `swagger.yaml` into [swagger editor online](https://editor.swagger.io/)

The UI for this backend service is available as [e-food-client](https://github.com/shashankvivek/e-food-client/tree/master/e-food) on github.

Reference: 
1. https://github.com/go-swagger/go-swagger
2. https://goswagger.io/install.html


It includes:

1. Guest cart session. (After closing the browser, the app maintain the list of products for guest user)
2. User Registration 
3. User Login using `JWT` with validity set to 300 mins
4. Cart Rules:
      - If 7 or more apples are added to the cart, a 10% discount is applied to all apples
      -  For each set of 4 pears and 2 bananas, a 30% discount is applied, to each set.
      -  These sets must be added to their own cart item entry.
      -  If pears or bananas already exist in the cart, this discount must be recalculated when new pears or bananas are added
5. Coupons can be generated. To generate a coupon, make a `POST` call to below endpoint:
     > http://127.0.0.1:8080/v1/generateCouponCode
      - coupon code can be used to get a 30% discount on oranges, if applied to the cart, otherwise oranges are full price
      - the `/generateCouponCode` can easily be configured to generate different coupons with different product rules
      - a coupon can only be applied once (this is also configurable in `/generateCouponCode` which accepts `userLimit` to restrict the number of times a coupon can be consumed)          
      - Has a configurable expiry timeout (90 seconds (10 sec is too short to test) has been hard-coded in API for testing purposes) once generated.
6. Checkout cart shows:
      - Total price
      - Total Saving
7. In a cart you can :
      - Adjust quantity.
      - Delete items from the cart.
      - Apply coupons. ( option available during checkout time only )
8. A test payment gateway has been provided to complete the journey. ( **https://razorpay.com/**)
9.  Architecture diagram can be found at `./resources/e-food.drawio` and can be opened using http://draw.io/
10. Order Management, Inventory management & User address have not been implemented yet.  
11. A sample unit test file can be found at `./handlers/delete_guest_cart_item_test.go`
12. **In case a new rule has to be added for cart items, make sure the `productId` defined in `rules.json` and `ecommerce.product` table are in sync**
     
Cmd to generate new changes of swagger :

> swagger generate server -f api/swagger.yaml --default-scheme http --exclude-main

To run the repository locally :

> go run cmd\e-food-server\main.go --scheme http --port=8080



# Prerequisite:
* Install mysql
* create a schema named `ecommerce` with:  [ TODO: We can create a separate user rather than using `root`]
> username = `root`,
> password = `root`
 
* Connect to schema and execute file `create.sql` located in repository under path: 
 > ./resources/create.sql

* Install dependencies using `go.mod` file
-----------------------
**Note**: Other areas of improvement has been noted down into `to-do.txt` file of this repo
 

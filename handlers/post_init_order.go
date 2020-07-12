package handlers

import (
	"database/sql"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
	"e-food/restapi/operations/user"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/razorpay/razorpay-go"
	"strconv"
)

type initPaymentOrderImpl struct {
	razorClient *razorpay.Client
	dbClient    *sql.DB
}

func NewUserInitPayHandler(razorClient *razorpay.Client, db *sql.DB) user.InitPayHandler {
	return &initPaymentOrderImpl{
		razorClient: razorClient,
		dbClient:    db,
	}
}

func (impl *initPaymentOrderImpl) Handle(params user.InitPayParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewInitPayInternalServerError().WithPayload("error in parsing token")
	}
	cartId, _, err := dao.CreateOrGetCartDetails(impl.dbClient, email.(string))
	if err != nil {
		return user.NewInitPayInternalServerError().WithPayload("error fetching cart ID")
	}
	cartItems, _, err := dao.GetCustomerCart(impl.dbClient, email.(string))
	if err != nil {
		return user.NewInitPayInternalServerError().WithPayload("error fetching cart items")
	}
	if len(cartItems) < 1 {
		return user.NewInitPayNotFound()
	}
	data := map[string]interface{}{
		"amount":          params.PreOrder.Amount,
		"currency":        "INR",
		"receipt_id":      strconv.FormatInt(cartId, 10),
		"payment_capture": 1,
	}
	//data := map[string]interface{}{
	//	"amount":          1234,
	//	"currency":        "INR",
	//	"receipt_id":      "some_receipt_id",
	//	"payment_capture": 1,
	//}
	xtra := map[string]string{
		"content-type": "application/json",
	}
	_, err = impl.razorClient.Order.Create(data, xtra)
	if err != nil {
		fmt.Println(err)
	}
	// returning hardcoded value as their is a known bug:
	// https://github.com/razorpay/razorpay-go/issues/14
	return user.NewInitPayOK().WithPayload(&user.InitPayOKBody{ID: "order_FDjBRxhFaBvO5L"})
}

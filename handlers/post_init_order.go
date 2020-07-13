package handlers

import (
	"database/sql"
	"e-food/api/restapi/operations/user"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
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
		"receipt":         strconv.FormatInt(cartId, 10),
		"payment_capture": 1,
	}
	body, err := impl.razorClient.Order.Create(data, nil)
	if err != nil {
		fmt.Println(err)
		return user.NewInitPayInternalServerError().WithPayload("error with payment gateway")
	}
	return user.NewInitPayOK().WithPayload(&user.InitPayOKBody{ID: body["id"].(string)})
}

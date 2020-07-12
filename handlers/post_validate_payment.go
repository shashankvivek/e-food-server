package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"e-food/constants"
	"e-food/models"
	"e-food/restapi/operations/user"
	"encoding/hex"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/razorpay/razorpay-go"
)

type validatePayImpl struct {
	client   *razorpay.Client
	dbClient *sql.DB
}

func NewUserPostValidatePaymentHandler(client *razorpay.Client, dbClient *sql.DB) user.PostValidatePaymentHandler {
	return &validatePayImpl{
		client:   client,
		dbClient: dbClient,
	}
}

func (impl *validatePayImpl) Handle(params user.PostValidatePaymentParams, p interface{}) middleware.Responder {

	data := params.PreOrder.RazorpayOrderID + "|" + params.PreOrder.RazorpayPaymentID
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(constants.MyRazorSecret))

	// Write Data to it
	_, err := h.Write([]byte(data))
	if err != nil {
		return user.NewPostValidatePaymentInternalServerError().WithPayload("error decoding")
	}

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	isSuccess := sha == params.PreOrder.RazorpaySignature
	if isSuccess {
		// fetch cart Detail from Razor pay
		body, err := impl.client.Order.Fetch(params.PreOrder.RazorpayOrderID, nil, nil)
		if err != nil {
			fmt.Println(err.Error())
			// todo: return a response when some error happens while calling
			// Razorpay endpoint
		}
		cartId := body["receipt"]
		fmt.Println(cartId)
		// convert cart to order

		// delete cart data from cart table
	}

	//return orderId
	return user.NewPostValidatePaymentOK().WithPayload(&models.SuccessResponse{Success: isSuccess})
}

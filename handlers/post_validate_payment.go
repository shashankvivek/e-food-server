package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"e-food/constants"
	"e-food/models"
	"e-food/pkg/dao"
	"e-food/restapi/operations/user"
	"encoding/hex"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/razorpay/razorpay-go"
	"strconv"
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
			return user.NewPostValidatePaymentInternalServerError().WithPayload("error fetching order info from 3rd party")
		}
		cartId, err := strconv.ParseInt(body["receipt"].(string), 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return user.NewPostValidatePaymentInternalServerError().WithPayload("error processing order number")
		}
		// TODO: convert cart Entity to order Entity
		// 1. do inventory management
		// 2. capture payment info under transaction entity

		couponId, err := dao.GetAppliedCouponIdOnCart(impl.dbClient, cartId)
		if err != nil {
			fmt.Println(err.Error())
			return user.NewPostValidatePaymentInternalServerError().WithPayload("error getting coupon info")
		}
		// reduce User Limit of coupon
		if couponId != "" {
			err = dao.ReduceUserLimit(impl.dbClient, couponId, 1)
			if err != nil {
				fmt.Println(err.Error())
				return user.NewPostValidatePaymentInternalServerError().WithPayload("error with coupon management")
			}
		}
		// delete cart data from cart table
		err = dao.RenewCart(impl.dbClient, cartId)
		if err != nil {
			fmt.Println(err.Error())
			return user.NewPostValidatePaymentInternalServerError().WithPayload("unknown cart identifier")
		}
	}
	//TODO: return orderId
	return user.NewPostValidatePaymentOK().WithPayload(&models.SuccessResponse{Success: isSuccess, Message: "new_order_id"})
}

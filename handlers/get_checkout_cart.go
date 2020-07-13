package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/user"
	"e-food/model"
	"e-food/pkg/dao"
	"e-food/pkg/integration"
	"e-food/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
)

type cartPreviewImpl struct {
	dbClient            *sql.DB
	couponHandler       dao.CouponHandler
	customerCartHandler dao.CustomerCartHandler
}

func NewCartCheckoutHandler(db *sql.DB, couponHandler dao.CouponHandler, customerCartHandler dao.CustomerCartHandler) user.CheckoutHandler {
	return &cartPreviewImpl{
		dbClient:            db,
		couponHandler:       couponHandler,
		customerCartHandler: customerCartHandler,
	}
}

func (impl *cartPreviewImpl) Handle(params user.CheckoutParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewCheckoutInternalServerError().WithPayload("error in parsing token")
	}
	cartItems, couponId, err := impl.customerCartHandler.GetCustomerCart(impl.dbClient, email.(string))
	if err != nil {
		return user.NewCheckoutInternalServerError().WithPayload("error getting cart details")
	}
	if len(cartItems) == 0 {
		return user.NewCheckoutOK().WithPayload(&models.BillableCart{})
	}
	var couponInfo *model.CouponEntity
	if couponId != "" {
		couponInfo, _ = impl.customerCartHandler.GetCouponDetails(impl.dbClient, couponId, email.(string))

	}
	billedCart, err := integration.PrepareBilling(cartItems, couponInfo)
	if err != nil {
		return user.NewCheckoutInternalServerError().WithPayload("error creating billing")
	}
	return user.NewCheckoutOK().WithPayload(billedCart)
}

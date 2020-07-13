package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/user"
	"e-food/pkg/dao"
	"e-food/pkg/entities"
	"e-food/pkg/integration"
	"e-food/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
)

type cartPreviewImpl struct {
	dbClient *sql.DB
}

func NewCartCheckoutHandler(db *sql.DB) user.CheckoutHandler {
	return &cartPreviewImpl{
		dbClient: db,
	}
}

func (impl *cartPreviewImpl) Handle(params user.CheckoutParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewCheckoutInternalServerError().WithPayload("error in parsing token")
	}
	cartItems, couponId, err := dao.GetCustomerCart(impl.dbClient, email.(string))
	if err != nil {
		return user.NewCheckoutInternalServerError().WithPayload("error getting cart details")
	}
	if len(cartItems) == 0 {
		return user.NewCheckoutOK().WithPayload(&models.BillableCart{})
	}
	var couponInfo *entities.CouponEntity
	if couponId != "" {
		couponInfo, _ = dao.GetCouponDetails(impl.dbClient, couponId, email.(string))

	}
	billedCart, err := integration.PrepareBilling(cartItems, couponInfo)
	if err != nil {
		return user.NewCheckoutInternalServerError().WithPayload("error creating billing")
	}
	return user.NewCheckoutOK().WithPayload(billedCart)
}

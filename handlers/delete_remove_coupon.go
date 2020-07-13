package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/user"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
)

type removeCouponImpl struct {
	dbClient            *sql.DB
	customerCartHandler dao.CustomerCartHandler
}

func NewUserRemoveCouponHandler(db *sql.DB, customerCartHandler dao.CustomerCartHandler) user.RemoveCouponHandler {
	return &removeCouponImpl{
		dbClient:            db,
		customerCartHandler: customerCartHandler,
	}
}

func (impl *removeCouponImpl) Handle(params user.RemoveCouponParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewRemoveCouponInternalServerError().WithPayload("error in parsing token")
	}
	err = impl.customerCartHandler.RemoveCouponFromCart(impl.dbClient, email.(string))
	if err != nil {
		return user.NewRemoveCouponInternalServerError().WithPayload("unable to remove coupon")
	}
	return user.NewRemoveCouponOK().WithPayload(&models.SuccessResponse{Success: true, Message: "coupon removed successfully"})
}

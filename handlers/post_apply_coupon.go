package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/user"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type applyCouponImpl struct {
	dbClient *sql.DB
}

func NewUserApplyCouponHandler(db *sql.DB) user.ApplyCouponHandler {
	return &applyCouponImpl{
		dbClient: db,
	}
}

func (impl *applyCouponImpl) Handle(params user.ApplyCouponParams, principa interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewApplyCouponInternalServerError().WithPayload("error in parsing token")
	}
	err = dao.ApplyCouponToCart(impl.dbClient, params.CouponCode, email.(string))
	if err != nil {
		fmt.Println(err.Error())
		return user.NewApplyCouponInternalServerError().WithPayload("error applying coupon")
	}
	return user.NewApplyCouponOK().WithPayload(&models.SuccessResponse{Success: true, Message: "coupon applied successfully"})
}

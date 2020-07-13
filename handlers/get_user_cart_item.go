package handlers

import (
	"database/sql"
	"e-food/api/restapi/operations/user"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type userCartItemsImpl struct {
	dbClient            *sql.DB
	customerCartHandler dao.CustomerCartHandler
}

func NewUserGetCartHandler(db *sql.DB, customerCartHandler dao.CustomerCartHandler) user.GetCartHandler {
	return &userCartItemsImpl{
		dbClient:            db,
		customerCartHandler: customerCartHandler,
	}
}

func (impl *userCartItemsImpl) Handle(params user.GetCartParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewGetCartInternalServerError().WithPayload("error in parsing token")
	}
	cartItems, _, err := impl.customerCartHandler.GetCustomerCart(impl.dbClient, email.(string))
	if err != nil {
		log.Errorf(err.Error())
		return user.NewGetCartInternalServerError().WithPayload("Error getting info")
	}
	return user.NewGetCartOK().WithPayload(cartItems)
}

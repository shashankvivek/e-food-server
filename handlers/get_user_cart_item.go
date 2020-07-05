package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/restapi/operations/user"
	"e-food/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type userCartItemsImpl struct {
	dbClient *sql.DB
}

func NewUserGetCartHandler(db *sql.DB) user.GetCartHandler {
	return &userCartItemsImpl{
		dbClient: db,
	}
}

func (impl *userCartItemsImpl) Handle(params user.GetCartParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewGetCartInternalServerError().WithPayload("error in parsing token")
	}
	cartItems, err := dao.GetUserCart(impl.dbClient, email.(string))
	if err != nil {
		log.Errorf(err.Error())
		return user.NewGetCartInternalServerError().WithPayload("Error getting info")
	}
	return user.NewGetCartOK().WithPayload(cartItems)
}

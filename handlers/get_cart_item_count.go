package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/cart"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type cartCountImpl struct {
	dbClient *sql.DB
}

func NewGetCartItemCountHandler(db *sql.DB) cart.ItemCountHandler {
	return &cartCountImpl{
		dbClient: db,
	}
}

func (impl *cartCountImpl) Handle(params cart.ItemCountParams) middleware.Responder {
	//TODO: add check for logged in user
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return cart.NewItemCountInternalServerError().WithPayload("fix/enable cookie")
	}
	if cookieInfo.Value == "" {
		return cart.NewItemCountOK().WithPayload(&models.CartItemCount{Count: 0})
	}
	itemCount, err := dao.GetGuestCartItemCount(impl.dbClient, cookieInfo.Value)
	if err != nil {
		return cart.NewItemCountInternalServerError().WithPayload("Error in looking for cart")
	}
	fmt.Println(itemCount)
	return cart.NewItemCountOK().WithPayload(&models.CartItemCount{Count: itemCount})

}

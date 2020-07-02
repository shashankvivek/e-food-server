package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/cart"
	"github.com/go-openapi/runtime/middleware"
)

type cartItemsImpl struct {
	dbClient *sql.DB
}

func NewCartGetItemsHandler(db *sql.DB) cart.GetItemsHandler {
	return &cartItemsImpl{
		dbClient: db,
	}
}

func (impl *cartItemsImpl) Handle(params cart.GetItemsParams) middleware.Responder {
	//TODO: add check for logged in user
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return cart.NewGetItemsInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return cart.NewGetItemsOK().WithPayload(models.CartPreview{})
	}
	items, err := dao.GetGuestCart(impl.dbClient, cookieInfo.Value)
	if err != nil {
		return cart.NewGetItemsInternalServerError().WithPayload("Error in looking for cart")
	}
	return cart.NewGetItemsOK().WithPayload(items)
}

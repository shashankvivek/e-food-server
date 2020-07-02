package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/cart"
	"github.com/go-openapi/runtime/middleware"
)

type addCartItemImpl struct {
	dbClient *sql.DB
}

func NewCartAddItemHandler(dbClient *sql.DB) cart.AddItemHandler {
	return &addCartItemImpl{
		dbClient: dbClient,
	}
}

func (impl *addCartItemImpl) Handle(params cart.AddItemParams) middleware.Responder {
	//TODO: add check for logged in user and add item to cart accordingly
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return cart.NewGetItemsInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return cart.NewGetItemsInternalServerError().WithPayload("Unable to add Item to cart")
	}

	isSuccess, err := dao.AddItemToGuestCart(impl.dbClient, cookieInfo.Value, params.Body.TotalQty, params.Body.ProductID)
	if err != nil {
		return cart.NewGetItemsInternalServerError().WithPayload("Error in adding Item to cart")
	}
	return cart.NewAddItemOK().WithPayload(&models.SuccessResponse{Success: isSuccess, Message: "Item Added"})
}

package handlers

import (
	"database/sql"
	"e-food/constants"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/cart"
	"fmt"
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

	if params.Body.TotalQty < 1 || params.Body.TotalQty > constants.MAX_ALLOWED_CART_ITEM_QTY {
		return cart.NewAddItemOK().WithPayload(&models.SuccessResponse{Success: false, Message: "Quantity must be between 1 and 12"})
	}
	//TODO: check available item in stock before adding and add items with respective msg in case of shortage
	isSuccess, msg, err := dao.AddItemToGuestCart(impl.dbClient, cookieInfo.Value, params.Body.TotalQty, params.Body.ProductID)
	if err != nil {
		fmt.Println(err.Error())
		return cart.NewGetItemsInternalServerError().WithPayload("Error in adding Item to cart")
	}
	if msg == "" {
		msg = "Item Added successfully"
	}
	return cart.NewAddItemOK().WithPayload(&models.SuccessResponse{Success: isSuccess, Message: msg})
}

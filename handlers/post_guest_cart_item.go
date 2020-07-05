package handlers

import (
	"database/sql"
	"e-food/constants"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/guest"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type addCartItemImpl struct {
	dbClient *sql.DB
}

func NewCartAddItemHandler(dbClient *sql.DB) guest.AddItemHandler {
	return &addCartItemImpl{
		dbClient: dbClient,
	}
}

func (impl *addCartItemImpl) Handle(params guest.AddItemParams) middleware.Responder {
	//TODO: add check for logged in user and add item to cart accordingly
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		log.Errorf(err.Error())
		return guest.NewGetItemsInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return guest.NewGetItemsInternalServerError().WithPayload("Unable to add Item to cart")
	}

	if params.Body.TotalQty < 1 || params.Body.TotalQty > constants.MAX_ALLOWED_CART_ITEM_QTY {
		return guest.NewAddItemOK().WithPayload(&models.CartSuccessResponse{Success: false, Message: "Quantity must be between 1 and 12", QtyAdded: 0})
	}
	retVal, err := dao.AddItemToGuestCart(impl.dbClient, cookieInfo.Value, params.Body.TotalQty, params.Body.ProductID)
	if err != nil {
		fmt.Println(err.Error())
		return guest.NewAddItemInternalServerError().WithPayload("Error in adding Item to cart")
	}
	return guest.NewAddItemOK().WithPayload(retVal)
}

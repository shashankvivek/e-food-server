package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/guest"
	"e-food/constants"
	"e-food/pkg/dao"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type addCartItemImpl struct {
	dbClient         *sql.DB
	guestCartHandler dao.GuestCartHandler
	productHandler   dao.ProductHandler
}

func NewGuestCartAddItemHandler(dbClient *sql.DB, gc dao.GuestCartHandler, productHandler dao.ProductHandler) guest.AddItemHandler {
	return &addCartItemImpl{
		dbClient:         dbClient,
		guestCartHandler: gc,
		productHandler:   productHandler,
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

	if *params.Body.TotalQty < 1 || *params.Body.TotalQty > constants.MAX_ALLOWED_CART_ITEM_QTY {
		return guest.NewAddItemOK().WithPayload(&models.CartSuccessResponse{Success: false, Message: "Quantity must be between 1 and 12", QtyAdded: 0})
	}
	retVal, err := impl.guestCartHandler.AddItemToGuestCart(impl.dbClient, impl.productHandler, cookieInfo.Value, *params.Body.TotalQty, *params.Body.ProductID)
	if err != nil {
		fmt.Println(err.Error())
		return guest.NewAddItemInternalServerError().WithPayload("Error in adding Item to cart")
	}
	return guest.NewAddItemOK().WithPayload(retVal)
}

package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/user"
	"e-food/constants"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type postsUserCartItem struct {
	dbClient            *sql.DB
	prodHandler         dao.ProductHandler
	customerCartHandler dao.CustomerCartHandler
}

func NewUserAddToCartHandler(db *sql.DB, prodHandler dao.ProductHandler, customerCartHandler dao.CustomerCartHandler) user.AddToCartHandler {
	return &postsUserCartItem{
		dbClient:            db,
		prodHandler:         prodHandler,
		customerCartHandler: customerCartHandler,
	}
}

func (impl *postsUserCartItem) Handle(params user.AddToCartParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewAddToCartInternalServerError().WithPayload("error in parsing token")
	}
	if *params.Body.TotalQty < 1 || *params.Body.TotalQty > constants.MAX_ALLOWED_CART_ITEM_QTY {
		return user.NewAddToCartOK().WithPayload(&models.CartSuccessResponse{Success: false, Message: "Quantity must be between 1 and 12", QtyAdded: 0})
	}
	retVal, err := impl.customerCartHandler.AddItemToCustomerCart(impl.dbClient, impl.prodHandler, email.(string), *params.Body.TotalQty, *params.Body.ProductID)
	if err != nil {
		fmt.Println(err.Error())
		return user.NewAddToCartInternalServerError().WithPayload("Error in adding Item to cart")
	}
	return user.NewAddToCartOK().WithPayload(retVal)
}

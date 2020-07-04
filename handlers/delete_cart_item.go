package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/cart"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type deleteCartItemImpl struct {
	dbClient *sql.DB
}

func NewCartRemoveItemHandler(db *sql.DB) cart.RemoveItemHandler {
	return &deleteCartItemImpl{
		dbClient: db,
	}
}

func (impl *deleteCartItemImpl) Handle(params cart.RemoveItemParams) middleware.Responder {
	//TODO: add check for logged in user
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return cart.NewRemoveItemInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return cart.NewRemoveItemInternalServerError().WithPayload("error with cookie")
	}
	isDelete, err := dao.RemoveItemFromGuestCart(impl.dbClient, params.ProductID, cookieInfo.Value)
	if err != nil {
		log.Errorf(err.Error())
		return cart.NewRemoveItemInternalServerError().WithPayload("error while deleting item")
	}
	return cart.NewRemoveItemOK().WithPayload(&models.SuccessResponse{Success: isDelete})
}

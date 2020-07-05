package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/guest"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type deleteCartItemImpl struct {
	dbClient *sql.DB
}

func NewCartRemoveItemHandler(db *sql.DB) guest.RemoveItemHandler {
	return &deleteCartItemImpl{
		dbClient: db,
	}
}

func (impl *deleteCartItemImpl) Handle(params guest.RemoveItemParams) middleware.Responder {
	//TODO: add check for logged in user
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return guest.NewRemoveItemInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return guest.NewRemoveItemInternalServerError().WithPayload("error with cookie")
	}
	isDelete, err := dao.RemoveItemFromGuestCart(impl.dbClient, params.ProductID, cookieInfo.Value)
	if err != nil {
		log.Errorf(err.Error())
		return guest.NewRemoveItemInternalServerError().WithPayload("error while deleting item")
	}
	return guest.NewRemoveItemOK().WithPayload(&models.SuccessResponse{Success: isDelete})
}

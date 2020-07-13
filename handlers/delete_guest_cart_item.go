package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/guest"
	"e-food/pkg/dao"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type deleteGuestCartItemImpl struct {
	dbClient         *sql.DB
	guestCartHandler dao.GuestCartHandler
}

func NewGuestCartRemoveItemHandler(db *sql.DB, gc dao.GuestCartHandler) guest.RemoveItemHandler {
	return &deleteGuestCartItemImpl{
		dbClient:         db,
		guestCartHandler: gc,
	}
}

func (impl *deleteGuestCartItemImpl) Handle(params guest.RemoveItemParams) middleware.Responder {
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return guest.NewRemoveItemInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return guest.NewRemoveItemInternalServerError().WithPayload("error with cookie")
	}
	err = impl.guestCartHandler.RemoveItemFromGuestCart(impl.dbClient, params.ProductID, cookieInfo.Value)
	if err != nil {
		log.Errorf(err.Error())
		return guest.NewRemoveItemInternalServerError().WithPayload("error while deleting item")
	}
	return guest.NewRemoveItemOK().WithPayload(&models.SuccessResponse{Success: true})
}

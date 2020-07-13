package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/guest"
	"e-food/pkg/dao"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type guestCartItemsImpl struct {
	dbClient         *sql.DB
	guestCartHandler dao.GuestCartHandler
}

func NewGuestCartGetItemsHandler(db *sql.DB, gc dao.GuestCartHandler) guest.GetItemsHandler {
	return &guestCartItemsImpl{
		dbClient:         db,
		guestCartHandler: gc,
	}
}

func (impl *guestCartItemsImpl) Handle(params guest.GetItemsParams) middleware.Responder {
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return guest.NewGetItemsInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return guest.NewGetItemsOK().WithPayload(models.CartPreview{})
	}
	items, err := impl.guestCartHandler.GetGuestCart(impl.dbClient, cookieInfo.Value)
	if err != nil {
		log.Errorf(err.Error())
		return guest.NewGetItemsInternalServerError().WithPayload("Error in looking for cart")
	}
	return guest.NewGetItemsOK().WithPayload(items)
}

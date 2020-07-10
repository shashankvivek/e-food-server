package handlers

import (
	"database/sql"
	"e-food/models"
	"e-food/pkg/dao"
	"e-food/restapi/operations/guest"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type guestCartItemsImpl struct {
	dbClient *sql.DB
}

func NewCartGetItemsHandler(db *sql.DB) guest.GetItemsHandler {
	return &guestCartItemsImpl{
		dbClient: db,
	}
}

func (impl *guestCartItemsImpl) Handle(params guest.GetItemsParams) middleware.Responder {
	//TODO: add check for logged in user
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return guest.NewGetItemsInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return guest.NewGetItemsOK().WithPayload(models.CartPreview{})
	}
	items, err := dao.GetGuestCart(impl.dbClient, cookieInfo.Value)
	if err != nil {
		log.Errorf(err.Error())
		return guest.NewGetItemsInternalServerError().WithPayload("Error in looking for cart")
	}
	return guest.NewGetItemsOK().WithPayload(items)
}

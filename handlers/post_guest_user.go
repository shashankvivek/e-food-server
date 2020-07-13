package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/guest"
	"e-food/pkg/dao"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
	"strings"
)

type guestUserImpl struct {
	dbClient *sql.DB
}

func NewGuestAddSessionHandler(dbClient *sql.DB) guest.AddSessionHandler {
	return &guestUserImpl{
		dbClient: dbClient,
	}
}

func (impl *guestUserImpl) Handle(params guest.AddSessionParams) middleware.Responder {
	//TODO: add check for logged in user and add item to cart accordingly
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		log.Errorf(err.Error())
		return guest.NewAddSessionInternalServerError().WithPayload("error with cookie")
	}
	if cookieInfo.Value == "" {
		return guest.NewAddSessionInternalServerError().WithPayload("Unable to add Item to cart")
	}
	isSuccess, err := dao.AddGuestSessionDetail(impl.dbClient, cookieInfo.Value, params.SessionInfo.ExtraInfo)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return guest.NewAddSessionOK().WithPayload(&models.SuccessResponse{Success: true, Message: "Session already present"})
		}
		log.Errorf(err.Error())
		return guest.NewAddSessionInternalServerError().WithPayload("Error adding guest session Info")
	}
	if !isSuccess {
		return guest.NewAddSessionInternalServerError().WithPayload("User Session info not added")
	}
	return guest.NewAddSessionOK().WithPayload(&models.SuccessResponse{Success: isSuccess, Message: "Session added"})

}

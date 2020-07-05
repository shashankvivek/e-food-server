package handlers

import (
	"database/sql"
	"e-food/models"
	"e-food/restapi/operations/user"
	"e-food/utils"
	"github.com/go-openapi/runtime/middleware"
)

type loginImpl struct {
	dbClient *sql.DB
}

func NewUserLoginHandler(db *sql.DB) user.LoginHandler {
	return &loginImpl{
		dbClient: db,
	}
}

func (impl *loginImpl) Handle(params user.LoginParams) middleware.Responder {
	//TODO: check username and pwd in DB and then generate token
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("error with cookie")
	}
	email := "test@gmail.com"
	if cookieInfo.Value != "" {
		err := utils.ShiftGuestCartItemsToUserCart(impl.dbClient, cookieInfo.Value, email)
		if err != nil {
			user.NewLoginInternalServerError().WithPayload("Error shifting cart items")
		}
	}
	token, err := utils.GenerateJWT(email, "Shashank", "Vivek")
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}

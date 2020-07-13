package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/user"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

type loginImpl struct {
	dbClient            *sql.DB
	guestCartHandler    dao.GuestCartHandler
	customerInfoHandler dao.CustomerInfoHandler
	prodHandler         dao.ProductHandler
	customerCartHandler dao.CustomerCartHandler
}

func NewUserLoginHandler(db *sql.DB, gc dao.GuestCartHandler, customerInfoHandler dao.CustomerInfoHandler, prodHandler dao.ProductHandler, customerCartHandler dao.CustomerCartHandler) user.LoginHandler {
	return &loginImpl{
		dbClient:            db,
		guestCartHandler:    gc,
		customerInfoHandler: customerInfoHandler,
		prodHandler:         prodHandler,
		customerCartHandler: customerCartHandler,
	}
}

func (impl *loginImpl) Handle(params user.LoginParams) middleware.Responder {
	cookieInfo, err := params.HTTPRequest.Cookie("guest_session")
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("error with cookie")
	}
	email := params.Login.Email
	userInfo, err := impl.customerInfoHandler.FetchUserDetails(impl.dbClient, *email)
	if err != nil {
		fmt.Println(err.Error())
		return user.NewLoginInternalServerError().WithPayload("Error fetching user details")
	}
	err = bcrypt.CompareHashAndPassword([]byte(*userInfo.Password), []byte(*params.Login.Password))
	if err != nil {
		fmt.Println(err)
		return user.NewRegisterNotFound()
	}
	if cookieInfo.Value != "" {
		err := impl.customerCartHandler.ShiftGuestCartItemsToCustomer(impl.dbClient, impl.prodHandler, impl.guestCartHandler, cookieInfo.Value, *email)
		if err != nil {
			user.NewLoginInternalServerError().WithPayload("Error shifting cart items")
		}
	}
	token, err := utils.GenerateJWT(*email, *userInfo.Fname, userInfo.Lname)
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}

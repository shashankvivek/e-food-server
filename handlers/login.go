package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/user"
	"e-food/utils"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
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
	email := params.Login.Email
	userInfo, err := dao.FetchUserDetails(impl.dbClient, email)
	if err != nil {
		fmt.Println(err.Error())
		return user.NewLoginInternalServerError().WithPayload("Error fetching user details")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(params.Login.Password))
	if err != nil {
		fmt.Println(err)
		return user.NewRegisterNotFound()
	}
	if cookieInfo.Value != "" {
		err := utils.ShiftGuestCartItemsToUserCart(impl.dbClient, cookieInfo.Value, email)
		if err != nil {
			user.NewLoginInternalServerError().WithPayload("Error shifting cart items")
		}
	}
	token, err := utils.GenerateJWT(email, userInfo.Fname, userInfo.Lname)
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}

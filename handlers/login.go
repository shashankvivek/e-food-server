package handlers

import (
	"database/sql"
	"e-food/constants"
	"e-food/models"
	"e-food/restapi/operations/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
	"time"
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
	// Once validated, shift any guestCart item to customer cart item
	// delete all entry from guestCart
	token, err := generateJWT("test@gmail.com", "Shashank", "Vivek")
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}

func generateJWT(userEmail, fname, lname string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = userEmail
	claims["fname"] = fname
	claims["lname"] = lname
	claims["exp"] = time.Now().Add(time.Minute * 90).Unix()

	tokenString, err := token.SignedString(constants.MySecretKeyForJWT)
	if err != nil {
		log.Errorf("Error generating Token: " + err.Error())
		return "", err
	}
	return tokenString, nil
}

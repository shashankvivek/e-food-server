package handlers

import (
	"database/sql"
	"e-food/constants"
	"e-food/models"
	"e-food/restapi/operations/user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
	"strings"
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
	token, err := generateJWT()
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "test@gmail.com"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(constants.MySecretKeyForJWT)
	if err != nil {
		log.Errorf("Error generating Token: " + err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return constants.MySecretKeyForJWT, nil
	})
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return token.Valid, nil
}

package handlers

import (
	"database/sql"
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
	token, err := genarateJWT()
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}

func genarateJWT() (string, error) {
	var mySecretKey = []byte("123123123123123") // To be stored in secret.json of Kubernetes
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "test@gmail.com"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		log.Errorf("Error generating Token: " + err.Error())
		return "", err
	}
	return tokenString, nil
}

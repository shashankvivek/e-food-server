package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/user"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type registerImpl struct {
	dbClient *sql.DB
}

func NewUserRegisterHandler(db *sql.DB) user.RegisterHandler {
	return &registerImpl{
		dbClient: db,
	}
}

func (impl *registerImpl) Handle(params user.RegisterParams) middleware.Responder {
	//TODO: check if user already present
	err := dao.RegisterNewUser(impl.dbClient, params.Signup)
	if err != nil {
		fmt.Println(err.Error())
		return user.NewRegisterInternalServerError().WithPayload("Error registering user")
	}
	return user.NewRegisterOK().WithPayload(&models.SuccessResponse{Success: true, Message: "User Registered successfully"})

}

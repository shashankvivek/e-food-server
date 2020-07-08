package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/user"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"strings"
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
	err := dao.RegisterNewUser(impl.dbClient, params.Signup)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "phoneNo_UNIQUE") {
				return user.NewRegisterOK().WithPayload(&models.SuccessResponse{Success: false, Message: "Mobile already registered"})
			}
			if strings.Contains(err.Error(), "email_UNIQUE") {
				return user.NewRegisterOK().WithPayload(&models.SuccessResponse{Success: false, Message: "Email already registered"})
			}
		}
		return user.NewRegisterInternalServerError().WithPayload("Error registering user")
	}
	return user.NewRegisterOK().WithPayload(&models.SuccessResponse{Success: true, Message: "User Registered successfully"})

}

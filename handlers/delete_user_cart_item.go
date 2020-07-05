package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/models"
	"e-food/restapi/operations/user"
	"e-food/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type delUserCartItemImpl struct {
	dbClient *sql.DB
}

func NewUserRemoveFromCartHandler(db *sql.DB) user.RemoveFromCartHandler {
	return &delUserCartItemImpl{
		dbClient: db,
	}
}

func (impl *delUserCartItemImpl) Handle(params user.RemoveFromCartParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewRemoveFromCartInternalServerError().WithPayload("error in parsing token")
	}
	isDelete, err := dao.RemoveItemFromUserCart(impl.dbClient, params.ProductID, email.(string))
	if err != nil {
		log.Errorf(err.Error())
		return user.NewRemoveFromCartInternalServerError().WithPayload("error while deleting item")
	}
	return user.NewRemoveFromCartOK().WithPayload(&models.SuccessResponse{Success: isDelete})
}

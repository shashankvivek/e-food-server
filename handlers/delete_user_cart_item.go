package handlers

import (
	"database/sql"
	"e-food/models"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
	"e-food/restapi/operations/user"
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
	err = dao.RemoveItemFromCustomerCart(impl.dbClient, params.ProductID, email.(string))
	if err != nil {
		log.Errorf(err.Error())
		return user.NewRemoveFromCartInternalServerError().WithPayload("error while deleting item")
	}
	return user.NewRemoveFromCartOK().WithPayload(&models.SuccessResponse{Success: true})
}

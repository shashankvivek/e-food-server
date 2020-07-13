package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/user"
	"e-food/pkg/dao"
	"e-food/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type delUserCartItemImpl struct {
	dbClient            *sql.DB
	customerCartHandler dao.CustomerCartHandler
}

func NewUserRemoveFromCartHandler(db *sql.DB, customerCartHandler dao.CustomerCartHandler) user.RemoveFromCartHandler {
	return &delUserCartItemImpl{
		dbClient:            db,
		customerCartHandler: customerCartHandler,
	}
}

func (impl *delUserCartItemImpl) Handle(params user.RemoveFromCartParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewRemoveFromCartInternalServerError().WithPayload("error in parsing token")
	}
	err = impl.customerCartHandler.RemoveItemFromCustomerCart(impl.dbClient, params.ProductID, email.(string))
	if err != nil {
		log.Errorf(err.Error())
		return user.NewRemoveFromCartInternalServerError().WithPayload("error while deleting item")
	}
	return user.NewRemoveFromCartOK().WithPayload(&models.SuccessResponse{Success: true})
}

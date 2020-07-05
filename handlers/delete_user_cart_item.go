package handlers

import (
	"database/sql"
	"e-food/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
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
	return nil
}

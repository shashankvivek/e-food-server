package handlers

import (
	"database/sql"
	"e-food/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

type removeCouponImpl struct {
	dbClient *sql.DB
}

func NewUserRemoveCouponHandler(db *sql.DB) user.RemoveCouponHandler {
	return &removeCouponImpl{
		dbClient: db,
	}
}

func (impl *removeCouponImpl) Handle(param user.RemoveCouponParams, principal interface{}) middleware.Responder {
	return nil
}

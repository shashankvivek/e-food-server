package handlers

import (
	"database/sql"
	"e-food/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

type cartPreviewImpl struct {
	dbClient *sql.DB
}

func NewCartCheckoutHandler(db *sql.DB) user.CheckoutHandler {
	return &cartPreviewImpl{
		dbClient: db,
	}
}

func (impl *cartPreviewImpl) Handle(params user.CheckoutParams, principal interface{}) middleware.Responder {
	return user.NewCheckoutOK()
}

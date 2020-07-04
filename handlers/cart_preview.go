package handlers

import (
	"database/sql"
	"e-food/restapi/operations/cart"
	"github.com/go-openapi/runtime/middleware"
)

type cartPreviewImpl struct {
	dbClient *sql.DB
}

func NewCartCheckoutHandler(db *sql.DB) cart.CheckoutHandler {
	return &cartPreviewImpl{
		dbClient: db,
	}
}

func (impl *cartPreviewImpl) Handle(params cart.CheckoutParams, principal interface{}) middleware.Responder {
	return cart.NewCheckoutOK()
}

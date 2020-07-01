package handlers

import (
	"database/sql"
	"e-food/restapi/operations/cart"
	"github.com/go-openapi/runtime/middleware"
)

type addCartItemImpl struct {
	dbClient *sql.DB
}

func NewCartAddItemHandler(dbClient *sql.DB) cart.AddItemHandler {
	return &addCartItemImpl{
		dbClient: dbClient,
	}
}

func (impl *addCartItemImpl) Handle(params cart.AddItemParams) middleware.Responder {
	return cart.NewAddItemOK()
}

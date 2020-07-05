package handlers

import (
	"database/sql"
	"e-food/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

type postsUserCartItem struct {
	dbClient *sql.DB
}

func NewUserAddToCartHandler(db *sql.DB) user.AddToCartHandler {
	return &postsUserCartItem{
		dbClient: db,
	}
}

func (impl *postsUserCartItem) Handle(params user.AddToCartParams, principal interface{}) middleware.Responder {
	return nil
}

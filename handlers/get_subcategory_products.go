package handlers

import (
	"database/sql"
	"e-food/restapi/operations/products"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type subCategoryImpl struct {
	dbClient *sql.DB
}

func NewProductsFromSubCategoryHandler(dbClient *sql.DB) products.GetFromSubCategoryHandler {
	return &subCategoryImpl{
		dbClient: dbClient,
	}
}

func (impl *subCategoryImpl) Handle(params products.GetFromSubCategoryParams) middleware.Responder {
	fmt.Print("hihih")
	return products.NewGetFromSubCategoryOK().WithPayload(nil)
}

package handlers

import (
	"database/sql"
	"e-food/api/restapi/operations/products"
	"e-food/pkg/dao"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type subCategoryImpl struct {
	dbClient    *sql.DB
	prodHandler dao.ProductHandler
}

func NewProductsFromSubCategoryHandler(dbClient *sql.DB, prodHandler dao.ProductHandler) products.GetFromSubCategoryHandler {
	return &subCategoryImpl{
		dbClient:    dbClient,
		prodHandler: prodHandler,
	}
}

func (impl *subCategoryImpl) Handle(params products.GetFromSubCategoryParams) middleware.Responder {
	productList, err := impl.prodHandler.GetProductsBySubCategory(impl.dbClient, params.ID)
	if err != nil {
		log.Errorf(err.Error())
		return products.NewGetFromSubCategoryInternalServerError()
	}
	if len(productList) == 0 {
		return products.NewGetFromSubCategoryNotFound()
	}
	return products.NewGetFromSubCategoryOK().WithPayload(productList)
}

package handlers

import (
	"database/sql"
	"e-food/pkg/dao"
	"e-food/restapi/operations/products"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
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
	productList, err := dao.GetProductsBySubCategory(impl.dbClient, params.ID)
	if err != nil {
		log.Errorf(err.Error())
		return products.NewGetFromSubCategoryInternalServerError()
	}
	if len(productList) == 0 {
		return products.NewGetFromSubCategoryNotFound()
	}
	return products.NewGetFromSubCategoryOK().WithPayload(productList)
}

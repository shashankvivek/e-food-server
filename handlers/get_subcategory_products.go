package handlers

import (
	"database/sql"
	"e-food/dao"
	"e-food/restapi/operations/products"
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
	//cookieInfo ,err := params.HTTPRequest.Cookie("guest_session")
	//if err != nil {
	//	return products.NewGetFromSubCategoryInternalServerError().WithPayload("unable to set cookie")
	//}
	productList, err := dao.GetProductsBySubCategory(impl.dbClient, params.ID)
	if err != nil {
		return products.NewGetFromSubCategoryInternalServerError()
	}
	if len(productList) == 0 {
		return products.NewGetFromSubCategoryNotFound()
	}
	return products.NewGetFromSubCategoryOK().WithPayload(productList)
}

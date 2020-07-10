package handlers

import (
	"database/sql"
	"e-food/pkg/dao"
	"e-food/restapi/operations/menu"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type menuImpl struct {
	dbClient *sql.DB
}

func NewMenuCategoryHandler(dbClient *sql.DB) menu.CategoryListHandler {
	return &menuImpl{
		dbClient: dbClient,
	}
}

func (impl *menuImpl) Handle(param menu.CategoryListParams) middleware.Responder {

	retVal, err := dao.GetMenuItems(impl.dbClient)
	if err != nil {
		log.Errorf(err.Error())
		return menu.NewCategoryListInternalServerError().WithPayload("Server ERROR")
	}

	if len(retVal) == 0 {
		return menu.NewCategoryListNotFound()
	}
	return menu.NewCategoryListOK().WithPayload(retVal)
}

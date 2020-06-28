package handlers

import (
	"database/sql"
	"e-food/models"
	"e-food/restapi/operations/menu"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
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
	//retVal, _ := mysql.GetBroadCategoryList(impl.dbClient)
	//fmt.Println(retVal)
	retVal, _ := models.GetMenuItems(impl.dbClient)
	fmt.Println(retVal)

	return menu.NewCategoryListOK().WithPayload(nil)
}

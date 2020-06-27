package handlers

import (
	"e-food/restapi/operations/menu"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type menuImpl struct {
}

func NewMenuCategoryHandler() menu.CategoryListHandler {
	return &menuImpl{}
}

func (impl *menuImpl) Handle(param menu.CategoryListParams) middleware.Responder {
	fmt.Print(param)
	return menu.NewCategoryListOK().WithPayload(nil)
}

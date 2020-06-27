// Code generated by go-swagger; DO NOT EDIT.

package menu

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CategoryListHandlerFunc turns a function with the right signature into a category list handler
type CategoryListHandlerFunc func(CategoryListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CategoryListHandlerFunc) Handle(params CategoryListParams) middleware.Responder {
	return fn(params)
}

// CategoryListHandler interface for that can handle valid category list params
type CategoryListHandler interface {
	Handle(CategoryListParams) middleware.Responder
}

// NewCategoryList creates a new http.Handler for the category list operation
func NewCategoryList(ctx *middleware.Context, handler CategoryListHandler) *CategoryList {
	return &CategoryList{Context: ctx, Handler: handler}
}

/*CategoryList swagger:route GET /get_categories menu categoryList

CategoryList category list API

*/
type CategoryList struct {
	Context *middleware.Context
	Handler CategoryListHandler
}

func (o *CategoryList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCategoryListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

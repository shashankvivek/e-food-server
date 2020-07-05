// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCartHandlerFunc turns a function with the right signature into a get cart handler
type GetCartHandlerFunc func(GetCartParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCartHandlerFunc) Handle(params GetCartParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCartHandler interface for that can handle valid get cart params
type GetCartHandler interface {
	Handle(GetCartParams, interface{}) middleware.Responder
}

// NewGetCart creates a new http.Handler for the get cart operation
func NewGetCart(ctx *middleware.Context, handler GetCartHandler) *GetCart {
	return &GetCart{Context: ctx, Handler: handler}
}

/*GetCart swagger:route GET /user/cart user getCart

Get All cart items

*/
type GetCart struct {
	Context *middleware.Context
	Handler GetCartHandler
}

func (o *GetCart) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCartParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

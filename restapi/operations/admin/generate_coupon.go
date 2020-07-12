// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GenerateCouponHandlerFunc turns a function with the right signature into a generate coupon handler
type GenerateCouponHandlerFunc func(GenerateCouponParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GenerateCouponHandlerFunc) Handle(params GenerateCouponParams) middleware.Responder {
	return fn(params)
}

// GenerateCouponHandler interface for that can handle valid generate coupon params
type GenerateCouponHandler interface {
	Handle(GenerateCouponParams) middleware.Responder
}

// NewGenerateCoupon creates a new http.Handler for the generate coupon operation
func NewGenerateCoupon(ctx *middleware.Context, handler GenerateCouponHandler) *GenerateCoupon {
	return &GenerateCoupon{Context: ctx, Handler: handler}
}

/*GenerateCoupon swagger:route POST /generateCouponCode admin generateCoupon

Generate coupon code

*/
type GenerateCoupon struct {
	Context *middleware.Context
	Handler GenerateCouponHandler
}

func (o *GenerateCoupon) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGenerateCouponParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

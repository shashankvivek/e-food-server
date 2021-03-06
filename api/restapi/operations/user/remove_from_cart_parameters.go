// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewRemoveFromCartParams creates a new RemoveFromCartParams object
// no default values defined in spec.
func NewRemoveFromCartParams() RemoveFromCartParams {

	return RemoveFromCartParams{}
}

// RemoveFromCartParams contains all the bound params for the remove from cart operation
// typically these are obtained from a http.Request
//
// swagger:parameters RemoveFromCart
type RemoveFromCartParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	ProductID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveFromCartParams() beforehand.
func (o *RemoveFromCartParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qProductID, qhkProductID, _ := qs.GetOK("productId")
	if err := o.bindProductID(qProductID, qhkProductID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProductID binds and validates parameter ProductID from query.
func (o *RemoveFromCartParams) bindProductID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("productId", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("productId", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("productId", "query", "int64", raw)
	}
	o.ProductID = value

	return nil
}

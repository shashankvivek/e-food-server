// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ItemInfo item info
//
// swagger:model ItemInfo
type ItemInfo struct {

	// product Id
	// Required: true
	ProductID *int64 `json:"productId"`

	// total qty
	// Required: true
	TotalQty *int64 `json:"totalQty"`
}

// Validate validates this item info
func (m *ItemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProductID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalQty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemInfo) validateProductID(formats strfmt.Registry) error {

	if err := validate.Required("productId", "body", m.ProductID); err != nil {
		return err
	}

	return nil
}

func (m *ItemInfo) validateTotalQty(formats strfmt.Registry) error {

	if err := validate.Required("totalQty", "body", m.TotalQty); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemInfo) UnmarshalBinary(b []byte) error {
	var res ItemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
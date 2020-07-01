// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CartItem cart item
//
// swagger:model CartItem
type CartItem struct {

	// product Id
	ProductID float64 `json:"productId,omitempty"`

	// total qty
	TotalQty float64 `json:"totalQty,omitempty"`

	// total saving
	TotalSaving float64 `json:"totalSaving,omitempty"`
}

// Validate validates this cart item
func (m *CartItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CartItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartItem) UnmarshalBinary(b []byte) error {
	var res CartItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

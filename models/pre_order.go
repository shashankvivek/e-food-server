// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PreOrder pre order
//
// swagger:model PreOrder
type PreOrder struct {

	// amount
	Amount int64 `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`
}

// Validate validates this pre order
func (m *PreOrder) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PreOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreOrder) UnmarshalBinary(b []byte) error {
	var res PreOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

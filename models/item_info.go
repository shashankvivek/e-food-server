// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ItemInfo item info
//
// swagger:model ItemInfo
type ItemInfo struct {

	// product Id
	ProductID float64 `json:"productId,omitempty"`

	// total qty
	TotalQty float64 `json:"totalQty,omitempty"`
}

// Validate validates this item info
func (m *ItemInfo) Validate(formats strfmt.Registry) error {
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

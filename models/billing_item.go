// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BillingItem billing item
//
// swagger:model BillingItem
type BillingItem struct {

	// currency
	Currency string `json:"currency,omitempty"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// product Id
	ProductID int64 `json:"productId,omitempty"`

	// product name
	ProductName string `json:"productName,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// total price
	TotalPrice float64 `json:"totalPrice,omitempty"`

	// unit price
	UnitPrice float64 `json:"unitPrice,omitempty"`
}

// Validate validates this billing item
func (m *BillingItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BillingItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingItem) UnmarshalBinary(b []byte) error {
	var res BillingItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

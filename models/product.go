// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Product product
// swagger:model Product
type Product struct {

	// Broad Category Id
	BcID int64 `json:"bcId,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Discount to be applied on Unit Price
	DiscountPercentage float64 `json:"discountPercentage,omitempty"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// product Id
	ProductID float64 `json:"productId,omitempty"`

	// Sub Category Id
	ScID int64 `json:"scId,omitempty"`

	// sku
	Sku string `json:"sku,omitempty"`

	// unit price
	UnitPrice float64 `json:"unitPrice,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

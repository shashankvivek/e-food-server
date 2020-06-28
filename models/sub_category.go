// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SubCategory sub category
// swagger:model SubCategory
type SubCategory struct {

	// sc Id
	ScID string `json:"scId,omitempty"`

	// sc image Url
	ScImageURL string `json:"scImageUrl,omitempty"`

	// sc is active
	ScIsActive bool `json:"scIsActive,omitempty"`

	// sc name
	ScName string `json:"scName,omitempty"`
}

// Validate validates this sub category
func (m *SubCategory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubCategory) UnmarshalBinary(b []byte) error {
	var res SubCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

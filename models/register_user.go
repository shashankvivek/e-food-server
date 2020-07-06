// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegisterUser register user
//
// swagger:model RegisterUser
type RegisterUser struct {

	// email
	Email string `json:"email,omitempty"`

	// fname
	Fname string `json:"fname,omitempty"`

	// lname
	Lname string `json:"lname,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// phone no
	PhoneNo int64 `json:"phoneNo,omitempty"`
}

// Validate validates this register user
func (m *RegisterUser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterUser) UnmarshalBinary(b []byte) error {
	var res RegisterUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
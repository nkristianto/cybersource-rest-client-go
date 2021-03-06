// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileBean file bean
//
// swagger:model FileBean
type FileBean struct {

	// Unique identifier for each file
	FileID string `json:"fileId,omitempty"`

	// href
	Href string `json:"href,omitempty"`

	// method
	Method string `json:"method,omitempty"`
}

// Validate validates this file bean
func (m *FileBean) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileBean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileBean) UnmarshalBinary(b []byte) error {
	var res FileBean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

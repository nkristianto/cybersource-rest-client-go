// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PIForIILinkLast p i for i i link last
// swagger:model PIForIILinkLast
type PIForIILinkLast struct {

	// A link to the last collection containing the remaining objects.
	Href string `json:"href,omitempty"`
}

// Validate validates this p i for i i link last
func (m *PIForIILinkLast) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PIForIILinkLast) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PIForIILinkLast) UnmarshalBinary(b []byte) error {
	var res PIForIILinkLast
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

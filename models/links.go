// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Links links
// swagger:model Links
type Links struct {

	// documentation
	Documentation []*LinksDocumentationItems0 `json:"documentation"`

	// next
	Next []*LinksNextItems0 `json:"next"`

	// self
	Self *LinksSelf `json:"self,omitempty"`
}

// Validate validates this links
func (m *Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Links) validateDocumentation(formats strfmt.Registry) error {

	if swag.IsZero(m.Documentation) { // not required
		return nil
	}

	for i := 0; i < len(m.Documentation); i++ {
		if swag.IsZero(m.Documentation[i]) { // not required
			continue
		}

		if m.Documentation[i] != nil {
			if err := m.Documentation[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documentation" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Links) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	for i := 0; i < len(m.Next); i++ {
		if swag.IsZero(m.Next[i]) { // not required
			continue
		}

		if m.Next[i] != nil {
			if err := m.Next[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("next" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Links) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Links) UnmarshalBinary(b []byte) error {
	var res Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LinksDocumentationItems0 links documentation items0
// swagger:model LinksDocumentationItems0
type LinksDocumentationItems0 struct {

	// URI of the linked resource.
	Href string `json:"href,omitempty"`

	// HTTP method of the linked resource.
	Method string `json:"method,omitempty"`

	// Label of the linked resource.
	Title string `json:"title,omitempty"`
}

// Validate validates this links documentation items0
func (m *LinksDocumentationItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinksDocumentationItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinksDocumentationItems0) UnmarshalBinary(b []byte) error {
	var res LinksDocumentationItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LinksNextItems0 links next items0
// swagger:model LinksNextItems0
type LinksNextItems0 struct {

	// URI of the linked resource.
	Href string `json:"href,omitempty"`

	// HTTP method of the linked resource.
	Method string `json:"method,omitempty"`

	// Label of the linked resource.
	Title string `json:"title,omitempty"`
}

// Validate validates this links next items0
func (m *LinksNextItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinksNextItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinksNextItems0) UnmarshalBinary(b []byte) error {
	var res LinksNextItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LinksSelf links self
// swagger:model LinksSelf
type LinksSelf struct {

	// URI of the linked resource.
	Href string `json:"href,omitempty"`

	// HTTP method of the linked resource.
	Method string `json:"method,omitempty"`

	// Label of the linked resource.
	Title string `json:"title,omitempty"`
}

// Validate validates this links self
func (m *LinksSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinksSelf) UnmarshalBinary(b []byte) error {
	var res LinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

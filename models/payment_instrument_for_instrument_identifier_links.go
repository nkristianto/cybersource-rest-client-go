// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentInstrumentForInstrumentIdentifierLinks payment instrument for instrument identifier links
//
// swagger:model PaymentInstrumentForInstrumentIdentifierLinks
type PaymentInstrumentForInstrumentIdentifierLinks struct {

	// first
	First *PaymentInstrumentForInstrumentIdentifierLinksFirst `json:"first,omitempty"`

	// last
	Last *PaymentInstrumentForInstrumentIdentifierLinksLast `json:"last,omitempty"`

	// next
	Next *PaymentInstrumentForInstrumentIdentifierLinksNext `json:"next,omitempty"`

	// prev
	Prev *PaymentInstrumentForInstrumentIdentifierLinksPrev `json:"prev,omitempty"`

	// self
	Self *PaymentInstrumentForInstrumentIdentifierLinksSelf `json:"self,omitempty"`
}

// Validate validates this payment instrument for instrument identifier links
func (m *PaymentInstrumentForInstrumentIdentifierLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
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

func (m *PaymentInstrumentForInstrumentIdentifierLinks) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(m.First) { // not required
		return nil
	}

	if m.First != nil {
		if err := m.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentInstrumentForInstrumentIdentifierLinks) validateLast(formats strfmt.Registry) error {

	if swag.IsZero(m.Last) { // not required
		return nil
	}

	if m.Last != nil {
		if err := m.Last.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentInstrumentForInstrumentIdentifierLinks) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentInstrumentForInstrumentIdentifierLinks) validatePrev(formats strfmt.Registry) error {

	if swag.IsZero(m.Prev) { // not required
		return nil
	}

	if m.Prev != nil {
		if err := m.Prev.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentInstrumentForInstrumentIdentifierLinks) validateSelf(formats strfmt.Registry) error {

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
func (m *PaymentInstrumentForInstrumentIdentifierLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinks) UnmarshalBinary(b []byte) error {
	var res PaymentInstrumentForInstrumentIdentifierLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentInstrumentForInstrumentIdentifierLinksFirst payment instrument for instrument identifier links first
//
// swagger:model PaymentInstrumentForInstrumentIdentifierLinksFirst
type PaymentInstrumentForInstrumentIdentifierLinksFirst struct {

	// A link to the collection starting at offset zero for the supplied limit.
	Href string `json:"href,omitempty"`
}

// Validate validates this payment instrument for instrument identifier links first
func (m *PaymentInstrumentForInstrumentIdentifierLinksFirst) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksFirst) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksFirst) UnmarshalBinary(b []byte) error {
	var res PaymentInstrumentForInstrumentIdentifierLinksFirst
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentInstrumentForInstrumentIdentifierLinksLast payment instrument for instrument identifier links last
//
// swagger:model PaymentInstrumentForInstrumentIdentifierLinksLast
type PaymentInstrumentForInstrumentIdentifierLinksLast struct {

	// A link to the last collection containing the remaining objects.
	Href string `json:"href,omitempty"`
}

// Validate validates this payment instrument for instrument identifier links last
func (m *PaymentInstrumentForInstrumentIdentifierLinksLast) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksLast) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksLast) UnmarshalBinary(b []byte) error {
	var res PaymentInstrumentForInstrumentIdentifierLinksLast
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentInstrumentForInstrumentIdentifierLinksNext payment instrument for instrument identifier links next
//
// swagger:model PaymentInstrumentForInstrumentIdentifierLinksNext
type PaymentInstrumentForInstrumentIdentifierLinksNext struct {

	// A link to the next collection starting at the supplied offset plus the supplied limit.
	Href string `json:"href,omitempty"`
}

// Validate validates this payment instrument for instrument identifier links next
func (m *PaymentInstrumentForInstrumentIdentifierLinksNext) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksNext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksNext) UnmarshalBinary(b []byte) error {
	var res PaymentInstrumentForInstrumentIdentifierLinksNext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentInstrumentForInstrumentIdentifierLinksPrev A link to the previous collection starting at the supplied offset minus the supplied limit.
//
// swagger:model PaymentInstrumentForInstrumentIdentifierLinksPrev
type PaymentInstrumentForInstrumentIdentifierLinksPrev struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this payment instrument for instrument identifier links prev
func (m *PaymentInstrumentForInstrumentIdentifierLinksPrev) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksPrev) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksPrev) UnmarshalBinary(b []byte) error {
	var res PaymentInstrumentForInstrumentIdentifierLinksPrev
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentInstrumentForInstrumentIdentifierLinksSelf payment instrument for instrument identifier links self
//
// swagger:model PaymentInstrumentForInstrumentIdentifierLinksSelf
type PaymentInstrumentForInstrumentIdentifierLinksSelf struct {

	// A link to the current requested collection.
	Href string `json:"href,omitempty"`
}

// Validate validates this payment instrument for instrument identifier links self
func (m *PaymentInstrumentForInstrumentIdentifierLinksSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentInstrumentForInstrumentIdentifierLinksSelf) UnmarshalBinary(b []byte) error {
	var res PaymentInstrumentForInstrumentIdentifierLinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

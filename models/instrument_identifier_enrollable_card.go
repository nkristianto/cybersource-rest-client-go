// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstrumentIdentifierEnrollableCard instrument identifier enrollable card
// swagger:model InstrumentIdentifierEnrollableCard
type InstrumentIdentifierEnrollableCard struct {

	// bank account
	BankAccount *InstrumentIdentifierEnrollableCardBankAccount `json:"BankAccount,omitempty"`

	// bill to
	BillTo *InstrumentIdentifierEnrollableCardBillTo `json:"billTo,omitempty"`

	// card
	Card *InstrumentIdentifierEnrollableCardCard `json:"card,omitempty"`

	// Type of Card
	Type string `json:"type,omitempty"`
}

// Validate validates this instrument identifier enrollable card
func (m *InstrumentIdentifierEnrollableCard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstrumentIdentifierEnrollableCard) validateBankAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.BankAccount) { // not required
		return nil
	}

	if m.BankAccount != nil {
		if err := m.BankAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BankAccount")
			}
			return err
		}
	}

	return nil
}

func (m *InstrumentIdentifierEnrollableCard) validateBillTo(formats strfmt.Registry) error {

	if swag.IsZero(m.BillTo) { // not required
		return nil
	}

	if m.BillTo != nil {
		if err := m.BillTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billTo")
			}
			return err
		}
	}

	return nil
}

func (m *InstrumentIdentifierEnrollableCard) validateCard(formats strfmt.Registry) error {

	if swag.IsZero(m.Card) { // not required
		return nil
	}

	if m.Card != nil {
		if err := m.Card.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstrumentIdentifierEnrollableCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstrumentIdentifierEnrollableCard) UnmarshalBinary(b []byte) error {
	var res InstrumentIdentifierEnrollableCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InstrumentIdentifierEnrollableCardBankAccount instrument identifier enrollable card bank account
// swagger:model InstrumentIdentifierEnrollableCardBankAccount
type InstrumentIdentifierEnrollableCardBankAccount struct {

	// Checking account number.
	// Max Length: 19
	// Min Length: 1
	Number string `json:"number,omitempty"`

	// Routing number.
	// Max Length: 9
	// Min Length: 1
	RoutingNumber string `json:"routingNumber,omitempty"`
}

// Validate validates this instrument identifier enrollable card bank account
func (m *InstrumentIdentifierEnrollableCardBankAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstrumentIdentifierEnrollableCardBankAccount) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MinLength("BankAccount"+"."+"number", "body", string(m.Number), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("BankAccount"+"."+"number", "body", string(m.Number), 19); err != nil {
		return err
	}

	return nil
}

func (m *InstrumentIdentifierEnrollableCardBankAccount) validateRoutingNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingNumber) { // not required
		return nil
	}

	if err := validate.MinLength("BankAccount"+"."+"routingNumber", "body", string(m.RoutingNumber), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("BankAccount"+"."+"routingNumber", "body", string(m.RoutingNumber), 9); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstrumentIdentifierEnrollableCardBankAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstrumentIdentifierEnrollableCardBankAccount) UnmarshalBinary(b []byte) error {
	var res InstrumentIdentifierEnrollableCardBankAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InstrumentIdentifierEnrollableCardBillTo instrument identifier enrollable card bill to
// swagger:model InstrumentIdentifierEnrollableCardBillTo
type InstrumentIdentifierEnrollableCardBillTo struct {

	// First address line registered with card.
	Address1 string `json:"address1,omitempty"`

	// Second address line registered with card.
	Address2 string `json:"address2,omitempty"`

	// Administrative area registered with card.
	AdministrativeArea string `json:"administrativeArea,omitempty"`

	// Country registered with card.
	Country string `json:"country,omitempty"`

	// Locality registered with card.
	Locality string `json:"locality,omitempty"`

	// Postal code registered with card.
	PostalCode string `json:"postalCode,omitempty"`
}

// Validate validates this instrument identifier enrollable card bill to
func (m *InstrumentIdentifierEnrollableCardBillTo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstrumentIdentifierEnrollableCardBillTo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstrumentIdentifierEnrollableCardBillTo) UnmarshalBinary(b []byte) error {
	var res InstrumentIdentifierEnrollableCardBillTo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InstrumentIdentifierEnrollableCardCard instrument identifier enrollable card card
// swagger:model InstrumentIdentifierEnrollableCardCard
type InstrumentIdentifierEnrollableCardCard struct {

	// Card expiration month.
	//
	// Format: `MM`.
	// Possible values: `01` through `12`.
	//
	// Max Length: 2
	// Min Length: 2
	ExpirationMonth string `json:"expirationMonth,omitempty"`

	// Card expiration year.
	//
	// Format: `YYYY`.
	// Possible values: `1900` through `2099`.
	//
	// Max Length: 4
	// Min Length: 4
	ExpirationYear string `json:"expirationYear,omitempty"`

	// Credit card number (PAN).
	// Max Length: 19
	// Min Length: 12
	Number string `json:"number,omitempty"`

	// Card security code.
	// Max Length: 4
	// Min Length: 3
	SecurityCode string `json:"securityCode,omitempty"`
}

// Validate validates this instrument identifier enrollable card card
func (m *InstrumentIdentifierEnrollableCardCard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationYear(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstrumentIdentifierEnrollableCardCard) validateExpirationMonth(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationMonth) { // not required
		return nil
	}

	if err := validate.MinLength("card"+"."+"expirationMonth", "body", string(m.ExpirationMonth), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("card"+"."+"expirationMonth", "body", string(m.ExpirationMonth), 2); err != nil {
		return err
	}

	return nil
}

func (m *InstrumentIdentifierEnrollableCardCard) validateExpirationYear(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationYear) { // not required
		return nil
	}

	if err := validate.MinLength("card"+"."+"expirationYear", "body", string(m.ExpirationYear), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("card"+"."+"expirationYear", "body", string(m.ExpirationYear), 4); err != nil {
		return err
	}

	return nil
}

func (m *InstrumentIdentifierEnrollableCardCard) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MinLength("card"+"."+"number", "body", string(m.Number), 12); err != nil {
		return err
	}

	if err := validate.MaxLength("card"+"."+"number", "body", string(m.Number), 19); err != nil {
		return err
	}

	return nil
}

func (m *InstrumentIdentifierEnrollableCardCard) validateSecurityCode(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityCode) { // not required
		return nil
	}

	if err := validate.MinLength("card"+"."+"securityCode", "body", string(m.SecurityCode), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("card"+"."+"securityCode", "body", string(m.SecurityCode), 4); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstrumentIdentifierEnrollableCardCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstrumentIdentifierEnrollableCardCard) UnmarshalBinary(b []byte) error {
	var res InstrumentIdentifierEnrollableCardCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

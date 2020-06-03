// Code generated by go-swagger; DO NOT EDIT.

package instrument_identifier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteInstrumentIdentifierReader is a Reader for the DeleteInstrumentIdentifier structure.
type DeleteInstrumentIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstrumentIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteInstrumentIdentifierNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteInstrumentIdentifierForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInstrumentIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteInstrumentIdentifierConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteInstrumentIdentifierGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 424:
		result := NewDeleteInstrumentIdentifierFailedDependency()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInstrumentIdentifierInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteInstrumentIdentifierNoContent creates a DeleteInstrumentIdentifierNoContent with default headers values
func NewDeleteInstrumentIdentifierNoContent() *DeleteInstrumentIdentifierNoContent {
	return &DeleteInstrumentIdentifierNoContent{}
}

/*DeleteInstrumentIdentifierNoContent handles this case with default header values.

An existing Instrument Identifier associated with the supplied `tokenId` has been deleted.
*/
type DeleteInstrumentIdentifierNoContent struct {
	/*A globally unique ID associated with your request.
	 */
	UniqueTransactionID string
}

func (o *DeleteInstrumentIdentifierNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tms/v1/instrumentidentifiers/{tokenId}][%d] deleteInstrumentIdentifierNoContent ", 204)
}

func (o *DeleteInstrumentIdentifierNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	return nil
}

// NewDeleteInstrumentIdentifierForbidden creates a DeleteInstrumentIdentifierForbidden with default headers values
func NewDeleteInstrumentIdentifierForbidden() *DeleteInstrumentIdentifierForbidden {
	return &DeleteInstrumentIdentifierForbidden{}
}

/*DeleteInstrumentIdentifierForbidden handles this case with default header values.

Forbidden. The profile might not have permission to perform the token operation.
*/
type DeleteInstrumentIdentifierForbidden struct {
	/*A globally unique ID associated with your request.
	 */
	UniqueTransactionID string

	Payload []*DeleteInstrumentIdentifierForbiddenBodyItems0
}

func (o *DeleteInstrumentIdentifierForbidden) Error() string {
	return fmt.Sprintf("[DELETE /tms/v1/instrumentidentifiers/{tokenId}][%d] deleteInstrumentIdentifierForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInstrumentIdentifierForbidden) GetPayload() []*DeleteInstrumentIdentifierForbiddenBodyItems0 {
	return o.Payload
}

func (o *DeleteInstrumentIdentifierForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstrumentIdentifierNotFound creates a DeleteInstrumentIdentifierNotFound with default headers values
func NewDeleteInstrumentIdentifierNotFound() *DeleteInstrumentIdentifierNotFound {
	return &DeleteInstrumentIdentifierNotFound{}
}

/*DeleteInstrumentIdentifierNotFound handles this case with default header values.

Token Not Found. The `tokenid` may not exist or was entered incorrectly.
*/
type DeleteInstrumentIdentifierNotFound struct {
	/*A globally unique ID associated with your request.
	 */
	UniqueTransactionID string

	Payload []*DeleteInstrumentIdentifierNotFoundBodyItems0
}

func (o *DeleteInstrumentIdentifierNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tms/v1/instrumentidentifiers/{tokenId}][%d] deleteInstrumentIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInstrumentIdentifierNotFound) GetPayload() []*DeleteInstrumentIdentifierNotFoundBodyItems0 {
	return o.Payload
}

func (o *DeleteInstrumentIdentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstrumentIdentifierConflict creates a DeleteInstrumentIdentifierConflict with default headers values
func NewDeleteInstrumentIdentifierConflict() *DeleteInstrumentIdentifierConflict {
	return &DeleteInstrumentIdentifierConflict{}
}

/*DeleteInstrumentIdentifierConflict handles this case with default header values.

Conflict. The token is linked to a Payment Instrument.
*/
type DeleteInstrumentIdentifierConflict struct {
	/*A globally unique ID associated with your request.
	 */
	UniqueTransactionID string

	Payload []*DeleteInstrumentIdentifierConflictBodyItems0
}

func (o *DeleteInstrumentIdentifierConflict) Error() string {
	return fmt.Sprintf("[DELETE /tms/v1/instrumentidentifiers/{tokenId}][%d] deleteInstrumentIdentifierConflict  %+v", 409, o.Payload)
}

func (o *DeleteInstrumentIdentifierConflict) GetPayload() []*DeleteInstrumentIdentifierConflictBodyItems0 {
	return o.Payload
}

func (o *DeleteInstrumentIdentifierConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstrumentIdentifierGone creates a DeleteInstrumentIdentifierGone with default headers values
func NewDeleteInstrumentIdentifierGone() *DeleteInstrumentIdentifierGone {
	return &DeleteInstrumentIdentifierGone{}
}

/*DeleteInstrumentIdentifierGone handles this case with default header values.

Token Not Available. The token has been deleted.
*/
type DeleteInstrumentIdentifierGone struct {
	/*A globally unique ID associated with your request.
	 */
	UniqueTransactionID string

	Payload []*DeleteInstrumentIdentifierGoneBodyItems0
}

func (o *DeleteInstrumentIdentifierGone) Error() string {
	return fmt.Sprintf("[DELETE /tms/v1/instrumentidentifiers/{tokenId}][%d] deleteInstrumentIdentifierGone  %+v", 410, o.Payload)
}

func (o *DeleteInstrumentIdentifierGone) GetPayload() []*DeleteInstrumentIdentifierGoneBodyItems0 {
	return o.Payload
}

func (o *DeleteInstrumentIdentifierGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstrumentIdentifierFailedDependency creates a DeleteInstrumentIdentifierFailedDependency with default headers values
func NewDeleteInstrumentIdentifierFailedDependency() *DeleteInstrumentIdentifierFailedDependency {
	return &DeleteInstrumentIdentifierFailedDependency{}
}

/*DeleteInstrumentIdentifierFailedDependency handles this case with default header values.

Failed Dependency: e.g. The profile represented by the profile-id may not exist or the profile-id was entered incorrectly.
*/
type DeleteInstrumentIdentifierFailedDependency struct {
	/*A globally unique id associated with your request.
	 */
	UniqueTransactionID string

	Payload []*DeleteInstrumentIdentifierFailedDependencyBodyItems0
}

func (o *DeleteInstrumentIdentifierFailedDependency) Error() string {
	return fmt.Sprintf("[DELETE /tms/v1/instrumentidentifiers/{tokenId}][%d] deleteInstrumentIdentifierFailedDependency  %+v", 424, o.Payload)
}

func (o *DeleteInstrumentIdentifierFailedDependency) GetPayload() []*DeleteInstrumentIdentifierFailedDependencyBodyItems0 {
	return o.Payload
}

func (o *DeleteInstrumentIdentifierFailedDependency) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstrumentIdentifierInternalServerError creates a DeleteInstrumentIdentifierInternalServerError with default headers values
func NewDeleteInstrumentIdentifierInternalServerError() *DeleteInstrumentIdentifierInternalServerError {
	return &DeleteInstrumentIdentifierInternalServerError{}
}

/*DeleteInstrumentIdentifierInternalServerError handles this case with default header values.

Unexpected error.
*/
type DeleteInstrumentIdentifierInternalServerError struct {
	/*A globally unique id associated with your request.
	 */
	UniqueTransactionID string

	Payload []*DeleteInstrumentIdentifierInternalServerErrorBodyItems0
}

func (o *DeleteInstrumentIdentifierInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /tms/v1/instrumentidentifiers/{tokenId}][%d] deleteInstrumentIdentifierInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteInstrumentIdentifierInternalServerError) GetPayload() []*DeleteInstrumentIdentifierInternalServerErrorBodyItems0 {
	return o.Payload
}

func (o *DeleteInstrumentIdentifierInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteInstrumentIdentifierConflictBodyItems0 delete instrument identifier conflict body items0
swagger:model DeleteInstrumentIdentifierConflictBodyItems0
*/
type DeleteInstrumentIdentifierConflictBodyItems0 struct {

	// details
	Details *DeleteInstrumentIdentifierConflictBodyItems0Details `json:"details,omitempty"`

	// The detailed message related to the type stated above.
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this delete instrument identifier conflict body items0
func (o *DeleteInstrumentIdentifierConflictBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInstrumentIdentifierConflictBodyItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	if o.Details != nil {
		if err := o.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierConflictBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierConflictBodyItems0) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierConflictBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierConflictBodyItems0Details delete instrument identifier conflict body items0 details
swagger:model DeleteInstrumentIdentifierConflictBodyItems0Details
*/
type DeleteInstrumentIdentifierConflictBodyItems0Details struct {

	// The location of the field that threw the error.
	Location string `json:"location,omitempty"`

	// The name of the field that threw the error.
	Name string `json:"name,omitempty"`
}

// Validate validates this delete instrument identifier conflict body items0 details
func (o *DeleteInstrumentIdentifierConflictBodyItems0Details) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierConflictBodyItems0Details) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierConflictBodyItems0Details) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierConflictBodyItems0Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierConflictBodyLinks delete instrument identifier conflict body links
swagger:model DeleteInstrumentIdentifierConflictBodyLinks
*/
type DeleteInstrumentIdentifierConflictBodyLinks struct {

	// payment instruments
	PaymentInstruments *DeleteInstrumentIdentifierConflictBodyLinksPaymentInstruments `json:"paymentInstruments,omitempty"`
}

// Validate validates this delete instrument identifier conflict body links
func (o *DeleteInstrumentIdentifierConflictBodyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePaymentInstruments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInstrumentIdentifierConflictBodyLinks) validatePaymentInstruments(formats strfmt.Registry) error {

	if swag.IsZero(o.PaymentInstruments) { // not required
		return nil
	}

	if o.PaymentInstruments != nil {
		if err := o.PaymentInstruments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteInstrumentIdentifierConflict" + "." + "_links" + "." + "paymentInstruments")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierConflictBodyLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierConflictBodyLinks) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierConflictBodyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierConflictBodyLinksPaymentInstruments delete instrument identifier conflict body links payment instruments
swagger:model DeleteInstrumentIdentifierConflictBodyLinksPaymentInstruments
*/
type DeleteInstrumentIdentifierConflictBodyLinksPaymentInstruments struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this delete instrument identifier conflict body links payment instruments
func (o *DeleteInstrumentIdentifierConflictBodyLinksPaymentInstruments) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierConflictBodyLinksPaymentInstruments) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierConflictBodyLinksPaymentInstruments) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierConflictBodyLinksPaymentInstruments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierFailedDependencyBodyItems0 delete instrument identifier failed dependency body items0
swagger:model DeleteInstrumentIdentifierFailedDependencyBodyItems0
*/
type DeleteInstrumentIdentifierFailedDependencyBodyItems0 struct {

	// details
	Details *DeleteInstrumentIdentifierFailedDependencyBodyItems0Details `json:"details,omitempty"`

	// The detailed message related to the type stated above.
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this delete instrument identifier failed dependency body items0
func (o *DeleteInstrumentIdentifierFailedDependencyBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInstrumentIdentifierFailedDependencyBodyItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	if o.Details != nil {
		if err := o.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierFailedDependencyBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierFailedDependencyBodyItems0) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierFailedDependencyBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierFailedDependencyBodyItems0Details delete instrument identifier failed dependency body items0 details
swagger:model DeleteInstrumentIdentifierFailedDependencyBodyItems0Details
*/
type DeleteInstrumentIdentifierFailedDependencyBodyItems0Details struct {

	// The location of the field that threw the error.
	Location string `json:"location,omitempty"`

	// The name of the field that threw the error.
	Name string `json:"name,omitempty"`
}

// Validate validates this delete instrument identifier failed dependency body items0 details
func (o *DeleteInstrumentIdentifierFailedDependencyBodyItems0Details) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierFailedDependencyBodyItems0Details) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierFailedDependencyBodyItems0Details) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierFailedDependencyBodyItems0Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierForbiddenBodyItems0 delete instrument identifier forbidden body items0
swagger:model DeleteInstrumentIdentifierForbiddenBodyItems0
*/
type DeleteInstrumentIdentifierForbiddenBodyItems0 struct {

	// details
	Details *DeleteInstrumentIdentifierForbiddenBodyItems0Details `json:"details,omitempty"`

	// The detailed message related to the type stated above.
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this delete instrument identifier forbidden body items0
func (o *DeleteInstrumentIdentifierForbiddenBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInstrumentIdentifierForbiddenBodyItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	if o.Details != nil {
		if err := o.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierForbiddenBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierForbiddenBodyItems0) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierForbiddenBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierForbiddenBodyItems0Details delete instrument identifier forbidden body items0 details
swagger:model DeleteInstrumentIdentifierForbiddenBodyItems0Details
*/
type DeleteInstrumentIdentifierForbiddenBodyItems0Details struct {

	// The location of the field that threw the error.
	Location string `json:"location,omitempty"`

	// The name of the field that threw the error.
	Name string `json:"name,omitempty"`
}

// Validate validates this delete instrument identifier forbidden body items0 details
func (o *DeleteInstrumentIdentifierForbiddenBodyItems0Details) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierForbiddenBodyItems0Details) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierForbiddenBodyItems0Details) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierForbiddenBodyItems0Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierGoneBodyItems0 delete instrument identifier gone body items0
swagger:model DeleteInstrumentIdentifierGoneBodyItems0
*/
type DeleteInstrumentIdentifierGoneBodyItems0 struct {

	// details
	Details *DeleteInstrumentIdentifierGoneBodyItems0Details `json:"details,omitempty"`

	// The detailed message related to the type stated above.
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this delete instrument identifier gone body items0
func (o *DeleteInstrumentIdentifierGoneBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInstrumentIdentifierGoneBodyItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	if o.Details != nil {
		if err := o.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierGoneBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierGoneBodyItems0) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierGoneBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierGoneBodyItems0Details delete instrument identifier gone body items0 details
swagger:model DeleteInstrumentIdentifierGoneBodyItems0Details
*/
type DeleteInstrumentIdentifierGoneBodyItems0Details struct {

	// The location of the field that threw the error.
	Location string `json:"location,omitempty"`

	// The name of the field that threw the error.
	Name string `json:"name,omitempty"`
}

// Validate validates this delete instrument identifier gone body items0 details
func (o *DeleteInstrumentIdentifierGoneBodyItems0Details) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierGoneBodyItems0Details) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierGoneBodyItems0Details) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierGoneBodyItems0Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierInternalServerErrorBodyItems0 delete instrument identifier internal server error body items0
swagger:model DeleteInstrumentIdentifierInternalServerErrorBodyItems0
*/
type DeleteInstrumentIdentifierInternalServerErrorBodyItems0 struct {

	// details
	Details *DeleteInstrumentIdentifierInternalServerErrorBodyItems0Details `json:"details,omitempty"`

	// The detailed message related to the type stated above.
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this delete instrument identifier internal server error body items0
func (o *DeleteInstrumentIdentifierInternalServerErrorBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInstrumentIdentifierInternalServerErrorBodyItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	if o.Details != nil {
		if err := o.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierInternalServerErrorBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierInternalServerErrorBodyItems0) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierInternalServerErrorBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierInternalServerErrorBodyItems0Details delete instrument identifier internal server error body items0 details
swagger:model DeleteInstrumentIdentifierInternalServerErrorBodyItems0Details
*/
type DeleteInstrumentIdentifierInternalServerErrorBodyItems0Details struct {

	// The location of the field that threw the error.
	Location string `json:"location,omitempty"`

	// The name of the field that threw the error.
	Name string `json:"name,omitempty"`
}

// Validate validates this delete instrument identifier internal server error body items0 details
func (o *DeleteInstrumentIdentifierInternalServerErrorBodyItems0Details) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierInternalServerErrorBodyItems0Details) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierInternalServerErrorBodyItems0Details) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierInternalServerErrorBodyItems0Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierNotFoundBodyItems0 delete instrument identifier not found body items0
swagger:model DeleteInstrumentIdentifierNotFoundBodyItems0
*/
type DeleteInstrumentIdentifierNotFoundBodyItems0 struct {

	// details
	Details *DeleteInstrumentIdentifierNotFoundBodyItems0Details `json:"details,omitempty"`

	// The detailed message related to the type stated above.
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this delete instrument identifier not found body items0
func (o *DeleteInstrumentIdentifierNotFoundBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInstrumentIdentifierNotFoundBodyItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	if o.Details != nil {
		if err := o.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierNotFoundBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierNotFoundBodyItems0) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierNotFoundBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInstrumentIdentifierNotFoundBodyItems0Details delete instrument identifier not found body items0 details
swagger:model DeleteInstrumentIdentifierNotFoundBodyItems0Details
*/
type DeleteInstrumentIdentifierNotFoundBodyItems0Details struct {

	// The location of the field that threw the error.
	Location string `json:"location,omitempty"`

	// The name of the field that threw the error.
	Name string `json:"name,omitempty"`
}

// Validate validates this delete instrument identifier not found body items0 details
func (o *DeleteInstrumentIdentifierNotFoundBodyItems0Details) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierNotFoundBodyItems0Details) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInstrumentIdentifierNotFoundBodyItems0Details) UnmarshalBinary(b []byte) error {
	var res DeleteInstrumentIdentifierNotFoundBodyItems0Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package purchase_and_refund_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetPurchaseAndRefundDetailsReader is a Reader for the GetPurchaseAndRefundDetails structure.
type GetPurchaseAndRefundDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPurchaseAndRefundDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPurchaseAndRefundDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPurchaseAndRefundDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPurchaseAndRefundDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPurchaseAndRefundDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPurchaseAndRefundDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPurchaseAndRefundDetailsOK creates a GetPurchaseAndRefundDetailsOK with default headers values
func NewGetPurchaseAndRefundDetailsOK() *GetPurchaseAndRefundDetailsOK {
	return &GetPurchaseAndRefundDetailsOK{}
}

/*GetPurchaseAndRefundDetailsOK handles this case with default header values.

Ok
*/
type GetPurchaseAndRefundDetailsOK struct {
	Payload *GetPurchaseAndRefundDetailsOKBody
}

func (o *GetPurchaseAndRefundDetailsOK) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/purchase-refund-details][%d] getPurchaseAndRefundDetailsOK  %+v", 200, o.Payload)
}

func (o *GetPurchaseAndRefundDetailsOK) GetPayload() *GetPurchaseAndRefundDetailsOKBody {
	return o.Payload
}

func (o *GetPurchaseAndRefundDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPurchaseAndRefundDetailsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseAndRefundDetailsBadRequest creates a GetPurchaseAndRefundDetailsBadRequest with default headers values
func NewGetPurchaseAndRefundDetailsBadRequest() *GetPurchaseAndRefundDetailsBadRequest {
	return &GetPurchaseAndRefundDetailsBadRequest{}
}

/*GetPurchaseAndRefundDetailsBadRequest handles this case with default header values.

Invalid request
*/
type GetPurchaseAndRefundDetailsBadRequest struct {
	Payload *GetPurchaseAndRefundDetailsBadRequestBody
}

func (o *GetPurchaseAndRefundDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/purchase-refund-details][%d] getPurchaseAndRefundDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPurchaseAndRefundDetailsBadRequest) GetPayload() *GetPurchaseAndRefundDetailsBadRequestBody {
	return o.Payload
}

func (o *GetPurchaseAndRefundDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPurchaseAndRefundDetailsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseAndRefundDetailsUnauthorized creates a GetPurchaseAndRefundDetailsUnauthorized with default headers values
func NewGetPurchaseAndRefundDetailsUnauthorized() *GetPurchaseAndRefundDetailsUnauthorized {
	return &GetPurchaseAndRefundDetailsUnauthorized{}
}

/*GetPurchaseAndRefundDetailsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPurchaseAndRefundDetailsUnauthorized struct {
	Payload *GetPurchaseAndRefundDetailsUnauthorizedBody
}

func (o *GetPurchaseAndRefundDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/purchase-refund-details][%d] getPurchaseAndRefundDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPurchaseAndRefundDetailsUnauthorized) GetPayload() *GetPurchaseAndRefundDetailsUnauthorizedBody {
	return o.Payload
}

func (o *GetPurchaseAndRefundDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPurchaseAndRefundDetailsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseAndRefundDetailsNotFound creates a GetPurchaseAndRefundDetailsNotFound with default headers values
func NewGetPurchaseAndRefundDetailsNotFound() *GetPurchaseAndRefundDetailsNotFound {
	return &GetPurchaseAndRefundDetailsNotFound{}
}

/*GetPurchaseAndRefundDetailsNotFound handles this case with default header values.

Report not found
*/
type GetPurchaseAndRefundDetailsNotFound struct {
	Payload *GetPurchaseAndRefundDetailsNotFoundBody
}

func (o *GetPurchaseAndRefundDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/purchase-refund-details][%d] getPurchaseAndRefundDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetPurchaseAndRefundDetailsNotFound) GetPayload() *GetPurchaseAndRefundDetailsNotFoundBody {
	return o.Payload
}

func (o *GetPurchaseAndRefundDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPurchaseAndRefundDetailsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseAndRefundDetailsInternalServerError creates a GetPurchaseAndRefundDetailsInternalServerError with default headers values
func NewGetPurchaseAndRefundDetailsInternalServerError() *GetPurchaseAndRefundDetailsInternalServerError {
	return &GetPurchaseAndRefundDetailsInternalServerError{}
}

/*GetPurchaseAndRefundDetailsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPurchaseAndRefundDetailsInternalServerError struct {
	Payload *GetPurchaseAndRefundDetailsInternalServerErrorBody
}

func (o *GetPurchaseAndRefundDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/purchase-refund-details][%d] getPurchaseAndRefundDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPurchaseAndRefundDetailsInternalServerError) GetPayload() *GetPurchaseAndRefundDetailsInternalServerErrorBody {
	return o.Payload
}

func (o *GetPurchaseAndRefundDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPurchaseAndRefundDetailsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AuthorizationsItems0 Authorization Info Values
swagger:model AuthorizationsItems0
*/
type AuthorizationsItems0 struct {

	// Authorization Amount
	Amount string `json:"amount,omitempty"`

	// Authorization Request Id
	AuthorizationRequestID string `json:"authorizationRequestId,omitempty"`

	// Authorization Code
	Code string `json:"code,omitempty"`

	// Valid ISO 4217 ALPHA-3 currency code
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Authorization RCode
	Rcode string `json:"rcode,omitempty"`

	// An unique identification number assigned by CyberSource to identify the submitted request.
	RequestID string `json:"requestId,omitempty"`

	// Authorization Date
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// Authorization Transaction Reference Number
	TransactionReferenceNumber string `json:"transactionReferenceNumber,omitempty"`
}

// Validate validates this authorizations items0
func (o *AuthorizationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AuthorizationsItems0) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(o.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", o.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AuthorizationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AuthorizationsItems0) UnmarshalBinary(b []byte) error {
	var res AuthorizationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 Provides failed validation input field detail
//
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// Field in request that caused an error
	//
	Field string `json:"field,omitempty"`

	// Documented reason code
	//
	Reason string `json:"reason,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*FeeAndFundingDetailsItems0 Fee Funding Section
swagger:model FeeAndFundingDetailsItems0
*/
type FeeAndFundingDetailsItems0 struct {

	// Discount Amount
	DiscountAmount string `json:"discountAmount,omitempty"`

	// Discount Per Item Fee
	DiscountPerItemFee string `json:"discountPerItemFee,omitempty"`

	// Discount Percentage
	DiscountPercentage string `json:"discountPercentage,omitempty"`

	// Dues Assessments
	DuesAssessments string `json:"duesAssessments,omitempty"`

	// Fee Currency
	FeeCurrency string `json:"feeCurrency,omitempty"`

	// Funding Amount
	FundingAmount string `json:"fundingAmount,omitempty"`

	// Funding Currency (ISO 4217)
	FundingCurrency string `json:"fundingCurrency,omitempty"`

	// interchange Description
	InterchangeDescription string `json:"interchangeDescription,omitempty"`

	// interchange Per Item Fee
	InterchangePerItemFee string `json:"interchangePerItemFee,omitempty"`

	// interchange Percentage
	InterchangePercentage string `json:"interchangePercentage,omitempty"`

	// interchange Percentage Amount
	InterchangePercentageAmount string `json:"interchangePercentageAmount,omitempty"`

	// An unique identification number assigned by CyberSource to identify the submitted request.
	// Max Length: 26
	RequestID string `json:"requestId,omitempty"`

	// Total Fee
	TotalFee string `json:"totalFee,omitempty"`
}

// Validate validates this fee and funding details items0
func (o *FeeAndFundingDetailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FeeAndFundingDetailsItems0) validateRequestID(formats strfmt.Registry) error {

	if swag.IsZero(o.RequestID) { // not required
		return nil
	}

	if err := validate.MaxLength("requestId", "body", string(o.RequestID), 26); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FeeAndFundingDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FeeAndFundingDetailsItems0) UnmarshalBinary(b []byte) error {
	var res FeeAndFundingDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPurchaseAndRefundDetailsBadRequestBody reportingV3PurchaseRefundDetailsGet400Response
//
// HTTP status code for client application
swagger:model GetPurchaseAndRefundDetailsBadRequestBody
*/
type GetPurchaseAndRefundDetailsBadRequestBody struct {

	// Error field list
	//
	// Required: true
	Details []*DetailsItems0 `json:"details"`

	// Short descriptive message to the user.
	//
	// Required: true
	Message *string `json:"message"`

	// Documented reason code
	//
	// Required: true
	Reason *string `json:"reason"`

	// Time of request in UTC.
	//
	// Required: true
	// Format: date-time
	SubmitTimeUtc *strfmt.DateTime `json:"submitTimeUtc"`
}

// Validate validates this get purchase and refund details bad request body
func (o *GetPurchaseAndRefundDetailsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmitTimeUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPurchaseAndRefundDetailsBadRequestBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsBadRequest"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsBadRequestBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsBadRequest"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsBadRequestBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsBadRequest"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("getPurchaseAndRefundDetailsBadRequest"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetPurchaseAndRefundDetailsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPurchaseAndRefundDetailsInternalServerErrorBody reportingV3PurchaseRefundDetailsGet500Response
//
// HTTP status code for client application
swagger:model GetPurchaseAndRefundDetailsInternalServerErrorBody
*/
type GetPurchaseAndRefundDetailsInternalServerErrorBody struct {

	// Error field list
	//
	// Required: true
	Details []*DetailsItems0 `json:"details"`

	// Short descriptive message to the user.
	//
	// Required: true
	Message *string `json:"message"`

	// Documented reason code
	//
	// Required: true
	Reason *string `json:"reason"`

	// Time of request in UTC.
	//
	// Required: true
	// Format: date-time
	SubmitTimeUtc *strfmt.DateTime `json:"submitTimeUtc"`
}

// Validate validates this get purchase and refund details internal server error body
func (o *GetPurchaseAndRefundDetailsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmitTimeUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPurchaseAndRefundDetailsInternalServerErrorBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsInternalServerError"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsInternalServerError" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsInternalServerErrorBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsInternalServerError"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsInternalServerErrorBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsInternalServerError"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsInternalServerErrorBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsInternalServerError"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("getPurchaseAndRefundDetailsInternalServerError"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetPurchaseAndRefundDetailsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPurchaseAndRefundDetailsNotFoundBody reportingV3PurchaseRefundDetailsGet404Response
//
// HTTP status code for client application
swagger:model GetPurchaseAndRefundDetailsNotFoundBody
*/
type GetPurchaseAndRefundDetailsNotFoundBody struct {

	// Error field list
	//
	// Required: true
	Details []*DetailsItems0 `json:"details"`

	// Short descriptive message to the user.
	//
	// Required: true
	Message *string `json:"message"`

	// Documented reason code
	//
	// Required: true
	Reason *string `json:"reason"`

	// Time of request in UTC.
	//
	// Required: true
	// Format: date-time
	SubmitTimeUtc *strfmt.DateTime `json:"submitTimeUtc"`
}

// Validate validates this get purchase and refund details not found body
func (o *GetPurchaseAndRefundDetailsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmitTimeUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPurchaseAndRefundDetailsNotFoundBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsNotFound"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsNotFound" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsNotFoundBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsNotFound"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsNotFoundBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsNotFound"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("getPurchaseAndRefundDetailsNotFound"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetPurchaseAndRefundDetailsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPurchaseAndRefundDetailsOKBody reportingV3PurchaseRefundDetailsGet200Response
swagger:model GetPurchaseAndRefundDetailsOKBody
*/
type GetPurchaseAndRefundDetailsOKBody struct {

	// List of Authorization Info values
	Authorizations []*AuthorizationsItems0 `json:"authorizations"`

	// List of Fee Funding Info values
	FeeAndFundingDetails []*FeeAndFundingDetailsItems0 `json:"feeAndFundingDetails"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`

	// List of Other Info values
	Others []*OthersItems0 `json:"others"`

	// page results
	PageResults int64 `json:"pageResults,omitempty"`

	// List of Request Info values
	RequestDetails []*RequestDetailsItems0 `json:"requestDetails"`

	// List of Settlement Status Info values
	SettlementStatuses []*SettlementStatusesItems0 `json:"settlementStatuses"`

	// List of Settlement Info values
	Settlements []*SettlementsItems0 `json:"settlements"`
}

// Validate validates this get purchase and refund details o k body
func (o *GetPurchaseAndRefundDetailsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFeeAndFundingDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOthers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRequestDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSettlementStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSettlements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPurchaseAndRefundDetailsOKBody) validateAuthorizations(formats strfmt.Registry) error {

	if swag.IsZero(o.Authorizations) { // not required
		return nil
	}

	for i := 0; i < len(o.Authorizations); i++ {
		if swag.IsZero(o.Authorizations[i]) { // not required
			continue
		}

		if o.Authorizations[i] != nil {
			if err := o.Authorizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsOK" + "." + "authorizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsOKBody) validateFeeAndFundingDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.FeeAndFundingDetails) { // not required
		return nil
	}

	for i := 0; i < len(o.FeeAndFundingDetails); i++ {
		if swag.IsZero(o.FeeAndFundingDetails[i]) { // not required
			continue
		}

		if o.FeeAndFundingDetails[i] != nil {
			if err := o.FeeAndFundingDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsOK" + "." + "feeAndFundingDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsOKBody) validateOthers(formats strfmt.Registry) error {

	if swag.IsZero(o.Others) { // not required
		return nil
	}

	for i := 0; i < len(o.Others); i++ {
		if swag.IsZero(o.Others[i]) { // not required
			continue
		}

		if o.Others[i] != nil {
			if err := o.Others[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsOK" + "." + "others" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsOKBody) validateRequestDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.RequestDetails) { // not required
		return nil
	}

	for i := 0; i < len(o.RequestDetails); i++ {
		if swag.IsZero(o.RequestDetails[i]) { // not required
			continue
		}

		if o.RequestDetails[i] != nil {
			if err := o.RequestDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsOK" + "." + "requestDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsOKBody) validateSettlementStatuses(formats strfmt.Registry) error {

	if swag.IsZero(o.SettlementStatuses) { // not required
		return nil
	}

	for i := 0; i < len(o.SettlementStatuses); i++ {
		if swag.IsZero(o.SettlementStatuses[i]) { // not required
			continue
		}

		if o.SettlementStatuses[i] != nil {
			if err := o.SettlementStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsOK" + "." + "settlementStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsOKBody) validateSettlements(formats strfmt.Registry) error {

	if swag.IsZero(o.Settlements) { // not required
		return nil
	}

	for i := 0; i < len(o.Settlements); i++ {
		if swag.IsZero(o.Settlements[i]) { // not required
			continue
		}

		if o.Settlements[i] != nil {
			if err := o.Settlements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsOK" + "." + "settlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPurchaseAndRefundDetailsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPurchaseAndRefundDetailsUnauthorizedBody reportingV3PurchaseRefundDetailsGet401Response
//
// HTTP status code for client application
swagger:model GetPurchaseAndRefundDetailsUnauthorizedBody
*/
type GetPurchaseAndRefundDetailsUnauthorizedBody struct {

	// Error field list
	//
	// Required: true
	Details []*DetailsItems0 `json:"details"`

	// Short descriptive message to the user.
	//
	// Required: true
	Message *string `json:"message"`

	// Documented reason code
	//
	// Required: true
	Reason *string `json:"reason"`

	// Time of request in UTC.
	//
	// Required: true
	// Format: date-time
	SubmitTimeUtc *strfmt.DateTime `json:"submitTimeUtc"`
}

// Validate validates this get purchase and refund details unauthorized body
func (o *GetPurchaseAndRefundDetailsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmitTimeUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPurchaseAndRefundDetailsUnauthorizedBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsUnauthorized"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPurchaseAndRefundDetailsUnauthorized" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsUnauthorizedBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsUnauthorized"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsUnauthorizedBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsUnauthorized"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *GetPurchaseAndRefundDetailsUnauthorizedBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("getPurchaseAndRefundDetailsUnauthorized"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("getPurchaseAndRefundDetailsUnauthorized"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPurchaseAndRefundDetailsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetPurchaseAndRefundDetailsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*OthersItems0 Other Merchant Details Values.
swagger:model OthersItems0
*/
type OthersItems0 struct {

	// First Name
	FirstName string `json:"firstName,omitempty"`

	// Last Name
	LastName string `json:"lastName,omitempty"`

	// Merchant Defined Data
	MerchantData1 string `json:"merchantData1,omitempty"`

	// Merchant Defined Data
	MerchantData2 string `json:"merchantData2,omitempty"`

	// Merchant Defined Data
	MerchantData3 string `json:"merchantData3,omitempty"`

	// Merchant Defined Data
	MerchantData4 string `json:"merchantData4,omitempty"`

	// An unique identification number assigned by CyberSource to identify the submitted request.
	// Max Length: 26
	RequestID string `json:"requestId,omitempty"`
}

// Validate validates this others items0
func (o *OthersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OthersItems0) validateRequestID(formats strfmt.Registry) error {

	if swag.IsZero(o.RequestID) { // not required
		return nil
	}

	if err := validate.MaxLength("requestId", "body", string(o.RequestID), 26); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OthersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OthersItems0) UnmarshalBinary(b []byte) error {
	var res OthersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RequestDetailsItems0 Request Info Section
swagger:model RequestDetailsItems0
*/
type RequestDetailsItems0 struct {

	// Cybersource Merchant Id
	CybersourceMerchantID string `json:"cybersourceMerchantId,omitempty"`

	// Group Name
	GroupName string `json:"groupName,omitempty"`

	// Merchant Reference Number
	MerchantReferenceNumber string `json:"merchantReferenceNumber,omitempty"`

	// Cybersource Processor Merchant Id
	ProcessorMerchantID string `json:"processorMerchantId,omitempty"`

	// An unique identification number assigned by CyberSource to identify the submitted request.
	RequestID string `json:"requestId,omitempty"`

	// Transaction Reference Number
	TransactionReferenceNumber string `json:"transactionReferenceNumber,omitempty"`
}

// Validate validates this request details items0
func (o *RequestDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RequestDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RequestDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RequestDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SettlementStatusesItems0 Settlement Status Section Values.
swagger:model SettlementStatusesItems0
*/
type SettlementStatusesItems0 struct {

	// errorText
	ErrorText string `json:"errorText,omitempty"`

	// ReasonCode
	ReasonCode string `json:"reasonCode,omitempty"`

	// An unique identification number assigned by CyberSource to identify the submitted request.
	// Max Length: 26
	RequestID string `json:"requestId,omitempty"`

	// Settlement Date
	// Format: date-time
	SettlementTime strfmt.DateTime `json:"settlementTime,omitempty"`

	// Settlement Status
	Status string `json:"status,omitempty"`
}

// Validate validates this settlement statuses items0
func (o *SettlementStatusesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSettlementTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SettlementStatusesItems0) validateRequestID(formats strfmt.Registry) error {

	if swag.IsZero(o.RequestID) { // not required
		return nil
	}

	if err := validate.MaxLength("requestId", "body", string(o.RequestID), 26); err != nil {
		return err
	}

	return nil
}

func (o *SettlementStatusesItems0) validateSettlementTime(formats strfmt.Registry) error {

	if swag.IsZero(o.SettlementTime) { // not required
		return nil
	}

	if err := validate.FormatOf("settlementTime", "body", "date-time", o.SettlementTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SettlementStatusesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SettlementStatusesItems0) UnmarshalBinary(b []byte) error {
	var res SettlementStatusesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SettlementsItems0 settlements items0
swagger:model SettlementsItems0
*/
type SettlementsItems0 struct {

	// Account Suffix
	AccountSuffix string `json:"accountSuffix,omitempty"`

	// Amount
	Amount string `json:"amount,omitempty"`

	// Card Type
	CardType string `json:"cardType,omitempty"`

	// Valid ISO 4217 ALPHA-3 currency code
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Cybersource Batch Id
	CybersourceBatchID string `json:"cybersourceBatchId,omitempty"`

	// Cybersource Batch Time
	// Format: date-time
	CybersourceBatchTime strfmt.DateTime `json:"cybersourceBatchTime,omitempty"`

	// Debit Network
	DebitNetwork string `json:"debitNetwork,omitempty"`

	// payment method
	PaymentMethod string `json:"paymentMethod,omitempty"`

	// Payment Type
	PaymentType string `json:"paymentType,omitempty"`

	// An unique identification number assigned by CyberSource to identify the submitted request.
	RequestID string `json:"requestId,omitempty"`

	// Submission Date
	// Format: date-time
	SubmissionTime strfmt.DateTime `json:"submissionTime,omitempty"`

	// Transaction Type
	TransactionType string `json:"transactionType,omitempty"`

	// Solution Type (Wallet)
	WalletType string `json:"walletType,omitempty"`
}

// Validate validates this settlements items0
func (o *SettlementsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCybersourceBatchTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmissionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SettlementsItems0) validateCybersourceBatchTime(formats strfmt.Registry) error {

	if swag.IsZero(o.CybersourceBatchTime) { // not required
		return nil
	}

	if err := validate.FormatOf("cybersourceBatchTime", "body", "date-time", o.CybersourceBatchTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *SettlementsItems0) validateSubmissionTime(formats strfmt.Registry) error {

	if swag.IsZero(o.SubmissionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("submissionTime", "body", "date-time", o.SubmissionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SettlementsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SettlementsItems0) UnmarshalBinary(b []byte) error {
	var res SettlementsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

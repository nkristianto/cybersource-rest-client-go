// Code generated by go-swagger; DO NOT EDIT.

package report_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetResourceV2InfoReader is a Reader for the GetResourceV2Info structure.
type GetResourceV2InfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceV2InfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceV2InfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResourceV2InfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceV2InfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResourceV2InfoOK creates a GetResourceV2InfoOK with default headers values
func NewGetResourceV2InfoOK() *GetResourceV2InfoOK {
	return &GetResourceV2InfoOK{}
}

/*GetResourceV2InfoOK handles this case with default header values.

Ok
*/
type GetResourceV2InfoOK struct {
	Payload *GetResourceV2InfoOKBody
}

func (o *GetResourceV2InfoOK) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/report-definitions][%d] getResourceV2InfoOK  %+v", 200, o.Payload)
}

func (o *GetResourceV2InfoOK) GetPayload() *GetResourceV2InfoOKBody {
	return o.Payload
}

func (o *GetResourceV2InfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetResourceV2InfoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceV2InfoBadRequest creates a GetResourceV2InfoBadRequest with default headers values
func NewGetResourceV2InfoBadRequest() *GetResourceV2InfoBadRequest {
	return &GetResourceV2InfoBadRequest{}
}

/*GetResourceV2InfoBadRequest handles this case with default header values.

Invalid request
*/
type GetResourceV2InfoBadRequest struct {
	Payload *GetResourceV2InfoBadRequestBody
}

func (o *GetResourceV2InfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/report-definitions][%d] getResourceV2InfoBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourceV2InfoBadRequest) GetPayload() *GetResourceV2InfoBadRequestBody {
	return o.Payload
}

func (o *GetResourceV2InfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetResourceV2InfoBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceV2InfoNotFound creates a GetResourceV2InfoNotFound with default headers values
func NewGetResourceV2InfoNotFound() *GetResourceV2InfoNotFound {
	return &GetResourceV2InfoNotFound{}
}

/*GetResourceV2InfoNotFound handles this case with default header values.

Report not found
*/
type GetResourceV2InfoNotFound struct {
}

func (o *GetResourceV2InfoNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/report-definitions][%d] getResourceV2InfoNotFound ", 404)
}

func (o *GetResourceV2InfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetResourceV2InfoBadRequestBody reportingV3ReportDefinitionsGet400Response
//
// HTTP status code for client application
swagger:model GetResourceV2InfoBadRequestBody
*/
type GetResourceV2InfoBadRequestBody struct {

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

// Validate validates this get resource v2 info bad request body
func (o *GetResourceV2InfoBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *GetResourceV2InfoBadRequestBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("getResourceV2InfoBadRequest"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getResourceV2InfoBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetResourceV2InfoBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getResourceV2InfoBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetResourceV2InfoBadRequestBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("getResourceV2InfoBadRequest"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *GetResourceV2InfoBadRequestBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("getResourceV2InfoBadRequest"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("getResourceV2InfoBadRequest"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceV2InfoBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceV2InfoBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetResourceV2InfoBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourceV2InfoOKBody reportingV3ReportDefinitionsGet200Response
swagger:model GetResourceV2InfoOKBody
*/
type GetResourceV2InfoOKBody struct {

	// report definitions
	ReportDefinitions []*ReportDefinitionsItems0 `json:"reportDefinitions"`
}

// Validate validates this get resource v2 info o k body
func (o *GetResourceV2InfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReportDefinitions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourceV2InfoOKBody) validateReportDefinitions(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportDefinitions) { // not required
		return nil
	}

	for i := 0; i < len(o.ReportDefinitions); i++ {
		if swag.IsZero(o.ReportDefinitions[i]) { // not required
			continue
		}

		if o.ReportDefinitions[i] != nil {
			if err := o.ReportDefinitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getResourceV2InfoOK" + "." + "reportDefinitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceV2InfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceV2InfoOKBody) UnmarshalBinary(b []byte) error {
	var res GetResourceV2InfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ReportDefinitionsItems0 report definitions items0
swagger:model ReportDefinitionsItems0
*/
type ReportDefinitionsItems0 struct {

	// description
	Description string `json:"description,omitempty"`

	// | Id  |         Definition Class          |
	// | --- | --------------------------------- |
	// | 210 | TransactionRequestClass           |
	// | 211 | PaymentBatchDetailClass           |
	// | 212 | ExceptionDetailClass              |
	// | 213 | ProcessorSettlementDetailClass    |
	// | 214 | ProcessorEventsDetailClass        |
	// | 215 | FundingDetailClass                |
	// | 216 | AgingDetailClass                  |
	// | 217 | ChargebackAndRetrievalDetailClass |
	// | 218 | DepositDetailClass                |
	// | 219 | FeeDetailClass                    |
	// | 220 | InvoiceSummaryClass               |
	// | 221 | PayerAuthDetailClass              |
	// | 222 | ConversionDetailClass             |
	// | 270 | JPTransactionDetailClass          |
	// | 271 | ServiceFeeDetailClass             |
	// | 310 | GatewayTransactionRequestClass    |
	// | 400 | DecisionManagerEventDetailClass   |
	// | 401 | DecisionManagerDetailClass        |
	// | 410 | FeeSummaryClass                   |
	// | 420 | TaxCalculationClass               |
	// | 520 | POSTerminalExceptionClass         |
	// | 620 | SubscriptionDetailClass           |
	//
	ReportDefinitionID int32 `json:"reportDefinitionId,omitempty"`

	// report defintion name
	ReportDefintionName string `json:"reportDefintionName,omitempty"`

	// supported formats
	// Unique: true
	SupportedFormats []string `json:"supportedFormats"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this report definitions items0
func (o *ReportDefinitionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSupportedFormats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReportDefinitionsItems0) validateSupportedFormats(formats strfmt.Registry) error {

	if swag.IsZero(o.SupportedFormats) { // not required
		return nil
	}

	if err := validate.UniqueItems("supportedFormats", "body", o.SupportedFormats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReportDefinitionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReportDefinitionsItems0) UnmarshalBinary(b []byte) error {
	var res ReportDefinitionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

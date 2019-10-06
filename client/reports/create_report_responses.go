// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// CreateReportReader is a Reader for the CreateReport structure.
type CreateReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateReportCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCreateReportNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewCreateReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateReportCreated creates a CreateReportCreated with default headers values
func NewCreateReportCreated() *CreateReportCreated {
	return &CreateReportCreated{}
}

/*CreateReportCreated handles this case with default header values.

Created
*/
type CreateReportCreated struct {
}

func (o *CreateReportCreated) Error() string {
	return fmt.Sprintf("[POST /reporting/v3/reports][%d] createReportCreated ", 201)
}

func (o *CreateReportCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateReportNotModified creates a CreateReportNotModified with default headers values
func NewCreateReportNotModified() *CreateReportNotModified {
	return &CreateReportNotModified{}
}

/*CreateReportNotModified handles this case with default header values.

Not Modified
*/
type CreateReportNotModified struct {
}

func (o *CreateReportNotModified) Error() string {
	return fmt.Sprintf("[POST /reporting/v3/reports][%d] createReportNotModified ", 304)
}

func (o *CreateReportNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateReportBadRequest creates a CreateReportBadRequest with default headers values
func NewCreateReportBadRequest() *CreateReportBadRequest {
	return &CreateReportBadRequest{}
}

/*CreateReportBadRequest handles this case with default header values.

Invalid request
*/
type CreateReportBadRequest struct {
	Payload *CreateReportBadRequestBody
}

func (o *CreateReportBadRequest) Error() string {
	return fmt.Sprintf("[POST /reporting/v3/reports][%d] createReportBadRequest  %+v", 400, o.Payload)
}

func (o *CreateReportBadRequest) GetPayload() *CreateReportBadRequestBody {
	return o.Payload
}

func (o *CreateReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateReportBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateReportBadRequestBody reportingV3ReportsPost400Response
//
// HTTP status code for client application
swagger:model CreateReportBadRequestBody
*/
type CreateReportBadRequestBody struct {

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

// Validate validates this create report bad request body
func (o *CreateReportBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *CreateReportBadRequestBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("createReportBadRequest"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createReportBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateReportBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createReportBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *CreateReportBadRequestBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("createReportBadRequest"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *CreateReportBadRequestBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("createReportBadRequest"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("createReportBadRequest"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateReportBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateReportBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateReportBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateReportBody create report body
swagger:model CreateReportBody
*/
type CreateReportBody struct {

	// Specifies the group name
	// Pattern: [0-9]*
	GroupName string `json:"groupName,omitempty"`

	// Valid CyberSource Organization Id
	// Pattern: [a-zA-Z0-9-_]+
	OrganizationID string `json:"organizationId,omitempty"`

	// report definition name
	// Max Length: 80
	// Min Length: 1
	// Pattern: [a-zA-Z0-9-]+
	ReportDefinitionName string `json:"reportDefinitionName,omitempty"`

	// End time of the report
	// Format: date-time
	ReportEndTime strfmt.DateTime `json:"reportEndTime,omitempty"`

	// List of fields which needs to get included in a report
	ReportFields []string `json:"reportFields"`

	// List of filters to apply
	ReportFilters map[string][]string `json:"reportFilters,omitempty"`

	// 'Format of the report'
	//
	// Valid values:
	// - application/xml
	// - text/csv
	//
	ReportMimeType string `json:"reportMimeType,omitempty"`

	// Name of the report
	// Max Length: 128
	// Min Length: 1
	// Pattern: [a-zA-Z0-9-_ ]+
	ReportName string `json:"reportName,omitempty"`

	// report preferences
	ReportPreferences *CreateReportParamsBodyReportPreferences `json:"reportPreferences,omitempty"`

	// Start time of the report
	// Format: date-time
	ReportStartTime strfmt.DateTime `json:"reportStartTime,omitempty"`

	// Timezone of the report
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this create report body
func (o *CreateReportBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReportDefinitionName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReportEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReportName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReportPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReportStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateReportBody) validateGroupName(formats strfmt.Registry) error {

	if swag.IsZero(o.GroupName) { // not required
		return nil
	}

	if err := validate.Pattern("requestBody"+"."+"groupName", "body", string(o.GroupName), `[0-9]*`); err != nil {
		return err
	}

	return nil
}

func (o *CreateReportBody) validateOrganizationID(formats strfmt.Registry) error {

	if swag.IsZero(o.OrganizationID) { // not required
		return nil
	}

	if err := validate.Pattern("requestBody"+"."+"organizationId", "body", string(o.OrganizationID), `[a-zA-Z0-9-_]+`); err != nil {
		return err
	}

	return nil
}

func (o *CreateReportBody) validateReportDefinitionName(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportDefinitionName) { // not required
		return nil
	}

	if err := validate.MinLength("requestBody"+"."+"reportDefinitionName", "body", string(o.ReportDefinitionName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("requestBody"+"."+"reportDefinitionName", "body", string(o.ReportDefinitionName), 80); err != nil {
		return err
	}

	if err := validate.Pattern("requestBody"+"."+"reportDefinitionName", "body", string(o.ReportDefinitionName), `[a-zA-Z0-9-]+`); err != nil {
		return err
	}

	return nil
}

func (o *CreateReportBody) validateReportEndTime(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("requestBody"+"."+"reportEndTime", "body", "date-time", o.ReportEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateReportBody) validateReportName(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportName) { // not required
		return nil
	}

	if err := validate.MinLength("requestBody"+"."+"reportName", "body", string(o.ReportName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("requestBody"+"."+"reportName", "body", string(o.ReportName), 128); err != nil {
		return err
	}

	if err := validate.Pattern("requestBody"+"."+"reportName", "body", string(o.ReportName), `[a-zA-Z0-9-_ ]+`); err != nil {
		return err
	}

	return nil
}

func (o *CreateReportBody) validateReportPreferences(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportPreferences) { // not required
		return nil
	}

	if o.ReportPreferences != nil {
		if err := o.ReportPreferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestBody" + "." + "reportPreferences")
			}
			return err
		}
	}

	return nil
}

func (o *CreateReportBody) validateReportStartTime(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("requestBody"+"."+"reportStartTime", "body", "date-time", o.ReportStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateReportBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateReportBody) UnmarshalBinary(b []byte) error {
	var res CreateReportBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateReportParamsBodyReportPreferences Report Preferences
swagger:model CreateReportParamsBodyReportPreferences
*/
type CreateReportParamsBodyReportPreferences struct {

	// Specify the field naming convention to be followed in reports (applicable to only csv report formats)
	//
	// Valid values:
	// - SOAPI
	// - SCMP
	//
	FieldNameConvention string `json:"fieldNameConvention,omitempty"`

	// Indicator to determine whether negative sign infront of amount for all refunded transaction
	SignedAmounts bool `json:"signedAmounts,omitempty"`
}

// Validate validates this create report params body report preferences
func (o *CreateReportParamsBodyReportPreferences) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateReportParamsBodyReportPreferences) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateReportParamsBodyReportPreferences) UnmarshalBinary(b []byte) error {
	var res CreateReportParamsBodyReportPreferences
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
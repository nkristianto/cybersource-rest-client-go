// Code generated by go-swagger; DO NOT EDIT.

package conversion_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetConversionDetailParams creates a new GetConversionDetailParams object
// with the default values initialized.
func NewGetConversionDetailParams() *GetConversionDetailParams {
	var ()
	return &GetConversionDetailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversionDetailParamsWithTimeout creates a new GetConversionDetailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversionDetailParamsWithTimeout(timeout time.Duration) *GetConversionDetailParams {
	var ()
	return &GetConversionDetailParams{

		timeout: timeout,
	}
}

// NewGetConversionDetailParamsWithContext creates a new GetConversionDetailParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversionDetailParamsWithContext(ctx context.Context) *GetConversionDetailParams {
	var ()
	return &GetConversionDetailParams{

		Context: ctx,
	}
}

// NewGetConversionDetailParamsWithHTTPClient creates a new GetConversionDetailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversionDetailParamsWithHTTPClient(client *http.Client) *GetConversionDetailParams {
	var ()
	return &GetConversionDetailParams{
		HTTPClient: client,
	}
}

/*GetConversionDetailParams contains all the parameters to send to the API endpoint
for the get conversion detail operation typically these are written to a http.Request
*/
type GetConversionDetailParams struct {

	/*EndTime
	  Valid report End Time in **ISO 8601 format**
	Please refer the following link to know more about ISO 8601 format.[Rfc Date Format](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14)

	**Example date format:**
	  - yyyy-MM-dd'T'HH:mm:ss.SSSZ (e.g. 2018-01-01T00:00:00.000Z)


	*/
	EndTime strfmt.DateTime
	/*OrganizationID
	  Valid Cybersource Organization Id

	*/
	OrganizationID *string
	/*StartTime
	  Valid report Start Time in **ISO 8601 format**
	Please refer the following link to know more about ISO 8601 format.[Rfc Date Format](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14)

	**Example date format:**
	  - yyyy-MM-dd'T'HH:mm:ss.SSSZ (e.g. 2018-01-01T00:00:00.000Z)


	*/
	StartTime strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversion detail params
func (o *GetConversionDetailParams) WithTimeout(timeout time.Duration) *GetConversionDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversion detail params
func (o *GetConversionDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversion detail params
func (o *GetConversionDetailParams) WithContext(ctx context.Context) *GetConversionDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversion detail params
func (o *GetConversionDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversion detail params
func (o *GetConversionDetailParams) WithHTTPClient(client *http.Client) *GetConversionDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversion detail params
func (o *GetConversionDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTime adds the endTime to the get conversion detail params
func (o *GetConversionDetailParams) WithEndTime(endTime strfmt.DateTime) *GetConversionDetailParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get conversion detail params
func (o *GetConversionDetailParams) SetEndTime(endTime strfmt.DateTime) {
	o.EndTime = endTime
}

// WithOrganizationID adds the organizationID to the get conversion detail params
func (o *GetConversionDetailParams) WithOrganizationID(organizationID *string) *GetConversionDetailParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get conversion detail params
func (o *GetConversionDetailParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithStartTime adds the startTime to the get conversion detail params
func (o *GetConversionDetailParams) WithStartTime(startTime strfmt.DateTime) *GetConversionDetailParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get conversion detail params
func (o *GetConversionDetailParams) SetStartTime(startTime strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversionDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endTime
	qrEndTime := o.EndTime
	qEndTime := qrEndTime.String()
	if qEndTime != "" {
		if err := r.SetQueryParam("endTime", qEndTime); err != nil {
			return err
		}
	}

	if o.OrganizationID != nil {

		// query param organizationId
		var qrOrganizationID string
		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {
			if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
				return err
			}
		}

	}

	// query param startTime
	qrStartTime := o.StartTime
	qStartTime := qrStartTime.String()
	if qStartTime != "" {
		if err := r.SetQueryParam("startTime", qStartTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

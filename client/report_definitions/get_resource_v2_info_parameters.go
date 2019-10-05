// Code generated by go-swagger; DO NOT EDIT.

package report_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetResourceV2InfoParams creates a new GetResourceV2InfoParams object
// with the default values initialized.
func NewGetResourceV2InfoParams() *GetResourceV2InfoParams {
	var ()
	return &GetResourceV2InfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceV2InfoParamsWithTimeout creates a new GetResourceV2InfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResourceV2InfoParamsWithTimeout(timeout time.Duration) *GetResourceV2InfoParams {
	var ()
	return &GetResourceV2InfoParams{

		timeout: timeout,
	}
}

// NewGetResourceV2InfoParamsWithContext creates a new GetResourceV2InfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResourceV2InfoParamsWithContext(ctx context.Context) *GetResourceV2InfoParams {
	var ()
	return &GetResourceV2InfoParams{

		Context: ctx,
	}
}

// NewGetResourceV2InfoParamsWithHTTPClient creates a new GetResourceV2InfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResourceV2InfoParamsWithHTTPClient(client *http.Client) *GetResourceV2InfoParams {
	var ()
	return &GetResourceV2InfoParams{
		HTTPClient: client,
	}
}

/*GetResourceV2InfoParams contains all the parameters to send to the API endpoint
for the get resource v2 info operation typically these are written to a http.Request
*/
type GetResourceV2InfoParams struct {

	/*OrganizationID
	  Valid Cybersource Organization Id

	*/
	OrganizationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get resource v2 info params
func (o *GetResourceV2InfoParams) WithTimeout(timeout time.Duration) *GetResourceV2InfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource v2 info params
func (o *GetResourceV2InfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource v2 info params
func (o *GetResourceV2InfoParams) WithContext(ctx context.Context) *GetResourceV2InfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource v2 info params
func (o *GetResourceV2InfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource v2 info params
func (o *GetResourceV2InfoParams) WithHTTPClient(client *http.Client) *GetResourceV2InfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource v2 info params
func (o *GetResourceV2InfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get resource v2 info params
func (o *GetResourceV2InfoParams) WithOrganizationID(organizationID *string) *GetResourceV2InfoParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get resource v2 info params
func (o *GetResourceV2InfoParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceV2InfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package instrument_identifier

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

// NewCreateInstrumentIdentifierParams creates a new CreateInstrumentIdentifierParams object
// with the default values initialized.
func NewCreateInstrumentIdentifierParams() *CreateInstrumentIdentifierParams {
	var ()
	return &CreateInstrumentIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInstrumentIdentifierParamsWithTimeout creates a new CreateInstrumentIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInstrumentIdentifierParamsWithTimeout(timeout time.Duration) *CreateInstrumentIdentifierParams {
	var ()
	return &CreateInstrumentIdentifierParams{

		timeout: timeout,
	}
}

// NewCreateInstrumentIdentifierParamsWithContext creates a new CreateInstrumentIdentifierParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInstrumentIdentifierParamsWithContext(ctx context.Context) *CreateInstrumentIdentifierParams {
	var ()
	return &CreateInstrumentIdentifierParams{

		Context: ctx,
	}
}

// NewCreateInstrumentIdentifierParamsWithHTTPClient creates a new CreateInstrumentIdentifierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInstrumentIdentifierParamsWithHTTPClient(client *http.Client) *CreateInstrumentIdentifierParams {
	var ()
	return &CreateInstrumentIdentifierParams{
		HTTPClient: client,
	}
}

/*CreateInstrumentIdentifierParams contains all the parameters to send to the API endpoint
for the create instrument identifier operation typically these are written to a http.Request
*/
type CreateInstrumentIdentifierParams struct {

	/*CreateInstrumentIdentifierRequest
	  Please specify either a Card, Bank Account or Enrollable Card

	*/
	CreateInstrumentIdentifierRequest CreateInstrumentIdentifierBody
	/*ProfileID
	  The id of a profile containing user specific TMS configuration.

	*/
	ProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) WithTimeout(timeout time.Duration) *CreateInstrumentIdentifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) WithContext(ctx context.Context) *CreateInstrumentIdentifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) WithHTTPClient(client *http.Client) *CreateInstrumentIdentifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateInstrumentIdentifierRequest adds the createInstrumentIdentifierRequest to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) WithCreateInstrumentIdentifierRequest(createInstrumentIdentifierRequest CreateInstrumentIdentifierBody) *CreateInstrumentIdentifierParams {
	o.SetCreateInstrumentIdentifierRequest(createInstrumentIdentifierRequest)
	return o
}

// SetCreateInstrumentIdentifierRequest adds the createInstrumentIdentifierRequest to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) SetCreateInstrumentIdentifierRequest(createInstrumentIdentifierRequest CreateInstrumentIdentifierBody) {
	o.CreateInstrumentIdentifierRequest = createInstrumentIdentifierRequest
}

// WithProfileID adds the profileID to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) WithProfileID(profileID string) *CreateInstrumentIdentifierParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the create instrument identifier params
func (o *CreateInstrumentIdentifierParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInstrumentIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CreateInstrumentIdentifierRequest); err != nil {
		return err
	}

	// header param profile-id
	if err := r.SetHeaderParam("profile-id", o.ProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

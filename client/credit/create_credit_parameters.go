// Code generated by go-swagger; DO NOT EDIT.

package credit

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

// NewCreateCreditParams creates a new CreateCreditParams object
// with the default values initialized.
func NewCreateCreditParams() *CreateCreditParams {
	var ()
	return &CreateCreditParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCreditParamsWithTimeout creates a new CreateCreditParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCreditParamsWithTimeout(timeout time.Duration) *CreateCreditParams {
	var ()
	return &CreateCreditParams{

		timeout: timeout,
	}
}

// NewCreateCreditParamsWithContext creates a new CreateCreditParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCreditParamsWithContext(ctx context.Context) *CreateCreditParams {
	var ()
	return &CreateCreditParams{

		Context: ctx,
	}
}

// NewCreateCreditParamsWithHTTPClient creates a new CreateCreditParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCreditParamsWithHTTPClient(client *http.Client) *CreateCreditParams {
	var ()
	return &CreateCreditParams{
		HTTPClient: client,
	}
}

/*CreateCreditParams contains all the parameters to send to the API endpoint
for the create credit operation typically these are written to a http.Request
*/
type CreateCreditParams struct {

	/*CreateCreditRequest*/
	CreateCreditRequest CreateCreditBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create credit params
func (o *CreateCreditParams) WithTimeout(timeout time.Duration) *CreateCreditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create credit params
func (o *CreateCreditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create credit params
func (o *CreateCreditParams) WithContext(ctx context.Context) *CreateCreditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create credit params
func (o *CreateCreditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create credit params
func (o *CreateCreditParams) WithHTTPClient(client *http.Client) *CreateCreditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create credit params
func (o *CreateCreditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateCreditRequest adds the createCreditRequest to the create credit params
func (o *CreateCreditParams) WithCreateCreditRequest(createCreditRequest CreateCreditBody) *CreateCreditParams {
	o.SetCreateCreditRequest(createCreditRequest)
	return o
}

// SetCreateCreditRequest adds the createCreditRequest to the create credit params
func (o *CreateCreditParams) SetCreateCreditRequest(createCreditRequest CreateCreditBody) {
	o.CreateCreditRequest = createCreditRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCreditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CreateCreditRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

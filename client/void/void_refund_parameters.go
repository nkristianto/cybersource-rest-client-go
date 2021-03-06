// Code generated by go-swagger; DO NOT EDIT.

package void

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

// NewVoidRefundParams creates a new VoidRefundParams object
// with the default values initialized.
func NewVoidRefundParams() *VoidRefundParams {
	var ()
	return &VoidRefundParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoidRefundParamsWithTimeout creates a new VoidRefundParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoidRefundParamsWithTimeout(timeout time.Duration) *VoidRefundParams {
	var ()
	return &VoidRefundParams{

		timeout: timeout,
	}
}

// NewVoidRefundParamsWithContext creates a new VoidRefundParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoidRefundParamsWithContext(ctx context.Context) *VoidRefundParams {
	var ()
	return &VoidRefundParams{

		Context: ctx,
	}
}

// NewVoidRefundParamsWithHTTPClient creates a new VoidRefundParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoidRefundParamsWithHTTPClient(client *http.Client) *VoidRefundParams {
	var ()
	return &VoidRefundParams{
		HTTPClient: client,
	}
}

/*VoidRefundParams contains all the parameters to send to the API endpoint
for the void refund operation typically these are written to a http.Request
*/
type VoidRefundParams struct {

	/*ID
	  The refund ID returned from a previous refund request.

	*/
	ID string
	/*VoidRefundRequest*/
	VoidRefundRequest VoidRefundBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the void refund params
func (o *VoidRefundParams) WithTimeout(timeout time.Duration) *VoidRefundParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the void refund params
func (o *VoidRefundParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the void refund params
func (o *VoidRefundParams) WithContext(ctx context.Context) *VoidRefundParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the void refund params
func (o *VoidRefundParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the void refund params
func (o *VoidRefundParams) WithHTTPClient(client *http.Client) *VoidRefundParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the void refund params
func (o *VoidRefundParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the void refund params
func (o *VoidRefundParams) WithID(id string) *VoidRefundParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the void refund params
func (o *VoidRefundParams) SetID(id string) {
	o.ID = id
}

// WithVoidRefundRequest adds the voidRefundRequest to the void refund params
func (o *VoidRefundParams) WithVoidRefundRequest(voidRefundRequest VoidRefundBody) *VoidRefundParams {
	o.SetVoidRefundRequest(voidRefundRequest)
	return o
}

// SetVoidRefundRequest adds the voidRefundRequest to the void refund params
func (o *VoidRefundParams) SetVoidRefundRequest(voidRefundRequest VoidRefundBody) {
	o.VoidRefundRequest = voidRefundRequest
}

// WriteToRequest writes these params to a swagger request
func (o *VoidRefundParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.VoidRefundRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

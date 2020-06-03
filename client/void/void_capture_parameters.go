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

// NewVoidCaptureParams creates a new VoidCaptureParams object
// with the default values initialized.
func NewVoidCaptureParams() *VoidCaptureParams {
	var ()
	return &VoidCaptureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoidCaptureParamsWithTimeout creates a new VoidCaptureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoidCaptureParamsWithTimeout(timeout time.Duration) *VoidCaptureParams {
	var ()
	return &VoidCaptureParams{

		timeout: timeout,
	}
}

// NewVoidCaptureParamsWithContext creates a new VoidCaptureParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoidCaptureParamsWithContext(ctx context.Context) *VoidCaptureParams {
	var ()
	return &VoidCaptureParams{

		Context: ctx,
	}
}

// NewVoidCaptureParamsWithHTTPClient creates a new VoidCaptureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoidCaptureParamsWithHTTPClient(client *http.Client) *VoidCaptureParams {
	var ()
	return &VoidCaptureParams{
		HTTPClient: client,
	}
}

/*VoidCaptureParams contains all the parameters to send to the API endpoint
for the void capture operation typically these are written to a http.Request
*/
type VoidCaptureParams struct {

	/*ID
	  The capture ID returned from a previous capture request.

	*/
	ID string
	/*VoidCaptureRequest*/
	VoidCaptureRequest VoidCaptureBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the void capture params
func (o *VoidCaptureParams) WithTimeout(timeout time.Duration) *VoidCaptureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the void capture params
func (o *VoidCaptureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the void capture params
func (o *VoidCaptureParams) WithContext(ctx context.Context) *VoidCaptureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the void capture params
func (o *VoidCaptureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the void capture params
func (o *VoidCaptureParams) WithHTTPClient(client *http.Client) *VoidCaptureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the void capture params
func (o *VoidCaptureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the void capture params
func (o *VoidCaptureParams) WithID(id string) *VoidCaptureParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the void capture params
func (o *VoidCaptureParams) SetID(id string) {
	o.ID = id
}

// WithVoidCaptureRequest adds the voidCaptureRequest to the void capture params
func (o *VoidCaptureParams) WithVoidCaptureRequest(voidCaptureRequest VoidCaptureBody) *VoidCaptureParams {
	o.SetVoidCaptureRequest(voidCaptureRequest)
	return o
}

// SetVoidCaptureRequest adds the voidCaptureRequest to the void capture params
func (o *VoidCaptureParams) SetVoidCaptureRequest(voidCaptureRequest VoidCaptureBody) {
	o.VoidCaptureRequest = voidCaptureRequest
}

// WriteToRequest writes these params to a swagger request
func (o *VoidCaptureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.VoidCaptureRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

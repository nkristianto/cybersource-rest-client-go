// Code generated by go-swagger; DO NOT EDIT.

package key_generation

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

// NewGeneratePublicKeyParams creates a new GeneratePublicKeyParams object
// with the default values initialized.
func NewGeneratePublicKeyParams() *GeneratePublicKeyParams {
	var ()
	return &GeneratePublicKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGeneratePublicKeyParamsWithTimeout creates a new GeneratePublicKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGeneratePublicKeyParamsWithTimeout(timeout time.Duration) *GeneratePublicKeyParams {
	var ()
	return &GeneratePublicKeyParams{

		timeout: timeout,
	}
}

// NewGeneratePublicKeyParamsWithContext creates a new GeneratePublicKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGeneratePublicKeyParamsWithContext(ctx context.Context) *GeneratePublicKeyParams {
	var ()
	return &GeneratePublicKeyParams{

		Context: ctx,
	}
}

// NewGeneratePublicKeyParamsWithHTTPClient creates a new GeneratePublicKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGeneratePublicKeyParamsWithHTTPClient(client *http.Client) *GeneratePublicKeyParams {
	var ()
	return &GeneratePublicKeyParams{
		HTTPClient: client,
	}
}

/*GeneratePublicKeyParams contains all the parameters to send to the API endpoint
for the generate public key operation typically these are written to a http.Request
*/
type GeneratePublicKeyParams struct {

	/*GeneratePublicKeyRequest*/
	GeneratePublicKeyRequest GeneratePublicKeyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the generate public key params
func (o *GeneratePublicKeyParams) WithTimeout(timeout time.Duration) *GeneratePublicKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate public key params
func (o *GeneratePublicKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate public key params
func (o *GeneratePublicKeyParams) WithContext(ctx context.Context) *GeneratePublicKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate public key params
func (o *GeneratePublicKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate public key params
func (o *GeneratePublicKeyParams) WithHTTPClient(client *http.Client) *GeneratePublicKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate public key params
func (o *GeneratePublicKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneratePublicKeyRequest adds the generatePublicKeyRequest to the generate public key params
func (o *GeneratePublicKeyParams) WithGeneratePublicKeyRequest(generatePublicKeyRequest GeneratePublicKeyBody) *GeneratePublicKeyParams {
	o.SetGeneratePublicKeyRequest(generatePublicKeyRequest)
	return o
}

// SetGeneratePublicKeyRequest adds the generatePublicKeyRequest to the generate public key params
func (o *GeneratePublicKeyParams) SetGeneratePublicKeyRequest(generatePublicKeyRequest GeneratePublicKeyBody) {
	o.GeneratePublicKeyRequest = generatePublicKeyRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GeneratePublicKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.GeneratePublicKeyRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

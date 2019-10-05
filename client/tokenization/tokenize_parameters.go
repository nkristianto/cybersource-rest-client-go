// Code generated by go-swagger; DO NOT EDIT.

package tokenization

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

// NewTokenizeParams creates a new TokenizeParams object
// with the default values initialized.
func NewTokenizeParams() *TokenizeParams {
	var ()
	return &TokenizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenizeParamsWithTimeout creates a new TokenizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenizeParamsWithTimeout(timeout time.Duration) *TokenizeParams {
	var ()
	return &TokenizeParams{

		timeout: timeout,
	}
}

// NewTokenizeParamsWithContext creates a new TokenizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenizeParamsWithContext(ctx context.Context) *TokenizeParams {
	var ()
	return &TokenizeParams{

		Context: ctx,
	}
}

// NewTokenizeParamsWithHTTPClient creates a new TokenizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenizeParamsWithHTTPClient(client *http.Client) *TokenizeParams {
	var ()
	return &TokenizeParams{
		HTTPClient: client,
	}
}

/*TokenizeParams contains all the parameters to send to the API endpoint
for the tokenize operation typically these are written to a http.Request
*/
type TokenizeParams struct {

	/*TokenizeRequest*/
	TokenizeRequest TokenizeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tokenize params
func (o *TokenizeParams) WithTimeout(timeout time.Duration) *TokenizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tokenize params
func (o *TokenizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tokenize params
func (o *TokenizeParams) WithContext(ctx context.Context) *TokenizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tokenize params
func (o *TokenizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tokenize params
func (o *TokenizeParams) WithHTTPClient(client *http.Client) *TokenizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tokenize params
func (o *TokenizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTokenizeRequest adds the tokenizeRequest to the tokenize params
func (o *TokenizeParams) WithTokenizeRequest(tokenizeRequest TokenizeBody) *TokenizeParams {
	o.SetTokenizeRequest(tokenizeRequest)
	return o
}

// SetTokenizeRequest adds the tokenizeRequest to the tokenize params
func (o *TokenizeParams) SetTokenizeRequest(tokenizeRequest TokenizeBody) {
	o.TokenizeRequest = tokenizeRequest
}

// WriteToRequest writes these params to a swagger request
func (o *TokenizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.TokenizeRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

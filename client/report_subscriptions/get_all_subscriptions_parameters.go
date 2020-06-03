// Code generated by go-swagger; DO NOT EDIT.

package report_subscriptions

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

// NewGetAllSubscriptionsParams creates a new GetAllSubscriptionsParams object
// with the default values initialized.
func NewGetAllSubscriptionsParams() *GetAllSubscriptionsParams {

	return &GetAllSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllSubscriptionsParamsWithTimeout creates a new GetAllSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllSubscriptionsParamsWithTimeout(timeout time.Duration) *GetAllSubscriptionsParams {

	return &GetAllSubscriptionsParams{

		timeout: timeout,
	}
}

// NewGetAllSubscriptionsParamsWithContext creates a new GetAllSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllSubscriptionsParamsWithContext(ctx context.Context) *GetAllSubscriptionsParams {

	return &GetAllSubscriptionsParams{

		Context: ctx,
	}
}

// NewGetAllSubscriptionsParamsWithHTTPClient creates a new GetAllSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllSubscriptionsParamsWithHTTPClient(client *http.Client) *GetAllSubscriptionsParams {

	return &GetAllSubscriptionsParams{
		HTTPClient: client,
	}
}

/*GetAllSubscriptionsParams contains all the parameters to send to the API endpoint
for the get all subscriptions operation typically these are written to a http.Request
*/
type GetAllSubscriptionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithTimeout(timeout time.Duration) *GetAllSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all subscriptions params
func (o *GetAllSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithContext(ctx context.Context) *GetAllSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all subscriptions params
func (o *GetAllSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithHTTPClient(client *http.Client) *GetAllSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all subscriptions params
func (o *GetAllSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

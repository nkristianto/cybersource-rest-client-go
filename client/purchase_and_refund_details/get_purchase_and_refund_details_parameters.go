// Code generated by go-swagger; DO NOT EDIT.

package purchase_and_refund_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPurchaseAndRefundDetailsParams creates a new GetPurchaseAndRefundDetailsParams object
// with the default values initialized.
func NewGetPurchaseAndRefundDetailsParams() *GetPurchaseAndRefundDetailsParams {
	var (
		limitDefault          = int32(2000)
		paymentSubtypeDefault = string("ALL")
		viewByDefault         = string("requestDate")
	)
	return &GetPurchaseAndRefundDetailsParams{
		Limit:          &limitDefault,
		PaymentSubtype: &paymentSubtypeDefault,
		ViewBy:         &viewByDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPurchaseAndRefundDetailsParamsWithTimeout creates a new GetPurchaseAndRefundDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPurchaseAndRefundDetailsParamsWithTimeout(timeout time.Duration) *GetPurchaseAndRefundDetailsParams {
	var (
		limitDefault          = int32(2000)
		paymentSubtypeDefault = string("ALL")
		viewByDefault         = string("requestDate")
	)
	return &GetPurchaseAndRefundDetailsParams{
		Limit:          &limitDefault,
		PaymentSubtype: &paymentSubtypeDefault,
		ViewBy:         &viewByDefault,

		timeout: timeout,
	}
}

// NewGetPurchaseAndRefundDetailsParamsWithContext creates a new GetPurchaseAndRefundDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPurchaseAndRefundDetailsParamsWithContext(ctx context.Context) *GetPurchaseAndRefundDetailsParams {
	var (
		limitDefault          = int32(2000)
		paymentSubtypeDefault = string("ALL")
		viewByDefault         = string("requestDate")
	)
	return &GetPurchaseAndRefundDetailsParams{
		Limit:          &limitDefault,
		PaymentSubtype: &paymentSubtypeDefault,
		ViewBy:         &viewByDefault,

		Context: ctx,
	}
}

// NewGetPurchaseAndRefundDetailsParamsWithHTTPClient creates a new GetPurchaseAndRefundDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPurchaseAndRefundDetailsParamsWithHTTPClient(client *http.Client) *GetPurchaseAndRefundDetailsParams {
	var (
		limitDefault          = int32(2000)
		paymentSubtypeDefault = string("ALL")
		viewByDefault         = string("requestDate")
	)
	return &GetPurchaseAndRefundDetailsParams{
		Limit:          &limitDefault,
		PaymentSubtype: &paymentSubtypeDefault,
		ViewBy:         &viewByDefault,
		HTTPClient:     client,
	}
}

/*GetPurchaseAndRefundDetailsParams contains all the parameters to send to the API endpoint
for the get purchase and refund details operation typically these are written to a http.Request
*/
type GetPurchaseAndRefundDetailsParams struct {

	/*EndTime
	  Valid report End Time in **ISO 8601 format**
	Please refer the following link to know more about ISO 8601 format.[Rfc Date Format](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14)

	**Example date format:**
	  - yyyy-MM-dd'T'HH:mm:ss.SSSZ (e.g. 2018-01-01T00:00:00.000Z)


	*/
	EndTime strfmt.DateTime
	/*GroupName
	  Valid CyberSource Group Name.User can define groups using CBAPI and Group Management Module in EBC2. Groups are collection of organizationIds

	*/
	GroupName *string
	/*Limit
	  Results count per page. Range(1-2000)

	*/
	Limit *int32
	/*Offset
	  Offset of the Purchase and Refund Results.

	*/
	Offset *int32
	/*OrganizationID
	  Valid Cybersource Organization Id

	*/
	OrganizationID *string
	/*PaymentSubtype
	  Payment Subtypes.
	  - **ALL**:  All Payment Subtypes
	  - **VI** :  Visa
	  - **MC** :  Master Card
	  - **AX** :  American Express
	  - **DI** :  Discover
	  - **DP** :  Pinless Debit


	*/
	PaymentSubtype *string
	/*StartTime
	  Valid report Start Time in **ISO 8601 format**
	Please refer the following link to know more about ISO 8601 format.[Rfc Date Format](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14)

	**Example date format:**
	  - yyyy-MM-dd'T'HH:mm:ss.SSSZ (e.g. 2018-01-01T00:00:00.000Z)


	*/
	StartTime strfmt.DateTime
	/*ViewBy
	  View results by Request Date or Submission Date.
	  - **requestDate** : Request Date
	  - **submissionDate**: Submission Date


	*/
	ViewBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithTimeout(timeout time.Duration) *GetPurchaseAndRefundDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithContext(ctx context.Context) *GetPurchaseAndRefundDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithHTTPClient(client *http.Client) *GetPurchaseAndRefundDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTime adds the endTime to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithEndTime(endTime strfmt.DateTime) *GetPurchaseAndRefundDetailsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetEndTime(endTime strfmt.DateTime) {
	o.EndTime = endTime
}

// WithGroupName adds the groupName to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithGroupName(groupName *string) *GetPurchaseAndRefundDetailsParams {
	o.SetGroupName(groupName)
	return o
}

// SetGroupName adds the groupName to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetGroupName(groupName *string) {
	o.GroupName = groupName
}

// WithLimit adds the limit to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithLimit(limit *int32) *GetPurchaseAndRefundDetailsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithOffset(offset *int32) *GetPurchaseAndRefundDetailsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithOrganizationID(organizationID *string) *GetPurchaseAndRefundDetailsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithPaymentSubtype adds the paymentSubtype to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithPaymentSubtype(paymentSubtype *string) *GetPurchaseAndRefundDetailsParams {
	o.SetPaymentSubtype(paymentSubtype)
	return o
}

// SetPaymentSubtype adds the paymentSubtype to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetPaymentSubtype(paymentSubtype *string) {
	o.PaymentSubtype = paymentSubtype
}

// WithStartTime adds the startTime to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithStartTime(startTime strfmt.DateTime) *GetPurchaseAndRefundDetailsParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetStartTime(startTime strfmt.DateTime) {
	o.StartTime = startTime
}

// WithViewBy adds the viewBy to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) WithViewBy(viewBy *string) *GetPurchaseAndRefundDetailsParams {
	o.SetViewBy(viewBy)
	return o
}

// SetViewBy adds the viewBy to the get purchase and refund details params
func (o *GetPurchaseAndRefundDetailsParams) SetViewBy(viewBy *string) {
	o.ViewBy = viewBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetPurchaseAndRefundDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.GroupName != nil {

		// query param groupName
		var qrGroupName string
		if o.GroupName != nil {
			qrGroupName = *o.GroupName
		}
		qGroupName := qrGroupName
		if qGroupName != "" {
			if err := r.SetQueryParam("groupName", qGroupName); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
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

	if o.PaymentSubtype != nil {

		// query param paymentSubtype
		var qrPaymentSubtype string
		if o.PaymentSubtype != nil {
			qrPaymentSubtype = *o.PaymentSubtype
		}
		qPaymentSubtype := qrPaymentSubtype
		if qPaymentSubtype != "" {
			if err := r.SetQueryParam("paymentSubtype", qPaymentSubtype); err != nil {
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

	if o.ViewBy != nil {

		// query param viewBy
		var qrViewBy string
		if o.ViewBy != nil {
			qrViewBy = *o.ViewBy
		}
		qViewBy := qrViewBy
		if qViewBy != "" {
			if err := r.SetQueryParam("viewBy", qViewBy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

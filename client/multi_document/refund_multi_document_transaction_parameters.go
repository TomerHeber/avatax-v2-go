// Code generated by go-swagger; DO NOT EDIT.

package multi_document

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

	"github.com/TomerHeber/avatax-v2-go/models"
)

// NewRefundMultiDocumentTransactionParams creates a new RefundMultiDocumentTransactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRefundMultiDocumentTransactionParams() *RefundMultiDocumentTransactionParams {
	return &RefundMultiDocumentTransactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRefundMultiDocumentTransactionParamsWithTimeout creates a new RefundMultiDocumentTransactionParams object
// with the ability to set a timeout on a request.
func NewRefundMultiDocumentTransactionParamsWithTimeout(timeout time.Duration) *RefundMultiDocumentTransactionParams {
	return &RefundMultiDocumentTransactionParams{
		timeout: timeout,
	}
}

// NewRefundMultiDocumentTransactionParamsWithContext creates a new RefundMultiDocumentTransactionParams object
// with the ability to set a context for a request.
func NewRefundMultiDocumentTransactionParamsWithContext(ctx context.Context) *RefundMultiDocumentTransactionParams {
	return &RefundMultiDocumentTransactionParams{
		Context: ctx,
	}
}

// NewRefundMultiDocumentTransactionParamsWithHTTPClient creates a new RefundMultiDocumentTransactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRefundMultiDocumentTransactionParamsWithHTTPClient(client *http.Client) *RefundMultiDocumentTransactionParams {
	return &RefundMultiDocumentTransactionParams{
		HTTPClient: client,
	}
}

/* RefundMultiDocumentTransactionParams contains all the parameters to send to the API endpoint
   for the refund multi document transaction operation.

   Typically these are written to a http.Request.
*/
type RefundMultiDocumentTransactionParams struct {

	/* DollarInclude.

	   Specifies objects to include in the response after transaction is created
	*/
	DollarInclude *string

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   Information about the refund to create
	*/
	Body *models.RefundTransactionModel

	/* Code.

	   The code of this MultiDocument transaction
	*/
	Code string

	/* Type.

	   The type of this MultiDocument transaction
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the refund multi document transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefundMultiDocumentTransactionParams) WithDefaults() *RefundMultiDocumentTransactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the refund multi document transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefundMultiDocumentTransactionParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := RefundMultiDocumentTransactionParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) WithTimeout(timeout time.Duration) *RefundMultiDocumentTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) WithContext(ctx context.Context) *RefundMultiDocumentTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) WithHTTPClient(client *http.Client) *RefundMultiDocumentTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarInclude adds the dollarInclude to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) WithDollarInclude(dollarInclude *string) *RefundMultiDocumentTransactionParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithXAvalaraClient adds the xAvalaraClient to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) WithXAvalaraClient(xAvalaraClient *string) *RefundMultiDocumentTransactionParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) WithBody(body *models.RefundTransactionModel) *RefundMultiDocumentTransactionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) SetBody(body *models.RefundTransactionModel) {
	o.Body = body
}

// WithCode adds the code to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) WithCode(code string) *RefundMultiDocumentTransactionParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) SetCode(code string) {
	o.Code = code
}

// WithType adds the typeVar to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) WithType(typeVar string) *RefundMultiDocumentTransactionParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the refund multi document transaction params
func (o *RefundMultiDocumentTransactionParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *RefundMultiDocumentTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarInclude != nil {

		// query param $include
		var qrDollarInclude string

		if o.DollarInclude != nil {
			qrDollarInclude = *o.DollarInclude
		}
		qDollarInclude := qrDollarInclude
		if qDollarInclude != "" {

			if err := r.SetQueryParam("$include", qDollarInclude); err != nil {
				return err
			}
		}
	}

	if o.XAvalaraClient != nil {

		// header param X-Avalara-Client
		if err := r.SetHeaderParam("X-Avalara-Client", *o.XAvalaraClient); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param code
	if err := r.SetPathParam("code", o.Code); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

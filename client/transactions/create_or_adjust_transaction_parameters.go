// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// NewCreateOrAdjustTransactionParams creates a new CreateOrAdjustTransactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrAdjustTransactionParams() *CreateOrAdjustTransactionParams {
	return &CreateOrAdjustTransactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrAdjustTransactionParamsWithTimeout creates a new CreateOrAdjustTransactionParams object
// with the ability to set a timeout on a request.
func NewCreateOrAdjustTransactionParamsWithTimeout(timeout time.Duration) *CreateOrAdjustTransactionParams {
	return &CreateOrAdjustTransactionParams{
		timeout: timeout,
	}
}

// NewCreateOrAdjustTransactionParamsWithContext creates a new CreateOrAdjustTransactionParams object
// with the ability to set a context for a request.
func NewCreateOrAdjustTransactionParamsWithContext(ctx context.Context) *CreateOrAdjustTransactionParams {
	return &CreateOrAdjustTransactionParams{
		Context: ctx,
	}
}

// NewCreateOrAdjustTransactionParamsWithHTTPClient creates a new CreateOrAdjustTransactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrAdjustTransactionParamsWithHTTPClient(client *http.Client) *CreateOrAdjustTransactionParams {
	return &CreateOrAdjustTransactionParams{
		HTTPClient: client,
	}
}

/* CreateOrAdjustTransactionParams contains all the parameters to send to the API endpoint
   for the create or adjust transaction operation.

   Typically these are written to a http.Request.
*/
type CreateOrAdjustTransactionParams struct {

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

	   The transaction you wish to create or adjust
	*/
	Body *models.CreateOrAdjustTransactionModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create or adjust transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrAdjustTransactionParams) WithDefaults() *CreateOrAdjustTransactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create or adjust transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrAdjustTransactionParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := CreateOrAdjustTransactionParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) WithTimeout(timeout time.Duration) *CreateOrAdjustTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) WithContext(ctx context.Context) *CreateOrAdjustTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) WithHTTPClient(client *http.Client) *CreateOrAdjustTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarInclude adds the dollarInclude to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) WithDollarInclude(dollarInclude *string) *CreateOrAdjustTransactionParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithXAvalaraClient adds the xAvalaraClient to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) WithXAvalaraClient(xAvalaraClient *string) *CreateOrAdjustTransactionParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) WithBody(body *models.CreateOrAdjustTransactionModel) *CreateOrAdjustTransactionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create or adjust transaction params
func (o *CreateOrAdjustTransactionParams) SetBody(body *models.CreateOrAdjustTransactionModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrAdjustTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
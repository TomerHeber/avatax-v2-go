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
	"github.com/go-openapi/swag"
)

// NewGetMultiDocumentTransactionByIDParams creates a new GetMultiDocumentTransactionByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMultiDocumentTransactionByIDParams() *GetMultiDocumentTransactionByIDParams {
	return &GetMultiDocumentTransactionByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMultiDocumentTransactionByIDParamsWithTimeout creates a new GetMultiDocumentTransactionByIDParams object
// with the ability to set a timeout on a request.
func NewGetMultiDocumentTransactionByIDParamsWithTimeout(timeout time.Duration) *GetMultiDocumentTransactionByIDParams {
	return &GetMultiDocumentTransactionByIDParams{
		timeout: timeout,
	}
}

// NewGetMultiDocumentTransactionByIDParamsWithContext creates a new GetMultiDocumentTransactionByIDParams object
// with the ability to set a context for a request.
func NewGetMultiDocumentTransactionByIDParamsWithContext(ctx context.Context) *GetMultiDocumentTransactionByIDParams {
	return &GetMultiDocumentTransactionByIDParams{
		Context: ctx,
	}
}

// NewGetMultiDocumentTransactionByIDParamsWithHTTPClient creates a new GetMultiDocumentTransactionByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMultiDocumentTransactionByIDParamsWithHTTPClient(client *http.Client) *GetMultiDocumentTransactionByIDParams {
	return &GetMultiDocumentTransactionByIDParams{
		HTTPClient: client,
	}
}

/* GetMultiDocumentTransactionByIDParams contains all the parameters to send to the API endpoint
   for the get multi document transaction by Id operation.

   Typically these are written to a http.Request.
*/
type GetMultiDocumentTransactionByIDParams struct {

	/* DollarInclude.

	   Specifies objects to include in the response after transaction is created
	*/
	DollarInclude *string

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* ID.

	   The unique ID number of the MultiDocument transaction to retrieve

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get multi document transaction by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMultiDocumentTransactionByIDParams) WithDefaults() *GetMultiDocumentTransactionByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get multi document transaction by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMultiDocumentTransactionByIDParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := GetMultiDocumentTransactionByIDParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) WithTimeout(timeout time.Duration) *GetMultiDocumentTransactionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) WithContext(ctx context.Context) *GetMultiDocumentTransactionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) WithHTTPClient(client *http.Client) *GetMultiDocumentTransactionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarInclude adds the dollarInclude to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) WithDollarInclude(dollarInclude *string) *GetMultiDocumentTransactionByIDParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithXAvalaraClient adds the xAvalaraClient to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) WithXAvalaraClient(xAvalaraClient *string) *GetMultiDocumentTransactionByIDParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithID adds the id to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) WithID(id int64) *GetMultiDocumentTransactionByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get multi document transaction by Id params
func (o *GetMultiDocumentTransactionByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMultiDocumentTransactionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

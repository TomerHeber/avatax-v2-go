// Code generated by go-swagger; DO NOT EDIT.

package nexus

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

// NewGetNexusByFormCodeParams creates a new GetNexusByFormCodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNexusByFormCodeParams() *GetNexusByFormCodeParams {
	return &GetNexusByFormCodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNexusByFormCodeParamsWithTimeout creates a new GetNexusByFormCodeParams object
// with the ability to set a timeout on a request.
func NewGetNexusByFormCodeParamsWithTimeout(timeout time.Duration) *GetNexusByFormCodeParams {
	return &GetNexusByFormCodeParams{
		timeout: timeout,
	}
}

// NewGetNexusByFormCodeParamsWithContext creates a new GetNexusByFormCodeParams object
// with the ability to set a context for a request.
func NewGetNexusByFormCodeParamsWithContext(ctx context.Context) *GetNexusByFormCodeParams {
	return &GetNexusByFormCodeParams{
		Context: ctx,
	}
}

// NewGetNexusByFormCodeParamsWithHTTPClient creates a new GetNexusByFormCodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNexusByFormCodeParamsWithHTTPClient(client *http.Client) *GetNexusByFormCodeParams {
	return &GetNexusByFormCodeParams{
		HTTPClient: client,
	}
}

/* GetNexusByFormCodeParams contains all the parameters to send to the API endpoint
   for the get nexus by form code operation.

   Typically these are written to a http.Request.
*/
type GetNexusByFormCodeParams struct {

	// DollarInclude.
	DollarInclude *string

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyID.

	   The ID of the company that owns this nexus object

	   Format: int32
	*/
	CompanyID int32

	/* FormCode.

	   The form code that we are looking up the nexus for
	*/
	FormCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nexus by form code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNexusByFormCodeParams) WithDefaults() *GetNexusByFormCodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nexus by form code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNexusByFormCodeParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := GetNexusByFormCodeParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get nexus by form code params
func (o *GetNexusByFormCodeParams) WithTimeout(timeout time.Duration) *GetNexusByFormCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nexus by form code params
func (o *GetNexusByFormCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nexus by form code params
func (o *GetNexusByFormCodeParams) WithContext(ctx context.Context) *GetNexusByFormCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nexus by form code params
func (o *GetNexusByFormCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nexus by form code params
func (o *GetNexusByFormCodeParams) WithHTTPClient(client *http.Client) *GetNexusByFormCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nexus by form code params
func (o *GetNexusByFormCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarInclude adds the dollarInclude to the get nexus by form code params
func (o *GetNexusByFormCodeParams) WithDollarInclude(dollarInclude *string) *GetNexusByFormCodeParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the get nexus by form code params
func (o *GetNexusByFormCodeParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithXAvalaraClient adds the xAvalaraClient to the get nexus by form code params
func (o *GetNexusByFormCodeParams) WithXAvalaraClient(xAvalaraClient *string) *GetNexusByFormCodeParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the get nexus by form code params
func (o *GetNexusByFormCodeParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the get nexus by form code params
func (o *GetNexusByFormCodeParams) WithCompanyID(companyID int32) *GetNexusByFormCodeParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the get nexus by form code params
func (o *GetNexusByFormCodeParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithFormCode adds the formCode to the get nexus by form code params
func (o *GetNexusByFormCodeParams) WithFormCode(formCode string) *GetNexusByFormCodeParams {
	o.SetFormCode(formCode)
	return o
}

// SetFormCode adds the formCode to the get nexus by form code params
func (o *GetNexusByFormCodeParams) SetFormCode(formCode string) {
	o.FormCode = formCode
}

// WriteToRequest writes these params to a swagger request
func (o *GetNexusByFormCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param companyId
	if err := r.SetPathParam("companyId", swag.FormatInt32(o.CompanyID)); err != nil {
		return err
	}

	// path param formCode
	if err := r.SetPathParam("formCode", o.FormCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package companies

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

// NewGetCompanyParams creates a new GetCompanyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCompanyParams() *GetCompanyParams {
	return &GetCompanyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompanyParamsWithTimeout creates a new GetCompanyParams object
// with the ability to set a timeout on a request.
func NewGetCompanyParamsWithTimeout(timeout time.Duration) *GetCompanyParams {
	return &GetCompanyParams{
		timeout: timeout,
	}
}

// NewGetCompanyParamsWithContext creates a new GetCompanyParams object
// with the ability to set a context for a request.
func NewGetCompanyParamsWithContext(ctx context.Context) *GetCompanyParams {
	return &GetCompanyParams{
		Context: ctx,
	}
}

// NewGetCompanyParamsWithHTTPClient creates a new GetCompanyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCompanyParamsWithHTTPClient(client *http.Client) *GetCompanyParams {
	return &GetCompanyParams{
		HTTPClient: client,
	}
}

/* GetCompanyParams contains all the parameters to send to the API endpoint
   for the get company operation.

   Typically these are written to a http.Request.
*/
type GetCompanyParams struct {

	/* DollarInclude.

	   OPTIONAL: A comma separated list of special fetch options.

	            * Child objects - Specify one or more of the following to retrieve objects related to each company: "Contacts", "FilingCalendars", "Items", "Locations", "Nexus", "TaxCodes", "NonReportingChildren" or "TaxRules".
	            * Deleted objects - Specify "FetchDeleted" to retrieve information about previously deleted objects.
	*/
	DollarInclude *string

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* ID.

	   The ID of the company to retrieve.

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get company params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompanyParams) WithDefaults() *GetCompanyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get company params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompanyParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := GetCompanyParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get company params
func (o *GetCompanyParams) WithTimeout(timeout time.Duration) *GetCompanyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get company params
func (o *GetCompanyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get company params
func (o *GetCompanyParams) WithContext(ctx context.Context) *GetCompanyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get company params
func (o *GetCompanyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get company params
func (o *GetCompanyParams) WithHTTPClient(client *http.Client) *GetCompanyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get company params
func (o *GetCompanyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarInclude adds the dollarInclude to the get company params
func (o *GetCompanyParams) WithDollarInclude(dollarInclude *string) *GetCompanyParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the get company params
func (o *GetCompanyParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithXAvalaraClient adds the xAvalaraClient to the get company params
func (o *GetCompanyParams) WithXAvalaraClient(xAvalaraClient *string) *GetCompanyParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the get company params
func (o *GetCompanyParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithID adds the id to the get company params
func (o *GetCompanyParams) WithID(id int32) *GetCompanyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get company params
func (o *GetCompanyParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompanyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

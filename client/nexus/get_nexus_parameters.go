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

// NewGetNexusParams creates a new GetNexusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNexusParams() *GetNexusParams {
	return &GetNexusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNexusParamsWithTimeout creates a new GetNexusParams object
// with the ability to set a timeout on a request.
func NewGetNexusParamsWithTimeout(timeout time.Duration) *GetNexusParams {
	return &GetNexusParams{
		timeout: timeout,
	}
}

// NewGetNexusParamsWithContext creates a new GetNexusParams object
// with the ability to set a context for a request.
func NewGetNexusParamsWithContext(ctx context.Context) *GetNexusParams {
	return &GetNexusParams{
		Context: ctx,
	}
}

// NewGetNexusParamsWithHTTPClient creates a new GetNexusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNexusParamsWithHTTPClient(client *http.Client) *GetNexusParams {
	return &GetNexusParams{
		HTTPClient: client,
	}
}

/* GetNexusParams contains all the parameters to send to the API endpoint
   for the get nexus operation.

   Typically these are written to a http.Request.
*/
type GetNexusParams struct {

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

	/* ID.

	   The primary key of this nexus

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nexus params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNexusParams) WithDefaults() *GetNexusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nexus params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNexusParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := GetNexusParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get nexus params
func (o *GetNexusParams) WithTimeout(timeout time.Duration) *GetNexusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nexus params
func (o *GetNexusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nexus params
func (o *GetNexusParams) WithContext(ctx context.Context) *GetNexusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nexus params
func (o *GetNexusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nexus params
func (o *GetNexusParams) WithHTTPClient(client *http.Client) *GetNexusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nexus params
func (o *GetNexusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarInclude adds the dollarInclude to the get nexus params
func (o *GetNexusParams) WithDollarInclude(dollarInclude *string) *GetNexusParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the get nexus params
func (o *GetNexusParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithXAvalaraClient adds the xAvalaraClient to the get nexus params
func (o *GetNexusParams) WithXAvalaraClient(xAvalaraClient *string) *GetNexusParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the get nexus params
func (o *GetNexusParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the get nexus params
func (o *GetNexusParams) WithCompanyID(companyID int32) *GetNexusParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the get nexus params
func (o *GetNexusParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the get nexus params
func (o *GetNexusParams) WithID(id int32) *GetNexusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get nexus params
func (o *GetNexusParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetNexusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

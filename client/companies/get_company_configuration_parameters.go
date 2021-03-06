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

// NewGetCompanyConfigurationParams creates a new GetCompanyConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCompanyConfigurationParams() *GetCompanyConfigurationParams {
	return &GetCompanyConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompanyConfigurationParamsWithTimeout creates a new GetCompanyConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetCompanyConfigurationParamsWithTimeout(timeout time.Duration) *GetCompanyConfigurationParams {
	return &GetCompanyConfigurationParams{
		timeout: timeout,
	}
}

// NewGetCompanyConfigurationParamsWithContext creates a new GetCompanyConfigurationParams object
// with the ability to set a context for a request.
func NewGetCompanyConfigurationParamsWithContext(ctx context.Context) *GetCompanyConfigurationParams {
	return &GetCompanyConfigurationParams{
		Context: ctx,
	}
}

// NewGetCompanyConfigurationParamsWithHTTPClient creates a new GetCompanyConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCompanyConfigurationParamsWithHTTPClient(client *http.Client) *GetCompanyConfigurationParams {
	return &GetCompanyConfigurationParams{
		HTTPClient: client,
	}
}

/* GetCompanyConfigurationParams contains all the parameters to send to the API endpoint
   for the get company configuration operation.

   Typically these are written to a http.Request.
*/
type GetCompanyConfigurationParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get company configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompanyConfigurationParams) WithDefaults() *GetCompanyConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get company configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompanyConfigurationParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := GetCompanyConfigurationParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get company configuration params
func (o *GetCompanyConfigurationParams) WithTimeout(timeout time.Duration) *GetCompanyConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get company configuration params
func (o *GetCompanyConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get company configuration params
func (o *GetCompanyConfigurationParams) WithContext(ctx context.Context) *GetCompanyConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get company configuration params
func (o *GetCompanyConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get company configuration params
func (o *GetCompanyConfigurationParams) WithHTTPClient(client *http.Client) *GetCompanyConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get company configuration params
func (o *GetCompanyConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the get company configuration params
func (o *GetCompanyConfigurationParams) WithXAvalaraClient(xAvalaraClient *string) *GetCompanyConfigurationParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the get company configuration params
func (o *GetCompanyConfigurationParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithID adds the id to the get company configuration params
func (o *GetCompanyConfigurationParams) WithID(id int32) *GetCompanyConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get company configuration params
func (o *GetCompanyConfigurationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompanyConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

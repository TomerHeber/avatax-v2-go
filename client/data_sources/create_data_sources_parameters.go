// Code generated by go-swagger; DO NOT EDIT.

package data_sources

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

	"github.com/TomerHeber/avatax-v2-go/models"
)

// NewCreateDataSourcesParams creates a new CreateDataSourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDataSourcesParams() *CreateDataSourcesParams {
	return &CreateDataSourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDataSourcesParamsWithTimeout creates a new CreateDataSourcesParams object
// with the ability to set a timeout on a request.
func NewCreateDataSourcesParamsWithTimeout(timeout time.Duration) *CreateDataSourcesParams {
	return &CreateDataSourcesParams{
		timeout: timeout,
	}
}

// NewCreateDataSourcesParamsWithContext creates a new CreateDataSourcesParams object
// with the ability to set a context for a request.
func NewCreateDataSourcesParamsWithContext(ctx context.Context) *CreateDataSourcesParams {
	return &CreateDataSourcesParams{
		Context: ctx,
	}
}

// NewCreateDataSourcesParamsWithHTTPClient creates a new CreateDataSourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDataSourcesParamsWithHTTPClient(client *http.Client) *CreateDataSourcesParams {
	return &CreateDataSourcesParams{
		HTTPClient: client,
	}
}

/* CreateDataSourcesParams contains all the parameters to send to the API endpoint
   for the create data sources operation.

   Typically these are written to a http.Request.
*/
type CreateDataSourcesParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	// Body.
	Body []*models.DataSourceModel

	/* CompanyID.

	   The id of the company you which to create the datasources

	   Format: int32
	*/
	CompanyID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create data sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDataSourcesParams) WithDefaults() *CreateDataSourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create data sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDataSourcesParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := CreateDataSourcesParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create data sources params
func (o *CreateDataSourcesParams) WithTimeout(timeout time.Duration) *CreateDataSourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create data sources params
func (o *CreateDataSourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create data sources params
func (o *CreateDataSourcesParams) WithContext(ctx context.Context) *CreateDataSourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create data sources params
func (o *CreateDataSourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create data sources params
func (o *CreateDataSourcesParams) WithHTTPClient(client *http.Client) *CreateDataSourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create data sources params
func (o *CreateDataSourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the create data sources params
func (o *CreateDataSourcesParams) WithXAvalaraClient(xAvalaraClient *string) *CreateDataSourcesParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the create data sources params
func (o *CreateDataSourcesParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the create data sources params
func (o *CreateDataSourcesParams) WithBody(body []*models.DataSourceModel) *CreateDataSourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create data sources params
func (o *CreateDataSourcesParams) SetBody(body []*models.DataSourceModel) {
	o.Body = body
}

// WithCompanyID adds the companyID to the create data sources params
func (o *CreateDataSourcesParams) WithCompanyID(companyID int32) *CreateDataSourcesParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the create data sources params
func (o *CreateDataSourcesParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDataSourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param companyId
	if err := r.SetPathParam("companyId", swag.FormatInt32(o.CompanyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
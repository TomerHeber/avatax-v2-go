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
)

// NewDeleteDataSourceParams creates a new DeleteDataSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDataSourceParams() *DeleteDataSourceParams {
	return &DeleteDataSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDataSourceParamsWithTimeout creates a new DeleteDataSourceParams object
// with the ability to set a timeout on a request.
func NewDeleteDataSourceParamsWithTimeout(timeout time.Duration) *DeleteDataSourceParams {
	return &DeleteDataSourceParams{
		timeout: timeout,
	}
}

// NewDeleteDataSourceParamsWithContext creates a new DeleteDataSourceParams object
// with the ability to set a context for a request.
func NewDeleteDataSourceParamsWithContext(ctx context.Context) *DeleteDataSourceParams {
	return &DeleteDataSourceParams{
		Context: ctx,
	}
}

// NewDeleteDataSourceParamsWithHTTPClient creates a new DeleteDataSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDataSourceParamsWithHTTPClient(client *http.Client) *DeleteDataSourceParams {
	return &DeleteDataSourceParams{
		HTTPClient: client,
	}
}

/* DeleteDataSourceParams contains all the parameters to send to the API endpoint
   for the delete data source operation.

   Typically these are written to a http.Request.
*/
type DeleteDataSourceParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyID.

	   The id of the company the datasource belongs to.

	   Format: int32
	*/
	CompanyID int32

	/* ID.

	   The id of the datasource you wish to delete.

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete data source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataSourceParams) WithDefaults() *DeleteDataSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete data source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataSourceParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := DeleteDataSourceParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete data source params
func (o *DeleteDataSourceParams) WithTimeout(timeout time.Duration) *DeleteDataSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete data source params
func (o *DeleteDataSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete data source params
func (o *DeleteDataSourceParams) WithContext(ctx context.Context) *DeleteDataSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete data source params
func (o *DeleteDataSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete data source params
func (o *DeleteDataSourceParams) WithHTTPClient(client *http.Client) *DeleteDataSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete data source params
func (o *DeleteDataSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the delete data source params
func (o *DeleteDataSourceParams) WithXAvalaraClient(xAvalaraClient *string) *DeleteDataSourceParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the delete data source params
func (o *DeleteDataSourceParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the delete data source params
func (o *DeleteDataSourceParams) WithCompanyID(companyID int32) *DeleteDataSourceParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the delete data source params
func (o *DeleteDataSourceParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the delete data source params
func (o *DeleteDataSourceParams) WithID(id int32) *DeleteDataSourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete data source params
func (o *DeleteDataSourceParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDataSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
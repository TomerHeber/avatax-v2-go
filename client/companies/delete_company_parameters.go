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

// NewDeleteCompanyParams creates a new DeleteCompanyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCompanyParams() *DeleteCompanyParams {
	return &DeleteCompanyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCompanyParamsWithTimeout creates a new DeleteCompanyParams object
// with the ability to set a timeout on a request.
func NewDeleteCompanyParamsWithTimeout(timeout time.Duration) *DeleteCompanyParams {
	return &DeleteCompanyParams{
		timeout: timeout,
	}
}

// NewDeleteCompanyParamsWithContext creates a new DeleteCompanyParams object
// with the ability to set a context for a request.
func NewDeleteCompanyParamsWithContext(ctx context.Context) *DeleteCompanyParams {
	return &DeleteCompanyParams{
		Context: ctx,
	}
}

// NewDeleteCompanyParamsWithHTTPClient creates a new DeleteCompanyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCompanyParamsWithHTTPClient(client *http.Client) *DeleteCompanyParams {
	return &DeleteCompanyParams{
		HTTPClient: client,
	}
}

/* DeleteCompanyParams contains all the parameters to send to the API endpoint
   for the delete company operation.

   Typically these are written to a http.Request.
*/
type DeleteCompanyParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* ID.

	   The ID of the company you wish to delete.

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete company params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCompanyParams) WithDefaults() *DeleteCompanyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete company params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCompanyParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := DeleteCompanyParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete company params
func (o *DeleteCompanyParams) WithTimeout(timeout time.Duration) *DeleteCompanyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete company params
func (o *DeleteCompanyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete company params
func (o *DeleteCompanyParams) WithContext(ctx context.Context) *DeleteCompanyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete company params
func (o *DeleteCompanyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete company params
func (o *DeleteCompanyParams) WithHTTPClient(client *http.Client) *DeleteCompanyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete company params
func (o *DeleteCompanyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the delete company params
func (o *DeleteCompanyParams) WithXAvalaraClient(xAvalaraClient *string) *DeleteCompanyParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the delete company params
func (o *DeleteCompanyParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithID adds the id to the delete company params
func (o *DeleteCompanyParams) WithID(id int32) *DeleteCompanyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete company params
func (o *DeleteCompanyParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCompanyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

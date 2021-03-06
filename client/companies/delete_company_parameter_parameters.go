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

// NewDeleteCompanyParameterParams creates a new DeleteCompanyParameterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCompanyParameterParams() *DeleteCompanyParameterParams {
	return &DeleteCompanyParameterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCompanyParameterParamsWithTimeout creates a new DeleteCompanyParameterParams object
// with the ability to set a timeout on a request.
func NewDeleteCompanyParameterParamsWithTimeout(timeout time.Duration) *DeleteCompanyParameterParams {
	return &DeleteCompanyParameterParams{
		timeout: timeout,
	}
}

// NewDeleteCompanyParameterParamsWithContext creates a new DeleteCompanyParameterParams object
// with the ability to set a context for a request.
func NewDeleteCompanyParameterParamsWithContext(ctx context.Context) *DeleteCompanyParameterParams {
	return &DeleteCompanyParameterParams{
		Context: ctx,
	}
}

// NewDeleteCompanyParameterParamsWithHTTPClient creates a new DeleteCompanyParameterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCompanyParameterParamsWithHTTPClient(client *http.Client) *DeleteCompanyParameterParams {
	return &DeleteCompanyParameterParams{
		HTTPClient: client,
	}
}

/* DeleteCompanyParameterParams contains all the parameters to send to the API endpoint
   for the delete company parameter operation.

   Typically these are written to a http.Request.
*/
type DeleteCompanyParameterParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyID.

	   The company id

	   Format: int32
	*/
	CompanyID int32

	/* ID.

	   The parameter id

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete company parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCompanyParameterParams) WithDefaults() *DeleteCompanyParameterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete company parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCompanyParameterParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := DeleteCompanyParameterParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete company parameter params
func (o *DeleteCompanyParameterParams) WithTimeout(timeout time.Duration) *DeleteCompanyParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete company parameter params
func (o *DeleteCompanyParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete company parameter params
func (o *DeleteCompanyParameterParams) WithContext(ctx context.Context) *DeleteCompanyParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete company parameter params
func (o *DeleteCompanyParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete company parameter params
func (o *DeleteCompanyParameterParams) WithHTTPClient(client *http.Client) *DeleteCompanyParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete company parameter params
func (o *DeleteCompanyParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the delete company parameter params
func (o *DeleteCompanyParameterParams) WithXAvalaraClient(xAvalaraClient *string) *DeleteCompanyParameterParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the delete company parameter params
func (o *DeleteCompanyParameterParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the delete company parameter params
func (o *DeleteCompanyParameterParams) WithCompanyID(companyID int32) *DeleteCompanyParameterParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the delete company parameter params
func (o *DeleteCompanyParameterParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the delete company parameter params
func (o *DeleteCompanyParameterParams) WithID(id int64) *DeleteCompanyParameterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete company parameter params
func (o *DeleteCompanyParameterParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCompanyParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

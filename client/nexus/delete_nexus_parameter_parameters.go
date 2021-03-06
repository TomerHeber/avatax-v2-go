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

// NewDeleteNexusParameterParams creates a new DeleteNexusParameterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNexusParameterParams() *DeleteNexusParameterParams {
	return &DeleteNexusParameterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNexusParameterParamsWithTimeout creates a new DeleteNexusParameterParams object
// with the ability to set a timeout on a request.
func NewDeleteNexusParameterParamsWithTimeout(timeout time.Duration) *DeleteNexusParameterParams {
	return &DeleteNexusParameterParams{
		timeout: timeout,
	}
}

// NewDeleteNexusParameterParamsWithContext creates a new DeleteNexusParameterParams object
// with the ability to set a context for a request.
func NewDeleteNexusParameterParamsWithContext(ctx context.Context) *DeleteNexusParameterParams {
	return &DeleteNexusParameterParams{
		Context: ctx,
	}
}

// NewDeleteNexusParameterParamsWithHTTPClient creates a new DeleteNexusParameterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNexusParameterParamsWithHTTPClient(client *http.Client) *DeleteNexusParameterParams {
	return &DeleteNexusParameterParams{
		HTTPClient: client,
	}
}

/* DeleteNexusParameterParams contains all the parameters to send to the API endpoint
   for the delete nexus parameter operation.

   Typically these are written to a http.Request.
*/
type DeleteNexusParameterParams struct {

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

	/* NexusID.

	   The nexus id

	   Format: int32
	*/
	NexusID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete nexus parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNexusParameterParams) WithDefaults() *DeleteNexusParameterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete nexus parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNexusParameterParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := DeleteNexusParameterParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete nexus parameter params
func (o *DeleteNexusParameterParams) WithTimeout(timeout time.Duration) *DeleteNexusParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nexus parameter params
func (o *DeleteNexusParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nexus parameter params
func (o *DeleteNexusParameterParams) WithContext(ctx context.Context) *DeleteNexusParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nexus parameter params
func (o *DeleteNexusParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nexus parameter params
func (o *DeleteNexusParameterParams) WithHTTPClient(client *http.Client) *DeleteNexusParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nexus parameter params
func (o *DeleteNexusParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the delete nexus parameter params
func (o *DeleteNexusParameterParams) WithXAvalaraClient(xAvalaraClient *string) *DeleteNexusParameterParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the delete nexus parameter params
func (o *DeleteNexusParameterParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the delete nexus parameter params
func (o *DeleteNexusParameterParams) WithCompanyID(companyID int32) *DeleteNexusParameterParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the delete nexus parameter params
func (o *DeleteNexusParameterParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the delete nexus parameter params
func (o *DeleteNexusParameterParams) WithID(id int64) *DeleteNexusParameterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete nexus parameter params
func (o *DeleteNexusParameterParams) SetID(id int64) {
	o.ID = id
}

// WithNexusID adds the nexusID to the delete nexus parameter params
func (o *DeleteNexusParameterParams) WithNexusID(nexusID int32) *DeleteNexusParameterParams {
	o.SetNexusID(nexusID)
	return o
}

// SetNexusID adds the nexusId to the delete nexus parameter params
func (o *DeleteNexusParameterParams) SetNexusID(nexusID int32) {
	o.NexusID = nexusID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNexusParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nexusId
	if err := r.SetPathParam("nexusId", swag.FormatInt32(o.NexusID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

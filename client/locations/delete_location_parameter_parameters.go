// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// NewDeleteLocationParameterParams creates a new DeleteLocationParameterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLocationParameterParams() *DeleteLocationParameterParams {
	return &DeleteLocationParameterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLocationParameterParamsWithTimeout creates a new DeleteLocationParameterParams object
// with the ability to set a timeout on a request.
func NewDeleteLocationParameterParamsWithTimeout(timeout time.Duration) *DeleteLocationParameterParams {
	return &DeleteLocationParameterParams{
		timeout: timeout,
	}
}

// NewDeleteLocationParameterParamsWithContext creates a new DeleteLocationParameterParams object
// with the ability to set a context for a request.
func NewDeleteLocationParameterParamsWithContext(ctx context.Context) *DeleteLocationParameterParams {
	return &DeleteLocationParameterParams{
		Context: ctx,
	}
}

// NewDeleteLocationParameterParamsWithHTTPClient creates a new DeleteLocationParameterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLocationParameterParamsWithHTTPClient(client *http.Client) *DeleteLocationParameterParams {
	return &DeleteLocationParameterParams{
		HTTPClient: client,
	}
}

/* DeleteLocationParameterParams contains all the parameters to send to the API endpoint
   for the delete location parameter operation.

   Typically these are written to a http.Request.
*/
type DeleteLocationParameterParams struct {

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

	/* LocationID.

	   The location id

	   Format: int32
	*/
	LocationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete location parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLocationParameterParams) WithDefaults() *DeleteLocationParameterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete location parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLocationParameterParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := DeleteLocationParameterParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete location parameter params
func (o *DeleteLocationParameterParams) WithTimeout(timeout time.Duration) *DeleteLocationParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete location parameter params
func (o *DeleteLocationParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete location parameter params
func (o *DeleteLocationParameterParams) WithContext(ctx context.Context) *DeleteLocationParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete location parameter params
func (o *DeleteLocationParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete location parameter params
func (o *DeleteLocationParameterParams) WithHTTPClient(client *http.Client) *DeleteLocationParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete location parameter params
func (o *DeleteLocationParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the delete location parameter params
func (o *DeleteLocationParameterParams) WithXAvalaraClient(xAvalaraClient *string) *DeleteLocationParameterParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the delete location parameter params
func (o *DeleteLocationParameterParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the delete location parameter params
func (o *DeleteLocationParameterParams) WithCompanyID(companyID int32) *DeleteLocationParameterParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the delete location parameter params
func (o *DeleteLocationParameterParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the delete location parameter params
func (o *DeleteLocationParameterParams) WithID(id int64) *DeleteLocationParameterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete location parameter params
func (o *DeleteLocationParameterParams) SetID(id int64) {
	o.ID = id
}

// WithLocationID adds the locationID to the delete location parameter params
func (o *DeleteLocationParameterParams) WithLocationID(locationID int32) *DeleteLocationParameterParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the delete location parameter params
func (o *DeleteLocationParameterParams) SetLocationID(locationID int32) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLocationParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param locationId
	if err := r.SetPathParam("locationId", swag.FormatInt32(o.LocationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
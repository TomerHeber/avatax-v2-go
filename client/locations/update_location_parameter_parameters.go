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

	"github.com/TomerHeber/avatax-v2-go/models"
)

// NewUpdateLocationParameterParams creates a new UpdateLocationParameterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLocationParameterParams() *UpdateLocationParameterParams {
	return &UpdateLocationParameterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLocationParameterParamsWithTimeout creates a new UpdateLocationParameterParams object
// with the ability to set a timeout on a request.
func NewUpdateLocationParameterParamsWithTimeout(timeout time.Duration) *UpdateLocationParameterParams {
	return &UpdateLocationParameterParams{
		timeout: timeout,
	}
}

// NewUpdateLocationParameterParamsWithContext creates a new UpdateLocationParameterParams object
// with the ability to set a context for a request.
func NewUpdateLocationParameterParamsWithContext(ctx context.Context) *UpdateLocationParameterParams {
	return &UpdateLocationParameterParams{
		Context: ctx,
	}
}

// NewUpdateLocationParameterParamsWithHTTPClient creates a new UpdateLocationParameterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLocationParameterParamsWithHTTPClient(client *http.Client) *UpdateLocationParameterParams {
	return &UpdateLocationParameterParams{
		HTTPClient: client,
	}
}

/* UpdateLocationParameterParams contains all the parameters to send to the API endpoint
   for the update location parameter operation.

   Typically these are written to a http.Request.
*/
type UpdateLocationParameterParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The location parameter object you wish to update.
	*/
	Body *models.LocationParameterModel

	/* CompanyID.

	   The company id.

	   Format: int32
	*/
	CompanyID int32

	/* ID.

	   The location parameter id

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

// WithDefaults hydrates default values in the update location parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLocationParameterParams) WithDefaults() *UpdateLocationParameterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update location parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLocationParameterParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := UpdateLocationParameterParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update location parameter params
func (o *UpdateLocationParameterParams) WithTimeout(timeout time.Duration) *UpdateLocationParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update location parameter params
func (o *UpdateLocationParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update location parameter params
func (o *UpdateLocationParameterParams) WithContext(ctx context.Context) *UpdateLocationParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update location parameter params
func (o *UpdateLocationParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update location parameter params
func (o *UpdateLocationParameterParams) WithHTTPClient(client *http.Client) *UpdateLocationParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update location parameter params
func (o *UpdateLocationParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the update location parameter params
func (o *UpdateLocationParameterParams) WithXAvalaraClient(xAvalaraClient *string) *UpdateLocationParameterParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the update location parameter params
func (o *UpdateLocationParameterParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the update location parameter params
func (o *UpdateLocationParameterParams) WithBody(body *models.LocationParameterModel) *UpdateLocationParameterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update location parameter params
func (o *UpdateLocationParameterParams) SetBody(body *models.LocationParameterModel) {
	o.Body = body
}

// WithCompanyID adds the companyID to the update location parameter params
func (o *UpdateLocationParameterParams) WithCompanyID(companyID int32) *UpdateLocationParameterParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the update location parameter params
func (o *UpdateLocationParameterParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the update location parameter params
func (o *UpdateLocationParameterParams) WithID(id int64) *UpdateLocationParameterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update location parameter params
func (o *UpdateLocationParameterParams) SetID(id int64) {
	o.ID = id
}

// WithLocationID adds the locationID to the update location parameter params
func (o *UpdateLocationParameterParams) WithLocationID(locationID int32) *UpdateLocationParameterParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the update location parameter params
func (o *UpdateLocationParameterParams) SetLocationID(locationID int32) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLocationParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

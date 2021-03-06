// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewUpdateSettingParams creates a new UpdateSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSettingParams() *UpdateSettingParams {
	return &UpdateSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSettingParamsWithTimeout creates a new UpdateSettingParams object
// with the ability to set a timeout on a request.
func NewUpdateSettingParamsWithTimeout(timeout time.Duration) *UpdateSettingParams {
	return &UpdateSettingParams{
		timeout: timeout,
	}
}

// NewUpdateSettingParamsWithContext creates a new UpdateSettingParams object
// with the ability to set a context for a request.
func NewUpdateSettingParamsWithContext(ctx context.Context) *UpdateSettingParams {
	return &UpdateSettingParams{
		Context: ctx,
	}
}

// NewUpdateSettingParamsWithHTTPClient creates a new UpdateSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSettingParamsWithHTTPClient(client *http.Client) *UpdateSettingParams {
	return &UpdateSettingParams{
		HTTPClient: client,
	}
}

/* UpdateSettingParams contains all the parameters to send to the API endpoint
   for the update setting operation.

   Typically these are written to a http.Request.
*/
type UpdateSettingParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The setting you wish to update.
	*/
	Body *models.SettingModel

	/* CompanyID.

	   The ID of the company that this setting belongs to.

	   Format: int32
	*/
	CompanyID int32

	/* ID.

	   The ID of the setting you wish to update

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSettingParams) WithDefaults() *UpdateSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSettingParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := UpdateSettingParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update setting params
func (o *UpdateSettingParams) WithTimeout(timeout time.Duration) *UpdateSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update setting params
func (o *UpdateSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update setting params
func (o *UpdateSettingParams) WithContext(ctx context.Context) *UpdateSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update setting params
func (o *UpdateSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update setting params
func (o *UpdateSettingParams) WithHTTPClient(client *http.Client) *UpdateSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update setting params
func (o *UpdateSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the update setting params
func (o *UpdateSettingParams) WithXAvalaraClient(xAvalaraClient *string) *UpdateSettingParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the update setting params
func (o *UpdateSettingParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the update setting params
func (o *UpdateSettingParams) WithBody(body *models.SettingModel) *UpdateSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update setting params
func (o *UpdateSettingParams) SetBody(body *models.SettingModel) {
	o.Body = body
}

// WithCompanyID adds the companyID to the update setting params
func (o *UpdateSettingParams) WithCompanyID(companyID int32) *UpdateSettingParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the update setting params
func (o *UpdateSettingParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the update setting params
func (o *UpdateSettingParams) WithID(id int32) *UpdateSettingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update setting params
func (o *UpdateSettingParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

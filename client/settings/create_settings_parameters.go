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

// NewCreateSettingsParams creates a new CreateSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSettingsParams() *CreateSettingsParams {
	return &CreateSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSettingsParamsWithTimeout creates a new CreateSettingsParams object
// with the ability to set a timeout on a request.
func NewCreateSettingsParamsWithTimeout(timeout time.Duration) *CreateSettingsParams {
	return &CreateSettingsParams{
		timeout: timeout,
	}
}

// NewCreateSettingsParamsWithContext creates a new CreateSettingsParams object
// with the ability to set a context for a request.
func NewCreateSettingsParamsWithContext(ctx context.Context) *CreateSettingsParams {
	return &CreateSettingsParams{
		Context: ctx,
	}
}

// NewCreateSettingsParamsWithHTTPClient creates a new CreateSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSettingsParamsWithHTTPClient(client *http.Client) *CreateSettingsParams {
	return &CreateSettingsParams{
		HTTPClient: client,
	}
}

/* CreateSettingsParams contains all the parameters to send to the API endpoint
   for the create settings operation.

   Typically these are written to a http.Request.
*/
type CreateSettingsParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The setting you wish to create.
	*/
	Body []*models.SettingModel

	/* CompanyID.

	   The ID of the company that owns this setting.

	   Format: int32
	*/
	CompanyID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSettingsParams) WithDefaults() *CreateSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSettingsParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := CreateSettingsParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create settings params
func (o *CreateSettingsParams) WithTimeout(timeout time.Duration) *CreateSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create settings params
func (o *CreateSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create settings params
func (o *CreateSettingsParams) WithContext(ctx context.Context) *CreateSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create settings params
func (o *CreateSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create settings params
func (o *CreateSettingsParams) WithHTTPClient(client *http.Client) *CreateSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create settings params
func (o *CreateSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the create settings params
func (o *CreateSettingsParams) WithXAvalaraClient(xAvalaraClient *string) *CreateSettingsParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the create settings params
func (o *CreateSettingsParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the create settings params
func (o *CreateSettingsParams) WithBody(body []*models.SettingModel) *CreateSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create settings params
func (o *CreateSettingsParams) SetBody(body []*models.SettingModel) {
	o.Body = body
}

// WithCompanyID adds the companyID to the create settings params
func (o *CreateSettingsParams) WithCompanyID(companyID int32) *CreateSettingsParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the create settings params
func (o *CreateSettingsParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

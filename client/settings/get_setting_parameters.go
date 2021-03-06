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
)

// NewGetSettingParams creates a new GetSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSettingParams() *GetSettingParams {
	return &GetSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSettingParamsWithTimeout creates a new GetSettingParams object
// with the ability to set a timeout on a request.
func NewGetSettingParamsWithTimeout(timeout time.Duration) *GetSettingParams {
	return &GetSettingParams{
		timeout: timeout,
	}
}

// NewGetSettingParamsWithContext creates a new GetSettingParams object
// with the ability to set a context for a request.
func NewGetSettingParamsWithContext(ctx context.Context) *GetSettingParams {
	return &GetSettingParams{
		Context: ctx,
	}
}

// NewGetSettingParamsWithHTTPClient creates a new GetSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSettingParamsWithHTTPClient(client *http.Client) *GetSettingParams {
	return &GetSettingParams{
		HTTPClient: client,
	}
}

/* GetSettingParams contains all the parameters to send to the API endpoint
   for the get setting operation.

   Typically these are written to a http.Request.
*/
type GetSettingParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyID.

	   The ID of the company that owns this setting

	   Format: int32
	*/
	CompanyID int32

	/* ID.

	   The primary key of this setting

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSettingParams) WithDefaults() *GetSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSettingParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := GetSettingParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get setting params
func (o *GetSettingParams) WithTimeout(timeout time.Duration) *GetSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get setting params
func (o *GetSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get setting params
func (o *GetSettingParams) WithContext(ctx context.Context) *GetSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get setting params
func (o *GetSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get setting params
func (o *GetSettingParams) WithHTTPClient(client *http.Client) *GetSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get setting params
func (o *GetSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the get setting params
func (o *GetSettingParams) WithXAvalaraClient(xAvalaraClient *string) *GetSettingParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the get setting params
func (o *GetSettingParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the get setting params
func (o *GetSettingParams) WithCompanyID(companyID int32) *GetSettingParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the get setting params
func (o *GetSettingParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the get setting params
func (o *GetSettingParams) WithID(id int32) *GetSettingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get setting params
func (o *GetSettingParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

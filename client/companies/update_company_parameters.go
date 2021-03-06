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

	"github.com/TomerHeber/avatax-v2-go/models"
)

// NewUpdateCompanyParams creates a new UpdateCompanyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCompanyParams() *UpdateCompanyParams {
	return &UpdateCompanyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCompanyParamsWithTimeout creates a new UpdateCompanyParams object
// with the ability to set a timeout on a request.
func NewUpdateCompanyParamsWithTimeout(timeout time.Duration) *UpdateCompanyParams {
	return &UpdateCompanyParams{
		timeout: timeout,
	}
}

// NewUpdateCompanyParamsWithContext creates a new UpdateCompanyParams object
// with the ability to set a context for a request.
func NewUpdateCompanyParamsWithContext(ctx context.Context) *UpdateCompanyParams {
	return &UpdateCompanyParams{
		Context: ctx,
	}
}

// NewUpdateCompanyParamsWithHTTPClient creates a new UpdateCompanyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCompanyParamsWithHTTPClient(client *http.Client) *UpdateCompanyParams {
	return &UpdateCompanyParams{
		HTTPClient: client,
	}
}

/* UpdateCompanyParams contains all the parameters to send to the API endpoint
   for the update company operation.

   Typically these are written to a http.Request.
*/
type UpdateCompanyParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The company object you wish to update.
	*/
	Body *models.CompanyModel

	/* ID.

	   The ID of the company you wish to update.

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update company params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCompanyParams) WithDefaults() *UpdateCompanyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update company params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCompanyParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := UpdateCompanyParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update company params
func (o *UpdateCompanyParams) WithTimeout(timeout time.Duration) *UpdateCompanyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update company params
func (o *UpdateCompanyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update company params
func (o *UpdateCompanyParams) WithContext(ctx context.Context) *UpdateCompanyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update company params
func (o *UpdateCompanyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update company params
func (o *UpdateCompanyParams) WithHTTPClient(client *http.Client) *UpdateCompanyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update company params
func (o *UpdateCompanyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the update company params
func (o *UpdateCompanyParams) WithXAvalaraClient(xAvalaraClient *string) *UpdateCompanyParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the update company params
func (o *UpdateCompanyParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the update company params
func (o *UpdateCompanyParams) WithBody(body *models.CompanyModel) *UpdateCompanyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update company params
func (o *UpdateCompanyParams) SetBody(body *models.CompanyModel) {
	o.Body = body
}

// WithID adds the id to the update company params
func (o *UpdateCompanyParams) WithID(id int32) *UpdateCompanyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update company params
func (o *UpdateCompanyParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCompanyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

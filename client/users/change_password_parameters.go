// Code generated by go-swagger; DO NOT EDIT.

package users

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

	"github.com/TomerHeber/avatax-v2-go/models"
)

// NewChangePasswordParams creates a new ChangePasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangePasswordParams() *ChangePasswordParams {
	return &ChangePasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangePasswordParamsWithTimeout creates a new ChangePasswordParams object
// with the ability to set a timeout on a request.
func NewChangePasswordParamsWithTimeout(timeout time.Duration) *ChangePasswordParams {
	return &ChangePasswordParams{
		timeout: timeout,
	}
}

// NewChangePasswordParamsWithContext creates a new ChangePasswordParams object
// with the ability to set a context for a request.
func NewChangePasswordParamsWithContext(ctx context.Context) *ChangePasswordParams {
	return &ChangePasswordParams{
		Context: ctx,
	}
}

// NewChangePasswordParamsWithHTTPClient creates a new ChangePasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangePasswordParamsWithHTTPClient(client *http.Client) *ChangePasswordParams {
	return &ChangePasswordParams{
		HTTPClient: client,
	}
}

/* ChangePasswordParams contains all the parameters to send to the API endpoint
   for the change password operation.

   Typically these are written to a http.Request.
*/
type ChangePasswordParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   An object containing your current password and the new password.
	*/
	Body *models.PasswordChangeModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangePasswordParams) WithDefaults() *ChangePasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangePasswordParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ChangePasswordParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the change password params
func (o *ChangePasswordParams) WithTimeout(timeout time.Duration) *ChangePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change password params
func (o *ChangePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change password params
func (o *ChangePasswordParams) WithContext(ctx context.Context) *ChangePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change password params
func (o *ChangePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change password params
func (o *ChangePasswordParams) WithHTTPClient(client *http.Client) *ChangePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change password params
func (o *ChangePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the change password params
func (o *ChangePasswordParams) WithXAvalaraClient(xAvalaraClient *string) *ChangePasswordParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the change password params
func (o *ChangePasswordParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the change password params
func (o *ChangePasswordParams) WithBody(body *models.PasswordChangeModel) *ChangePasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change password params
func (o *ChangePasswordParams) SetBody(body *models.PasswordChangeModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
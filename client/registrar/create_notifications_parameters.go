// Code generated by go-swagger; DO NOT EDIT.

package registrar

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

// NewCreateNotificationsParams creates a new CreateNotificationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNotificationsParams() *CreateNotificationsParams {
	return &CreateNotificationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNotificationsParamsWithTimeout creates a new CreateNotificationsParams object
// with the ability to set a timeout on a request.
func NewCreateNotificationsParamsWithTimeout(timeout time.Duration) *CreateNotificationsParams {
	return &CreateNotificationsParams{
		timeout: timeout,
	}
}

// NewCreateNotificationsParamsWithContext creates a new CreateNotificationsParams object
// with the ability to set a context for a request.
func NewCreateNotificationsParamsWithContext(ctx context.Context) *CreateNotificationsParams {
	return &CreateNotificationsParams{
		Context: ctx,
	}
}

// NewCreateNotificationsParamsWithHTTPClient creates a new CreateNotificationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNotificationsParamsWithHTTPClient(client *http.Client) *CreateNotificationsParams {
	return &CreateNotificationsParams{
		HTTPClient: client,
	}
}

/* CreateNotificationsParams contains all the parameters to send to the API endpoint
   for the create notifications operation.

   Typically these are written to a http.Request.
*/
type CreateNotificationsParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The notifications you wish to create.
	*/
	Body []*models.NotificationModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNotificationsParams) WithDefaults() *CreateNotificationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNotificationsParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := CreateNotificationsParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create notifications params
func (o *CreateNotificationsParams) WithTimeout(timeout time.Duration) *CreateNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create notifications params
func (o *CreateNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create notifications params
func (o *CreateNotificationsParams) WithContext(ctx context.Context) *CreateNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create notifications params
func (o *CreateNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create notifications params
func (o *CreateNotificationsParams) WithHTTPClient(client *http.Client) *CreateNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create notifications params
func (o *CreateNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the create notifications params
func (o *CreateNotificationsParams) WithXAvalaraClient(xAvalaraClient *string) *CreateNotificationsParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the create notifications params
func (o *CreateNotificationsParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the create notifications params
func (o *CreateNotificationsParams) WithBody(body []*models.NotificationModel) *CreateNotificationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create notifications params
func (o *CreateNotificationsParams) SetBody(body []*models.NotificationModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
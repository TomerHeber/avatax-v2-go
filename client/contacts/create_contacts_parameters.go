// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewCreateContactsParams creates a new CreateContactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateContactsParams() *CreateContactsParams {
	return &CreateContactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateContactsParamsWithTimeout creates a new CreateContactsParams object
// with the ability to set a timeout on a request.
func NewCreateContactsParamsWithTimeout(timeout time.Duration) *CreateContactsParams {
	return &CreateContactsParams{
		timeout: timeout,
	}
}

// NewCreateContactsParamsWithContext creates a new CreateContactsParams object
// with the ability to set a context for a request.
func NewCreateContactsParamsWithContext(ctx context.Context) *CreateContactsParams {
	return &CreateContactsParams{
		Context: ctx,
	}
}

// NewCreateContactsParamsWithHTTPClient creates a new CreateContactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateContactsParamsWithHTTPClient(client *http.Client) *CreateContactsParams {
	return &CreateContactsParams{
		HTTPClient: client,
	}
}

/* CreateContactsParams contains all the parameters to send to the API endpoint
   for the create contacts operation.

   Typically these are written to a http.Request.
*/
type CreateContactsParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The contacts you wish to create.
	*/
	Body []*models.ContactModel

	/* CompanyID.

	   The ID of the company that owns this contact.

	   Format: int32
	*/
	CompanyID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create contacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateContactsParams) WithDefaults() *CreateContactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create contacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateContactsParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := CreateContactsParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create contacts params
func (o *CreateContactsParams) WithTimeout(timeout time.Duration) *CreateContactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create contacts params
func (o *CreateContactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create contacts params
func (o *CreateContactsParams) WithContext(ctx context.Context) *CreateContactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create contacts params
func (o *CreateContactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create contacts params
func (o *CreateContactsParams) WithHTTPClient(client *http.Client) *CreateContactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create contacts params
func (o *CreateContactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the create contacts params
func (o *CreateContactsParams) WithXAvalaraClient(xAvalaraClient *string) *CreateContactsParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the create contacts params
func (o *CreateContactsParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the create contacts params
func (o *CreateContactsParams) WithBody(body []*models.ContactModel) *CreateContactsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create contacts params
func (o *CreateContactsParams) SetBody(body []*models.ContactModel) {
	o.Body = body
}

// WithCompanyID adds the companyID to the create contacts params
func (o *CreateContactsParams) WithCompanyID(companyID int32) *CreateContactsParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the create contacts params
func (o *CreateContactsParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateContactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
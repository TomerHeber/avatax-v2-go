// Code generated by go-swagger; DO NOT EDIT.

package certificates

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

// NewListAttributesForCertificateParams creates a new ListAttributesForCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAttributesForCertificateParams() *ListAttributesForCertificateParams {
	return &ListAttributesForCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAttributesForCertificateParamsWithTimeout creates a new ListAttributesForCertificateParams object
// with the ability to set a timeout on a request.
func NewListAttributesForCertificateParamsWithTimeout(timeout time.Duration) *ListAttributesForCertificateParams {
	return &ListAttributesForCertificateParams{
		timeout: timeout,
	}
}

// NewListAttributesForCertificateParamsWithContext creates a new ListAttributesForCertificateParams object
// with the ability to set a context for a request.
func NewListAttributesForCertificateParamsWithContext(ctx context.Context) *ListAttributesForCertificateParams {
	return &ListAttributesForCertificateParams{
		Context: ctx,
	}
}

// NewListAttributesForCertificateParamsWithHTTPClient creates a new ListAttributesForCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAttributesForCertificateParamsWithHTTPClient(client *http.Client) *ListAttributesForCertificateParams {
	return &ListAttributesForCertificateParams{
		HTTPClient: client,
	}
}

/* ListAttributesForCertificateParams contains all the parameters to send to the API endpoint
   for the list attributes for certificate operation.

   Typically these are written to a http.Request.
*/
type ListAttributesForCertificateParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyID.

	   The unique ID number of the company that recorded this certificate

	   Format: int32
	*/
	CompanyID int32

	/* ID.

	   The unique ID number of this certificate

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list attributes for certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAttributesForCertificateParams) WithDefaults() *ListAttributesForCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list attributes for certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAttributesForCertificateParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListAttributesForCertificateParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) WithTimeout(timeout time.Duration) *ListAttributesForCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) WithContext(ctx context.Context) *ListAttributesForCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) WithHTTPClient(client *http.Client) *ListAttributesForCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) WithXAvalaraClient(xAvalaraClient *string) *ListAttributesForCertificateParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) WithCompanyID(companyID int32) *ListAttributesForCertificateParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) WithID(id int32) *ListAttributesForCertificateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list attributes for certificate params
func (o *ListAttributesForCertificateParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListAttributesForCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
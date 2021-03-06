// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewListValidCertificatesForCustomerParams creates a new ListValidCertificatesForCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListValidCertificatesForCustomerParams() *ListValidCertificatesForCustomerParams {
	return &ListValidCertificatesForCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListValidCertificatesForCustomerParamsWithTimeout creates a new ListValidCertificatesForCustomerParams object
// with the ability to set a timeout on a request.
func NewListValidCertificatesForCustomerParamsWithTimeout(timeout time.Duration) *ListValidCertificatesForCustomerParams {
	return &ListValidCertificatesForCustomerParams{
		timeout: timeout,
	}
}

// NewListValidCertificatesForCustomerParamsWithContext creates a new ListValidCertificatesForCustomerParams object
// with the ability to set a context for a request.
func NewListValidCertificatesForCustomerParamsWithContext(ctx context.Context) *ListValidCertificatesForCustomerParams {
	return &ListValidCertificatesForCustomerParams{
		Context: ctx,
	}
}

// NewListValidCertificatesForCustomerParamsWithHTTPClient creates a new ListValidCertificatesForCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewListValidCertificatesForCustomerParamsWithHTTPClient(client *http.Client) *ListValidCertificatesForCustomerParams {
	return &ListValidCertificatesForCustomerParams{
		HTTPClient: client,
	}
}

/* ListValidCertificatesForCustomerParams contains all the parameters to send to the API endpoint
   for the list valid certificates for customer operation.

   Typically these are written to a http.Request.
*/
type ListValidCertificatesForCustomerParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyID.

	   The unique ID number of the company that recorded this customer

	   Format: int32
	*/
	CompanyID int32

	/* Country.

	   Search for certificates matching this country.  Uses the ISO 3166 two character country code.
	*/
	Country string

	/* CustomerCode.

	   The unique code representing this customer
	*/
	CustomerCode string

	/* Region.

	   Search for certificates matching this region.  Uses the ISO 3166 two or three character state, region, or province code.
	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list valid certificates for customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListValidCertificatesForCustomerParams) WithDefaults() *ListValidCertificatesForCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list valid certificates for customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListValidCertificatesForCustomerParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListValidCertificatesForCustomerParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) WithTimeout(timeout time.Duration) *ListValidCertificatesForCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) WithContext(ctx context.Context) *ListValidCertificatesForCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) WithHTTPClient(client *http.Client) *ListValidCertificatesForCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) WithXAvalaraClient(xAvalaraClient *string) *ListValidCertificatesForCustomerParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) WithCompanyID(companyID int32) *ListValidCertificatesForCustomerParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithCountry adds the country to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) WithCountry(country string) *ListValidCertificatesForCustomerParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) SetCountry(country string) {
	o.Country = country
}

// WithCustomerCode adds the customerCode to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) WithCustomerCode(customerCode string) *ListValidCertificatesForCustomerParams {
	o.SetCustomerCode(customerCode)
	return o
}

// SetCustomerCode adds the customerCode to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) SetCustomerCode(customerCode string) {
	o.CustomerCode = customerCode
}

// WithRegion adds the region to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) WithRegion(region string) *ListValidCertificatesForCustomerParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the list valid certificates for customer params
func (o *ListValidCertificatesForCustomerParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *ListValidCertificatesForCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param country
	if err := r.SetPathParam("country", o.Country); err != nil {
		return err
	}

	// path param customerCode
	if err := r.SetPathParam("customerCode", o.CustomerCode); err != nil {
		return err
	}

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

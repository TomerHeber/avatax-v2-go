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

// NewListAttributesForCustomerParams creates a new ListAttributesForCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAttributesForCustomerParams() *ListAttributesForCustomerParams {
	return &ListAttributesForCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAttributesForCustomerParamsWithTimeout creates a new ListAttributesForCustomerParams object
// with the ability to set a timeout on a request.
func NewListAttributesForCustomerParamsWithTimeout(timeout time.Duration) *ListAttributesForCustomerParams {
	return &ListAttributesForCustomerParams{
		timeout: timeout,
	}
}

// NewListAttributesForCustomerParamsWithContext creates a new ListAttributesForCustomerParams object
// with the ability to set a context for a request.
func NewListAttributesForCustomerParamsWithContext(ctx context.Context) *ListAttributesForCustomerParams {
	return &ListAttributesForCustomerParams{
		Context: ctx,
	}
}

// NewListAttributesForCustomerParamsWithHTTPClient creates a new ListAttributesForCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAttributesForCustomerParamsWithHTTPClient(client *http.Client) *ListAttributesForCustomerParams {
	return &ListAttributesForCustomerParams{
		HTTPClient: client,
	}
}

/* ListAttributesForCustomerParams contains all the parameters to send to the API endpoint
   for the list attributes for customer operation.

   Typically these are written to a http.Request.
*/
type ListAttributesForCustomerParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyID.

	   The unique ID number of the company that recorded the provided customer

	   Format: int32
	*/
	CompanyID int32

	/* CustomerCode.

	   The unique code representing the current customer
	*/
	CustomerCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list attributes for customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAttributesForCustomerParams) WithDefaults() *ListAttributesForCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list attributes for customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAttributesForCustomerParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListAttributesForCustomerParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list attributes for customer params
func (o *ListAttributesForCustomerParams) WithTimeout(timeout time.Duration) *ListAttributesForCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list attributes for customer params
func (o *ListAttributesForCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list attributes for customer params
func (o *ListAttributesForCustomerParams) WithContext(ctx context.Context) *ListAttributesForCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list attributes for customer params
func (o *ListAttributesForCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list attributes for customer params
func (o *ListAttributesForCustomerParams) WithHTTPClient(client *http.Client) *ListAttributesForCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list attributes for customer params
func (o *ListAttributesForCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the list attributes for customer params
func (o *ListAttributesForCustomerParams) WithXAvalaraClient(xAvalaraClient *string) *ListAttributesForCustomerParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list attributes for customer params
func (o *ListAttributesForCustomerParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the list attributes for customer params
func (o *ListAttributesForCustomerParams) WithCompanyID(companyID int32) *ListAttributesForCustomerParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the list attributes for customer params
func (o *ListAttributesForCustomerParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithCustomerCode adds the customerCode to the list attributes for customer params
func (o *ListAttributesForCustomerParams) WithCustomerCode(customerCode string) *ListAttributesForCustomerParams {
	o.SetCustomerCode(customerCode)
	return o
}

// SetCustomerCode adds the customerCode to the list attributes for customer params
func (o *ListAttributesForCustomerParams) SetCustomerCode(customerCode string) {
	o.CustomerCode = customerCode
}

// WriteToRequest writes these params to a swagger request
func (o *ListAttributesForCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param customerCode
	if err := r.SetPathParam("customerCode", o.CustomerCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
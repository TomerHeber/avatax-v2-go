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

// NewQueryCustomersParams creates a new QueryCustomersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryCustomersParams() *QueryCustomersParams {
	return &QueryCustomersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryCustomersParamsWithTimeout creates a new QueryCustomersParams object
// with the ability to set a timeout on a request.
func NewQueryCustomersParamsWithTimeout(timeout time.Duration) *QueryCustomersParams {
	return &QueryCustomersParams{
		timeout: timeout,
	}
}

// NewQueryCustomersParamsWithContext creates a new QueryCustomersParams object
// with the ability to set a context for a request.
func NewQueryCustomersParamsWithContext(ctx context.Context) *QueryCustomersParams {
	return &QueryCustomersParams{
		Context: ctx,
	}
}

// NewQueryCustomersParamsWithHTTPClient creates a new QueryCustomersParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryCustomersParamsWithHTTPClient(client *http.Client) *QueryCustomersParams {
	return &QueryCustomersParams{
		HTTPClient: client,
	}
}

/* QueryCustomersParams contains all the parameters to send to the API endpoint
   for the query customers operation.

   Typically these are written to a http.Request.
*/
type QueryCustomersParams struct {

	/* DollarFilter.

	   A filter statement to identify specific records to retrieve.  For more information on filtering, see [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/).<br />*Not filterable:* shipTos
	*/
	DollarFilter *string

	/* DollarInclude.

	   OPTIONAL - You can specify the value `certificates` to fetch information about certificates linked to the customer.
	*/
	DollarInclude *string

	/* DollarOrderBy.

	   A comma separated list of sort statements in the format `(fieldname) [ASC|DESC]`, for example `id ASC`.
	*/
	DollarOrderBy *string

	/* DollarSkip.

	   If nonzero, skip this number of results before returning data.  Used with `$top` to provide pagination for large datasets.

	   Format: int32
	*/
	DollarSkip *int32

	/* DollarTop.

	   If nonzero, return no more than this number of results.  Used with `$skip` to provide pagination for large datasets.  Unless otherwise specified, the maximum number of records that can be returned from an API call is 1,000 records.

	   Format: int32
	*/
	DollarTop *int32

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query customers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryCustomersParams) WithDefaults() *QueryCustomersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query customers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryCustomersParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := QueryCustomersParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the query customers params
func (o *QueryCustomersParams) WithTimeout(timeout time.Duration) *QueryCustomersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query customers params
func (o *QueryCustomersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query customers params
func (o *QueryCustomersParams) WithContext(ctx context.Context) *QueryCustomersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query customers params
func (o *QueryCustomersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query customers params
func (o *QueryCustomersParams) WithHTTPClient(client *http.Client) *QueryCustomersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query customers params
func (o *QueryCustomersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the query customers params
func (o *QueryCustomersParams) WithDollarFilter(dollarFilter *string) *QueryCustomersParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the query customers params
func (o *QueryCustomersParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarInclude adds the dollarInclude to the query customers params
func (o *QueryCustomersParams) WithDollarInclude(dollarInclude *string) *QueryCustomersParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the query customers params
func (o *QueryCustomersParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithDollarOrderBy adds the dollarOrderBy to the query customers params
func (o *QueryCustomersParams) WithDollarOrderBy(dollarOrderBy *string) *QueryCustomersParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the query customers params
func (o *QueryCustomersParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the query customers params
func (o *QueryCustomersParams) WithDollarSkip(dollarSkip *int32) *QueryCustomersParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the query customers params
func (o *QueryCustomersParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the query customers params
func (o *QueryCustomersParams) WithDollarTop(dollarTop *int32) *QueryCustomersParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the query customers params
func (o *QueryCustomersParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the query customers params
func (o *QueryCustomersParams) WithXAvalaraClient(xAvalaraClient *string) *QueryCustomersParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the query customers params
func (o *QueryCustomersParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the query customers params
func (o *QueryCustomersParams) WithCompanyID(companyID int32) *QueryCustomersParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the query customers params
func (o *QueryCustomersParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryCustomersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string

		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {

			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}
	}

	if o.DollarInclude != nil {

		// query param $include
		var qrDollarInclude string

		if o.DollarInclude != nil {
			qrDollarInclude = *o.DollarInclude
		}
		qDollarInclude := qrDollarInclude
		if qDollarInclude != "" {

			if err := r.SetQueryParam("$include", qDollarInclude); err != nil {
				return err
			}
		}
	}

	if o.DollarOrderBy != nil {

		// query param $orderBy
		var qrDollarOrderBy string

		if o.DollarOrderBy != nil {
			qrDollarOrderBy = *o.DollarOrderBy
		}
		qDollarOrderBy := qrDollarOrderBy
		if qDollarOrderBy != "" {

			if err := r.SetQueryParam("$orderBy", qDollarOrderBy); err != nil {
				return err
			}
		}
	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int32

		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt32(qrDollarSkip)
		if qDollarSkip != "" {

			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}
	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int32

		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt32(qrDollarTop)
		if qDollarTop != "" {

			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package definitions

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

// NewListRateTypesByCountryParams creates a new ListRateTypesByCountryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRateTypesByCountryParams() *ListRateTypesByCountryParams {
	return &ListRateTypesByCountryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRateTypesByCountryParamsWithTimeout creates a new ListRateTypesByCountryParams object
// with the ability to set a timeout on a request.
func NewListRateTypesByCountryParamsWithTimeout(timeout time.Duration) *ListRateTypesByCountryParams {
	return &ListRateTypesByCountryParams{
		timeout: timeout,
	}
}

// NewListRateTypesByCountryParamsWithContext creates a new ListRateTypesByCountryParams object
// with the ability to set a context for a request.
func NewListRateTypesByCountryParamsWithContext(ctx context.Context) *ListRateTypesByCountryParams {
	return &ListRateTypesByCountryParams{
		Context: ctx,
	}
}

// NewListRateTypesByCountryParamsWithHTTPClient creates a new ListRateTypesByCountryParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRateTypesByCountryParamsWithHTTPClient(client *http.Client) *ListRateTypesByCountryParams {
	return &ListRateTypesByCountryParams{
		HTTPClient: client,
	}
}

/* ListRateTypesByCountryParams contains all the parameters to send to the API endpoint
   for the list rate types by country operation.

   Typically these are written to a http.Request.
*/
type ListRateTypesByCountryParams struct {

	/* DollarFilter.

	   A filter statement to identify specific records to retrieve.  For more information on filtering, see [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/).
	*/
	DollarFilter *string

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

	/* Country.

	   The country to examine for rate types
	*/
	Country string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list rate types by country params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRateTypesByCountryParams) WithDefaults() *ListRateTypesByCountryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list rate types by country params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRateTypesByCountryParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListRateTypesByCountryParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithTimeout(timeout time.Duration) *ListRateTypesByCountryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithContext(ctx context.Context) *ListRateTypesByCountryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithHTTPClient(client *http.Client) *ListRateTypesByCountryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithDollarFilter(dollarFilter *string) *ListRateTypesByCountryParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderBy adds the dollarOrderBy to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithDollarOrderBy(dollarOrderBy *string) *ListRateTypesByCountryParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithDollarSkip(dollarSkip *int32) *ListRateTypesByCountryParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithDollarTop(dollarTop *int32) *ListRateTypesByCountryParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithXAvalaraClient(xAvalaraClient *string) *ListRateTypesByCountryParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCountry adds the country to the list rate types by country params
func (o *ListRateTypesByCountryParams) WithCountry(country string) *ListRateTypesByCountryParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the list rate types by country params
func (o *ListRateTypesByCountryParams) SetCountry(country string) {
	o.Country = country
}

// WriteToRequest writes these params to a swagger request
func (o *ListRateTypesByCountryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param country
	if err := r.SetPathParam("country", o.Country); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
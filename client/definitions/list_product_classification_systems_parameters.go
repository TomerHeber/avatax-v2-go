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

// NewListProductClassificationSystemsParams creates a new ListProductClassificationSystemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProductClassificationSystemsParams() *ListProductClassificationSystemsParams {
	return &ListProductClassificationSystemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProductClassificationSystemsParamsWithTimeout creates a new ListProductClassificationSystemsParams object
// with the ability to set a timeout on a request.
func NewListProductClassificationSystemsParamsWithTimeout(timeout time.Duration) *ListProductClassificationSystemsParams {
	return &ListProductClassificationSystemsParams{
		timeout: timeout,
	}
}

// NewListProductClassificationSystemsParamsWithContext creates a new ListProductClassificationSystemsParams object
// with the ability to set a context for a request.
func NewListProductClassificationSystemsParamsWithContext(ctx context.Context) *ListProductClassificationSystemsParams {
	return &ListProductClassificationSystemsParams{
		Context: ctx,
	}
}

// NewListProductClassificationSystemsParamsWithHTTPClient creates a new ListProductClassificationSystemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProductClassificationSystemsParamsWithHTTPClient(client *http.Client) *ListProductClassificationSystemsParams {
	return &ListProductClassificationSystemsParams{
		HTTPClient: client,
	}
}

/* ListProductClassificationSystemsParams contains all the parameters to send to the API endpoint
   for the list product classification systems operation.

   Typically these are written to a http.Request.
*/
type ListProductClassificationSystemsParams struct {

	/* DollarCountryCode.

	   If not null, return all records with this code.
	*/
	DollarCountryCode *string

	/* DollarFilter.

	   A filter statement to identify specific records to retrieve.  For more information on filtering, see [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/).<br />*Not filterable:* countries
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list product classification systems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProductClassificationSystemsParams) WithDefaults() *ListProductClassificationSystemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list product classification systems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProductClassificationSystemsParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListProductClassificationSystemsParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithTimeout(timeout time.Duration) *ListProductClassificationSystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithContext(ctx context.Context) *ListProductClassificationSystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithHTTPClient(client *http.Client) *ListProductClassificationSystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCountryCode adds the dollarCountryCode to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithDollarCountryCode(dollarCountryCode *string) *ListProductClassificationSystemsParams {
	o.SetDollarCountryCode(dollarCountryCode)
	return o
}

// SetDollarCountryCode adds the dollarCountryCode to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetDollarCountryCode(dollarCountryCode *string) {
	o.DollarCountryCode = dollarCountryCode
}

// WithDollarFilter adds the dollarFilter to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithDollarFilter(dollarFilter *string) *ListProductClassificationSystemsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderBy adds the dollarOrderBy to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithDollarOrderBy(dollarOrderBy *string) *ListProductClassificationSystemsParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithDollarSkip(dollarSkip *int32) *ListProductClassificationSystemsParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithDollarTop(dollarTop *int32) *ListProductClassificationSystemsParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the list product classification systems params
func (o *ListProductClassificationSystemsParams) WithXAvalaraClient(xAvalaraClient *string) *ListProductClassificationSystemsParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list product classification systems params
func (o *ListProductClassificationSystemsParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WriteToRequest writes these params to a swagger request
func (o *ListProductClassificationSystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarCountryCode != nil {

		// query param $countryCode
		var qrDollarCountryCode string

		if o.DollarCountryCode != nil {
			qrDollarCountryCode = *o.DollarCountryCode
		}
		qDollarCountryCode := qrDollarCountryCode
		if qDollarCountryCode != "" {

			if err := r.SetQueryParam("$countryCode", qDollarCountryCode); err != nil {
				return err
			}
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
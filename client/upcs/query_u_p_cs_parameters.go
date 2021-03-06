// Code generated by go-swagger; DO NOT EDIT.

package upcs

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

// NewQueryUPCsParams creates a new QueryUPCsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryUPCsParams() *QueryUPCsParams {
	return &QueryUPCsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryUPCsParamsWithTimeout creates a new QueryUPCsParams object
// with the ability to set a timeout on a request.
func NewQueryUPCsParamsWithTimeout(timeout time.Duration) *QueryUPCsParams {
	return &QueryUPCsParams{
		timeout: timeout,
	}
}

// NewQueryUPCsParamsWithContext creates a new QueryUPCsParams object
// with the ability to set a context for a request.
func NewQueryUPCsParamsWithContext(ctx context.Context) *QueryUPCsParams {
	return &QueryUPCsParams{
		Context: ctx,
	}
}

// NewQueryUPCsParamsWithHTTPClient creates a new QueryUPCsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryUPCsParamsWithHTTPClient(client *http.Client) *QueryUPCsParams {
	return &QueryUPCsParams{
		HTTPClient: client,
	}
}

/* QueryUPCsParams contains all the parameters to send to the API endpoint
   for the query u p cs operation.

   Typically these are written to a http.Request.
*/
type QueryUPCsParams struct {

	/* DollarFilter.

	   A filter statement to identify specific records to retrieve.  For more information on filtering, see [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/).
	*/
	DollarFilter *string

	/* DollarInclude.

	   A comma separated list of additional data to retrieve.
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query u p cs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryUPCsParams) WithDefaults() *QueryUPCsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query u p cs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryUPCsParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := QueryUPCsParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the query u p cs params
func (o *QueryUPCsParams) WithTimeout(timeout time.Duration) *QueryUPCsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query u p cs params
func (o *QueryUPCsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query u p cs params
func (o *QueryUPCsParams) WithContext(ctx context.Context) *QueryUPCsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query u p cs params
func (o *QueryUPCsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query u p cs params
func (o *QueryUPCsParams) WithHTTPClient(client *http.Client) *QueryUPCsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query u p cs params
func (o *QueryUPCsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the query u p cs params
func (o *QueryUPCsParams) WithDollarFilter(dollarFilter *string) *QueryUPCsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the query u p cs params
func (o *QueryUPCsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarInclude adds the dollarInclude to the query u p cs params
func (o *QueryUPCsParams) WithDollarInclude(dollarInclude *string) *QueryUPCsParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the query u p cs params
func (o *QueryUPCsParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithDollarOrderBy adds the dollarOrderBy to the query u p cs params
func (o *QueryUPCsParams) WithDollarOrderBy(dollarOrderBy *string) *QueryUPCsParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the query u p cs params
func (o *QueryUPCsParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the query u p cs params
func (o *QueryUPCsParams) WithDollarSkip(dollarSkip *int32) *QueryUPCsParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the query u p cs params
func (o *QueryUPCsParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the query u p cs params
func (o *QueryUPCsParams) WithDollarTop(dollarTop *int32) *QueryUPCsParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the query u p cs params
func (o *QueryUPCsParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the query u p cs params
func (o *QueryUPCsParams) WithXAvalaraClient(xAvalaraClient *string) *QueryUPCsParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the query u p cs params
func (o *QueryUPCsParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WriteToRequest writes these params to a swagger request
func (o *QueryUPCsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewListNoticeFilingtypesParams creates a new ListNoticeFilingtypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNoticeFilingtypesParams() *ListNoticeFilingtypesParams {
	return &ListNoticeFilingtypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNoticeFilingtypesParamsWithTimeout creates a new ListNoticeFilingtypesParams object
// with the ability to set a timeout on a request.
func NewListNoticeFilingtypesParamsWithTimeout(timeout time.Duration) *ListNoticeFilingtypesParams {
	return &ListNoticeFilingtypesParams{
		timeout: timeout,
	}
}

// NewListNoticeFilingtypesParamsWithContext creates a new ListNoticeFilingtypesParams object
// with the ability to set a context for a request.
func NewListNoticeFilingtypesParamsWithContext(ctx context.Context) *ListNoticeFilingtypesParams {
	return &ListNoticeFilingtypesParams{
		Context: ctx,
	}
}

// NewListNoticeFilingtypesParamsWithHTTPClient creates a new ListNoticeFilingtypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNoticeFilingtypesParamsWithHTTPClient(client *http.Client) *ListNoticeFilingtypesParams {
	return &ListNoticeFilingtypesParams{
		HTTPClient: client,
	}
}

/* ListNoticeFilingtypesParams contains all the parameters to send to the API endpoint
   for the list notice filingtypes operation.

   Typically these are written to a http.Request.
*/
type ListNoticeFilingtypesParams struct {

	/* DollarFilter.

	   A filter statement to identify specific records to retrieve.  For more information on filtering, see [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/).<br />*Not filterable:* description, activeFlag, sortOrder
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

// WithDefaults hydrates default values in the list notice filingtypes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNoticeFilingtypesParams) WithDefaults() *ListNoticeFilingtypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list notice filingtypes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNoticeFilingtypesParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListNoticeFilingtypesParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) WithTimeout(timeout time.Duration) *ListNoticeFilingtypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) WithContext(ctx context.Context) *ListNoticeFilingtypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) WithHTTPClient(client *http.Client) *ListNoticeFilingtypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) WithDollarFilter(dollarFilter *string) *ListNoticeFilingtypesParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderBy adds the dollarOrderBy to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) WithDollarOrderBy(dollarOrderBy *string) *ListNoticeFilingtypesParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) WithDollarSkip(dollarSkip *int32) *ListNoticeFilingtypesParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) WithDollarTop(dollarTop *int32) *ListNoticeFilingtypesParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) WithXAvalaraClient(xAvalaraClient *string) *ListNoticeFilingtypesParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list notice filingtypes params
func (o *ListNoticeFilingtypesParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WriteToRequest writes these params to a swagger request
func (o *ListNoticeFilingtypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewListNotificationsParams creates a new ListNotificationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNotificationsParams() *ListNotificationsParams {
	return &ListNotificationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNotificationsParamsWithTimeout creates a new ListNotificationsParams object
// with the ability to set a timeout on a request.
func NewListNotificationsParamsWithTimeout(timeout time.Duration) *ListNotificationsParams {
	return &ListNotificationsParams{
		timeout: timeout,
	}
}

// NewListNotificationsParamsWithContext creates a new ListNotificationsParams object
// with the ability to set a context for a request.
func NewListNotificationsParamsWithContext(ctx context.Context) *ListNotificationsParams {
	return &ListNotificationsParams{
		Context: ctx,
	}
}

// NewListNotificationsParamsWithHTTPClient creates a new ListNotificationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNotificationsParamsWithHTTPClient(client *http.Client) *ListNotificationsParams {
	return &ListNotificationsParams{
		HTTPClient: client,
	}
}

/* ListNotificationsParams contains all the parameters to send to the API endpoint
   for the list notifications operation.

   Typically these are written to a http.Request.
*/
type ListNotificationsParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNotificationsParams) WithDefaults() *ListNotificationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNotificationsParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListNotificationsParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list notifications params
func (o *ListNotificationsParams) WithTimeout(timeout time.Duration) *ListNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list notifications params
func (o *ListNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list notifications params
func (o *ListNotificationsParams) WithContext(ctx context.Context) *ListNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list notifications params
func (o *ListNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list notifications params
func (o *ListNotificationsParams) WithHTTPClient(client *http.Client) *ListNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list notifications params
func (o *ListNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the list notifications params
func (o *ListNotificationsParams) WithDollarFilter(dollarFilter *string) *ListNotificationsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the list notifications params
func (o *ListNotificationsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderBy adds the dollarOrderBy to the list notifications params
func (o *ListNotificationsParams) WithDollarOrderBy(dollarOrderBy *string) *ListNotificationsParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the list notifications params
func (o *ListNotificationsParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the list notifications params
func (o *ListNotificationsParams) WithDollarSkip(dollarSkip *int32) *ListNotificationsParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list notifications params
func (o *ListNotificationsParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list notifications params
func (o *ListNotificationsParams) WithDollarTop(dollarTop *int32) *ListNotificationsParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list notifications params
func (o *ListNotificationsParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the list notifications params
func (o *ListNotificationsParams) WithXAvalaraClient(xAvalaraClient *string) *ListNotificationsParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list notifications params
func (o *ListNotificationsParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WriteToRequest writes these params to a swagger request
func (o *ListNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewListTaxSubTypesByJurisdictionAndRegionParams creates a new ListTaxSubTypesByJurisdictionAndRegionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTaxSubTypesByJurisdictionAndRegionParams() *ListTaxSubTypesByJurisdictionAndRegionParams {
	return &ListTaxSubTypesByJurisdictionAndRegionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTaxSubTypesByJurisdictionAndRegionParamsWithTimeout creates a new ListTaxSubTypesByJurisdictionAndRegionParams object
// with the ability to set a timeout on a request.
func NewListTaxSubTypesByJurisdictionAndRegionParamsWithTimeout(timeout time.Duration) *ListTaxSubTypesByJurisdictionAndRegionParams {
	return &ListTaxSubTypesByJurisdictionAndRegionParams{
		timeout: timeout,
	}
}

// NewListTaxSubTypesByJurisdictionAndRegionParamsWithContext creates a new ListTaxSubTypesByJurisdictionAndRegionParams object
// with the ability to set a context for a request.
func NewListTaxSubTypesByJurisdictionAndRegionParamsWithContext(ctx context.Context) *ListTaxSubTypesByJurisdictionAndRegionParams {
	return &ListTaxSubTypesByJurisdictionAndRegionParams{
		Context: ctx,
	}
}

// NewListTaxSubTypesByJurisdictionAndRegionParamsWithHTTPClient creates a new ListTaxSubTypesByJurisdictionAndRegionParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTaxSubTypesByJurisdictionAndRegionParamsWithHTTPClient(client *http.Client) *ListTaxSubTypesByJurisdictionAndRegionParams {
	return &ListTaxSubTypesByJurisdictionAndRegionParams{
		HTTPClient: client,
	}
}

/* ListTaxSubTypesByJurisdictionAndRegionParams contains all the parameters to send to the API endpoint
   for the list tax sub types by jurisdiction and region operation.

   Typically these are written to a http.Request.
*/
type ListTaxSubTypesByJurisdictionAndRegionParams struct {

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

	/* JurisdictionCode.

	   The jurisdiction code of the tax sub type.
	*/
	JurisdictionCode string

	/* Region.

	   The region of the tax sub type.
	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list tax sub types by jurisdiction and region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithDefaults() *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list tax sub types by jurisdiction and region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListTaxSubTypesByJurisdictionAndRegionParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithTimeout(timeout time.Duration) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithContext(ctx context.Context) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithHTTPClient(client *http.Client) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithDollarFilter(dollarFilter *string) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderBy adds the dollarOrderBy to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithDollarOrderBy(dollarOrderBy *string) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithDollarSkip(dollarSkip *int32) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithDollarTop(dollarTop *int32) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithXAvalaraClient(xAvalaraClient *string) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithJurisdictionCode adds the jurisdictionCode to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithJurisdictionCode(jurisdictionCode string) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetJurisdictionCode(jurisdictionCode)
	return o
}

// SetJurisdictionCode adds the jurisdictionCode to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetJurisdictionCode(jurisdictionCode string) {
	o.JurisdictionCode = jurisdictionCode
}

// WithRegion adds the region to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WithRegion(region string) *ListTaxSubTypesByJurisdictionAndRegionParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the list tax sub types by jurisdiction and region params
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *ListTaxSubTypesByJurisdictionAndRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param jurisdictionCode
	if err := r.SetPathParam("jurisdictionCode", o.JurisdictionCode); err != nil {
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
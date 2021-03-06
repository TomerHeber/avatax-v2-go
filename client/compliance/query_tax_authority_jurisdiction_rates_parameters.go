// Code generated by go-swagger; DO NOT EDIT.

package compliance

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

// NewQueryTaxAuthorityJurisdictionRatesParams creates a new QueryTaxAuthorityJurisdictionRatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryTaxAuthorityJurisdictionRatesParams() *QueryTaxAuthorityJurisdictionRatesParams {
	return &QueryTaxAuthorityJurisdictionRatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryTaxAuthorityJurisdictionRatesParamsWithTimeout creates a new QueryTaxAuthorityJurisdictionRatesParams object
// with the ability to set a timeout on a request.
func NewQueryTaxAuthorityJurisdictionRatesParamsWithTimeout(timeout time.Duration) *QueryTaxAuthorityJurisdictionRatesParams {
	return &QueryTaxAuthorityJurisdictionRatesParams{
		timeout: timeout,
	}
}

// NewQueryTaxAuthorityJurisdictionRatesParamsWithContext creates a new QueryTaxAuthorityJurisdictionRatesParams object
// with the ability to set a context for a request.
func NewQueryTaxAuthorityJurisdictionRatesParamsWithContext(ctx context.Context) *QueryTaxAuthorityJurisdictionRatesParams {
	return &QueryTaxAuthorityJurisdictionRatesParams{
		Context: ctx,
	}
}

// NewQueryTaxAuthorityJurisdictionRatesParamsWithHTTPClient creates a new QueryTaxAuthorityJurisdictionRatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryTaxAuthorityJurisdictionRatesParamsWithHTTPClient(client *http.Client) *QueryTaxAuthorityJurisdictionRatesParams {
	return &QueryTaxAuthorityJurisdictionRatesParams{
		HTTPClient: client,
	}
}

/* QueryTaxAuthorityJurisdictionRatesParams contains all the parameters to send to the API endpoint
   for the query tax authority jurisdiction rates operation.

   Typically these are written to a http.Request.
*/
type QueryTaxAuthorityJurisdictionRatesParams struct {

	/* DollarFilter.

	   A filter statement to identify specific records to retrieve.  For more information on filtering, see [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/).
	*/
	DollarFilter *string

	/* DollarInclude.

	   A comma separated list of objects to fetch underneath this jurisdiction.
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

	/* EffectiveDate.

	   Used to limit the jurisdictions returned.

	   Format: date-time
	*/
	EffectiveDate *strfmt.DateTime

	/* EndDate.

	   Used to limit the jurisdictions returned.

	   Format: date-time
	*/
	EndDate *strfmt.DateTime

	/* TaxAuthorityID.

	   Used to limit the jurisdictions returned.

	   Format: int32
	*/
	TaxAuthorityID *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query tax authority jurisdiction rates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithDefaults() *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query tax authority jurisdiction rates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := QueryTaxAuthorityJurisdictionRatesParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithTimeout(timeout time.Duration) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithContext(ctx context.Context) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithHTTPClient(client *http.Client) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithDollarFilter(dollarFilter *string) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarInclude adds the dollarInclude to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithDollarInclude(dollarInclude *string) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithDollarOrderBy adds the dollarOrderBy to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithDollarOrderBy(dollarOrderBy *string) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithDollarSkip(dollarSkip *int32) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithDollarTop(dollarTop *int32) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithXAvalaraClient(xAvalaraClient *string) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithEffectiveDate adds the effectiveDate to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithEffectiveDate(effectiveDate *strfmt.DateTime) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetEffectiveDate(effectiveDate *strfmt.DateTime) {
	o.EffectiveDate = effectiveDate
}

// WithEndDate adds the endDate to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithEndDate(endDate *strfmt.DateTime) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithTaxAuthorityID adds the taxAuthorityID to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) WithTaxAuthorityID(taxAuthorityID *int32) *QueryTaxAuthorityJurisdictionRatesParams {
	o.SetTaxAuthorityID(taxAuthorityID)
	return o
}

// SetTaxAuthorityID adds the taxAuthorityId to the query tax authority jurisdiction rates params
func (o *QueryTaxAuthorityJurisdictionRatesParams) SetTaxAuthorityID(taxAuthorityID *int32) {
	o.TaxAuthorityID = taxAuthorityID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryTaxAuthorityJurisdictionRatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EffectiveDate != nil {

		// query param effectiveDate
		var qrEffectiveDate strfmt.DateTime

		if o.EffectiveDate != nil {
			qrEffectiveDate = *o.EffectiveDate
		}
		qEffectiveDate := qrEffectiveDate.String()
		if qEffectiveDate != "" {

			if err := r.SetQueryParam("effectiveDate", qEffectiveDate); err != nil {
				return err
			}
		}
	}

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate strfmt.DateTime

		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {

			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}
	}

	if o.TaxAuthorityID != nil {

		// query param taxAuthorityId
		var qrTaxAuthorityID int32

		if o.TaxAuthorityID != nil {
			qrTaxAuthorityID = *o.TaxAuthorityID
		}
		qTaxAuthorityID := swag.FormatInt32(qrTaxAuthorityID)
		if qTaxAuthorityID != "" {

			if err := r.SetQueryParam("taxAuthorityId", qTaxAuthorityID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

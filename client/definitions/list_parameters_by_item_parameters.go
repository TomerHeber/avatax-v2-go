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

// NewListParametersByItemParams creates a new ListParametersByItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListParametersByItemParams() *ListParametersByItemParams {
	return &ListParametersByItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListParametersByItemParamsWithTimeout creates a new ListParametersByItemParams object
// with the ability to set a timeout on a request.
func NewListParametersByItemParamsWithTimeout(timeout time.Duration) *ListParametersByItemParams {
	return &ListParametersByItemParams{
		timeout: timeout,
	}
}

// NewListParametersByItemParamsWithContext creates a new ListParametersByItemParams object
// with the ability to set a context for a request.
func NewListParametersByItemParamsWithContext(ctx context.Context) *ListParametersByItemParams {
	return &ListParametersByItemParams{
		Context: ctx,
	}
}

// NewListParametersByItemParamsWithHTTPClient creates a new ListParametersByItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewListParametersByItemParamsWithHTTPClient(client *http.Client) *ListParametersByItemParams {
	return &ListParametersByItemParams{
		HTTPClient: client,
	}
}

/* ListParametersByItemParams contains all the parameters to send to the API endpoint
   for the list parameters by item operation.

   Typically these are written to a http.Request.
*/
type ListParametersByItemParams struct {

	/* DollarFilter.

	   A filter statement to identify specific records to retrieve.  For more information on filtering, see [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/).<br />*Not filterable:* serviceTypes, regularExpression, values
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

	/* CompanyCode.

	   Company code.
	*/
	CompanyCode string

	/* ItemCode.

	   Item code.
	*/
	ItemCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list parameters by item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListParametersByItemParams) WithDefaults() *ListParametersByItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list parameters by item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListParametersByItemParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListParametersByItemParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list parameters by item params
func (o *ListParametersByItemParams) WithTimeout(timeout time.Duration) *ListParametersByItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list parameters by item params
func (o *ListParametersByItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list parameters by item params
func (o *ListParametersByItemParams) WithContext(ctx context.Context) *ListParametersByItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list parameters by item params
func (o *ListParametersByItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list parameters by item params
func (o *ListParametersByItemParams) WithHTTPClient(client *http.Client) *ListParametersByItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list parameters by item params
func (o *ListParametersByItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the list parameters by item params
func (o *ListParametersByItemParams) WithDollarFilter(dollarFilter *string) *ListParametersByItemParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the list parameters by item params
func (o *ListParametersByItemParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderBy adds the dollarOrderBy to the list parameters by item params
func (o *ListParametersByItemParams) WithDollarOrderBy(dollarOrderBy *string) *ListParametersByItemParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the list parameters by item params
func (o *ListParametersByItemParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the list parameters by item params
func (o *ListParametersByItemParams) WithDollarSkip(dollarSkip *int32) *ListParametersByItemParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list parameters by item params
func (o *ListParametersByItemParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list parameters by item params
func (o *ListParametersByItemParams) WithDollarTop(dollarTop *int32) *ListParametersByItemParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list parameters by item params
func (o *ListParametersByItemParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the list parameters by item params
func (o *ListParametersByItemParams) WithXAvalaraClient(xAvalaraClient *string) *ListParametersByItemParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list parameters by item params
func (o *ListParametersByItemParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyCode adds the companyCode to the list parameters by item params
func (o *ListParametersByItemParams) WithCompanyCode(companyCode string) *ListParametersByItemParams {
	o.SetCompanyCode(companyCode)
	return o
}

// SetCompanyCode adds the companyCode to the list parameters by item params
func (o *ListParametersByItemParams) SetCompanyCode(companyCode string) {
	o.CompanyCode = companyCode
}

// WithItemCode adds the itemCode to the list parameters by item params
func (o *ListParametersByItemParams) WithItemCode(itemCode string) *ListParametersByItemParams {
	o.SetItemCode(itemCode)
	return o
}

// SetItemCode adds the itemCode to the list parameters by item params
func (o *ListParametersByItemParams) SetItemCode(itemCode string) {
	o.ItemCode = itemCode
}

// WriteToRequest writes these params to a swagger request
func (o *ListParametersByItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param companyCode
	if err := r.SetPathParam("companyCode", o.CompanyCode); err != nil {
		return err
	}

	// path param itemCode
	if err := r.SetPathParam("itemCode", o.ItemCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
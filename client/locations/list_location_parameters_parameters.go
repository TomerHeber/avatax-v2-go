// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// NewListLocationParametersParams creates a new ListLocationParametersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListLocationParametersParams() *ListLocationParametersParams {
	return &ListLocationParametersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListLocationParametersParamsWithTimeout creates a new ListLocationParametersParams object
// with the ability to set a timeout on a request.
func NewListLocationParametersParamsWithTimeout(timeout time.Duration) *ListLocationParametersParams {
	return &ListLocationParametersParams{
		timeout: timeout,
	}
}

// NewListLocationParametersParamsWithContext creates a new ListLocationParametersParams object
// with the ability to set a context for a request.
func NewListLocationParametersParamsWithContext(ctx context.Context) *ListLocationParametersParams {
	return &ListLocationParametersParams{
		Context: ctx,
	}
}

// NewListLocationParametersParamsWithHTTPClient creates a new ListLocationParametersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListLocationParametersParamsWithHTTPClient(client *http.Client) *ListLocationParametersParams {
	return &ListLocationParametersParams{
		HTTPClient: client,
	}
}

/* ListLocationParametersParams contains all the parameters to send to the API endpoint
   for the list location parameters operation.

   Typically these are written to a http.Request.
*/
type ListLocationParametersParams struct {

	/* DollarFilter.

	   A filter statement to identify specific records to retrieve.  For more information on filtering, see [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/).<br />*Not filterable:* name, unit
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

	/* CompanyID.

	   The company id

	   Format: int32
	*/
	CompanyID int32

	/* LocationID.

	   The ID of the location

	   Format: int32
	*/
	LocationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list location parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLocationParametersParams) WithDefaults() *ListLocationParametersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list location parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLocationParametersParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := ListLocationParametersParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list location parameters params
func (o *ListLocationParametersParams) WithTimeout(timeout time.Duration) *ListLocationParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list location parameters params
func (o *ListLocationParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list location parameters params
func (o *ListLocationParametersParams) WithContext(ctx context.Context) *ListLocationParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list location parameters params
func (o *ListLocationParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list location parameters params
func (o *ListLocationParametersParams) WithHTTPClient(client *http.Client) *ListLocationParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list location parameters params
func (o *ListLocationParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the list location parameters params
func (o *ListLocationParametersParams) WithDollarFilter(dollarFilter *string) *ListLocationParametersParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the list location parameters params
func (o *ListLocationParametersParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderBy adds the dollarOrderBy to the list location parameters params
func (o *ListLocationParametersParams) WithDollarOrderBy(dollarOrderBy *string) *ListLocationParametersParams {
	o.SetDollarOrderBy(dollarOrderBy)
	return o
}

// SetDollarOrderBy adds the dollarOrderBy to the list location parameters params
func (o *ListLocationParametersParams) SetDollarOrderBy(dollarOrderBy *string) {
	o.DollarOrderBy = dollarOrderBy
}

// WithDollarSkip adds the dollarSkip to the list location parameters params
func (o *ListLocationParametersParams) WithDollarSkip(dollarSkip *int32) *ListLocationParametersParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list location parameters params
func (o *ListLocationParametersParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list location parameters params
func (o *ListLocationParametersParams) WithDollarTop(dollarTop *int32) *ListLocationParametersParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list location parameters params
func (o *ListLocationParametersParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithXAvalaraClient adds the xAvalaraClient to the list location parameters params
func (o *ListLocationParametersParams) WithXAvalaraClient(xAvalaraClient *string) *ListLocationParametersParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the list location parameters params
func (o *ListLocationParametersParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the list location parameters params
func (o *ListLocationParametersParams) WithCompanyID(companyID int32) *ListLocationParametersParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the list location parameters params
func (o *ListLocationParametersParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithLocationID adds the locationID to the list location parameters params
func (o *ListLocationParametersParams) WithLocationID(locationID int32) *ListLocationParametersParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list location parameters params
func (o *ListLocationParametersParams) SetLocationID(locationID int32) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *ListLocationParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param companyId
	if err := r.SetPathParam("companyId", swag.FormatInt32(o.CompanyID)); err != nil {
		return err
	}

	// path param locationId
	if err := r.SetPathParam("locationId", swag.FormatInt32(o.LocationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
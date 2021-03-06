// Code generated by go-swagger; DO NOT EDIT.

package tax_content

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
)

// NewTaxRatesByAddressParams creates a new TaxRatesByAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTaxRatesByAddressParams() *TaxRatesByAddressParams {
	return &TaxRatesByAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTaxRatesByAddressParamsWithTimeout creates a new TaxRatesByAddressParams object
// with the ability to set a timeout on a request.
func NewTaxRatesByAddressParamsWithTimeout(timeout time.Duration) *TaxRatesByAddressParams {
	return &TaxRatesByAddressParams{
		timeout: timeout,
	}
}

// NewTaxRatesByAddressParamsWithContext creates a new TaxRatesByAddressParams object
// with the ability to set a context for a request.
func NewTaxRatesByAddressParamsWithContext(ctx context.Context) *TaxRatesByAddressParams {
	return &TaxRatesByAddressParams{
		Context: ctx,
	}
}

// NewTaxRatesByAddressParamsWithHTTPClient creates a new TaxRatesByAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewTaxRatesByAddressParamsWithHTTPClient(client *http.Client) *TaxRatesByAddressParams {
	return &TaxRatesByAddressParams{
		HTTPClient: client,
	}
}

/* TaxRatesByAddressParams contains all the parameters to send to the API endpoint
   for the tax rates by address operation.

   Typically these are written to a http.Request.
*/
type TaxRatesByAddressParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* City.

	   The city name of the location.
	*/
	City *string

	/* Country.

	     Name or ISO 3166 code identifying the country.

	This field supports many different country identifiers:
	 * Two character ISO 3166 codes
	 * Three character ISO 3166 codes
	 * Fully spelled out names of the country in ISO supported languages
	 * Common alternative spellings for many countries

	For a full list of all supported codes and names, please see the Definitions API `ListCountries`.
	*/
	Country string

	/* Line1.

	   The street address of the location.
	*/
	Line1 string

	/* Line2.

	   The street address of the location.
	*/
	Line2 *string

	/* Line3.

	   The street address of the location.
	*/
	Line3 *string

	/* PostalCode.

	   The postal code of the location.
	*/
	PostalCode string

	/* Region.

	     Name or ISO 3166 code identifying the region within the country.

	This field supports many different region identifiers:
	 * Two and three character ISO 3166 region codes
	 * Fully spelled out names of the region in ISO supported languages
	 * Common alternative spellings for many regions

	For a full list of all supported codes and names, please see the Definitions API `ListRegions`.
	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tax rates by address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaxRatesByAddressParams) WithDefaults() *TaxRatesByAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tax rates by address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaxRatesByAddressParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := TaxRatesByAddressParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the tax rates by address params
func (o *TaxRatesByAddressParams) WithTimeout(timeout time.Duration) *TaxRatesByAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tax rates by address params
func (o *TaxRatesByAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tax rates by address params
func (o *TaxRatesByAddressParams) WithContext(ctx context.Context) *TaxRatesByAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tax rates by address params
func (o *TaxRatesByAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tax rates by address params
func (o *TaxRatesByAddressParams) WithHTTPClient(client *http.Client) *TaxRatesByAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tax rates by address params
func (o *TaxRatesByAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the tax rates by address params
func (o *TaxRatesByAddressParams) WithXAvalaraClient(xAvalaraClient *string) *TaxRatesByAddressParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the tax rates by address params
func (o *TaxRatesByAddressParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCity adds the city to the tax rates by address params
func (o *TaxRatesByAddressParams) WithCity(city *string) *TaxRatesByAddressParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the tax rates by address params
func (o *TaxRatesByAddressParams) SetCity(city *string) {
	o.City = city
}

// WithCountry adds the country to the tax rates by address params
func (o *TaxRatesByAddressParams) WithCountry(country string) *TaxRatesByAddressParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the tax rates by address params
func (o *TaxRatesByAddressParams) SetCountry(country string) {
	o.Country = country
}

// WithLine1 adds the line1 to the tax rates by address params
func (o *TaxRatesByAddressParams) WithLine1(line1 string) *TaxRatesByAddressParams {
	o.SetLine1(line1)
	return o
}

// SetLine1 adds the line1 to the tax rates by address params
func (o *TaxRatesByAddressParams) SetLine1(line1 string) {
	o.Line1 = line1
}

// WithLine2 adds the line2 to the tax rates by address params
func (o *TaxRatesByAddressParams) WithLine2(line2 *string) *TaxRatesByAddressParams {
	o.SetLine2(line2)
	return o
}

// SetLine2 adds the line2 to the tax rates by address params
func (o *TaxRatesByAddressParams) SetLine2(line2 *string) {
	o.Line2 = line2
}

// WithLine3 adds the line3 to the tax rates by address params
func (o *TaxRatesByAddressParams) WithLine3(line3 *string) *TaxRatesByAddressParams {
	o.SetLine3(line3)
	return o
}

// SetLine3 adds the line3 to the tax rates by address params
func (o *TaxRatesByAddressParams) SetLine3(line3 *string) {
	o.Line3 = line3
}

// WithPostalCode adds the postalCode to the tax rates by address params
func (o *TaxRatesByAddressParams) WithPostalCode(postalCode string) *TaxRatesByAddressParams {
	o.SetPostalCode(postalCode)
	return o
}

// SetPostalCode adds the postalCode to the tax rates by address params
func (o *TaxRatesByAddressParams) SetPostalCode(postalCode string) {
	o.PostalCode = postalCode
}

// WithRegion adds the region to the tax rates by address params
func (o *TaxRatesByAddressParams) WithRegion(region string) *TaxRatesByAddressParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the tax rates by address params
func (o *TaxRatesByAddressParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *TaxRatesByAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XAvalaraClient != nil {

		// header param X-Avalara-Client
		if err := r.SetHeaderParam("X-Avalara-Client", *o.XAvalaraClient); err != nil {
			return err
		}
	}

	if o.City != nil {

		// query param city
		var qrCity string

		if o.City != nil {
			qrCity = *o.City
		}
		qCity := qrCity
		if qCity != "" {

			if err := r.SetQueryParam("city", qCity); err != nil {
				return err
			}
		}
	}

	// query param country
	qrCountry := o.Country
	qCountry := qrCountry
	if qCountry != "" {

		if err := r.SetQueryParam("country", qCountry); err != nil {
			return err
		}
	}

	// query param line1
	qrLine1 := o.Line1
	qLine1 := qrLine1
	if qLine1 != "" {

		if err := r.SetQueryParam("line1", qLine1); err != nil {
			return err
		}
	}

	if o.Line2 != nil {

		// query param line2
		var qrLine2 string

		if o.Line2 != nil {
			qrLine2 = *o.Line2
		}
		qLine2 := qrLine2
		if qLine2 != "" {

			if err := r.SetQueryParam("line2", qLine2); err != nil {
				return err
			}
		}
	}

	if o.Line3 != nil {

		// query param line3
		var qrLine3 string

		if o.Line3 != nil {
			qrLine3 = *o.Line3
		}
		qLine3 := qrLine3
		if qLine3 != "" {

			if err := r.SetQueryParam("line3", qLine3); err != nil {
				return err
			}
		}
	}

	// query param postalCode
	qrPostalCode := o.PostalCode
	qPostalCode := qrPostalCode
	if qPostalCode != "" {

		if err := r.SetQueryParam("postalCode", qPostalCode); err != nil {
			return err
		}
	}

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {

		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

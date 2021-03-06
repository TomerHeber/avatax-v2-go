// Code generated by go-swagger; DO NOT EDIT.

package tax_content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tax content API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tax content API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BuildTaxContentFile(params *BuildTaxContentFileParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*BuildTaxContentFileOK, error)

	BuildTaxContentFileForLocation(params *BuildTaxContentFileForLocationParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*BuildTaxContentFileForLocationOK, error)

	DownloadTaxRatesByZipCode(params *DownloadTaxRatesByZipCodeParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*DownloadTaxRatesByZipCodeOK, error)

	TaxRatesByAddress(params *TaxRatesByAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TaxRatesByAddressOK, error)

	TaxRatesByPostalCode(params *TaxRatesByPostalCodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TaxRatesByPostalCodeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BuildTaxContentFile builds a multi location tax content file

  Builds a tax content file containing information useful for a retail point-of-sale solution.

Since tax rates may change based on decisions made by a variety of tax authorities, we recommend
that users of this tax content API download new data every day.  Many tax authorities may finalize
decisions on tax changes at unexpected times and may make changes in response to legal issues or
governmental priorities.  Any tax content downloaded for future time periods is subject to change
if tax rates or tax laws change.

A TaxContent file contains a matrix of the taxes that would be charged when you sell any of your
Items at any of your Locations.  To create items, use `CreateItems()`.  To create locations, use
`CreateLocations()`.  The file is built by looking up the tax profile for your location and your
item and calculating taxes for each in turn.  To include a custom `TaxCode` in this tax content
file, first create the custom tax code using `CreateTaxCodes()` to create the custom tax code,
then use `CreateItems()` to create an item that uses the custom tax code.

This data file can be customized for specific partner devices and usage conditions.

The result of this API is the file you requested in the format you requested using the `responseType` field.

This API builds the file on demand, and is limited to files with no more than 7500 scenarios.  To build a tax content
file for a single location at a time, please use `BuildTaxContentFileForLocation`.

NOTE: This API does not work for Tennessee tax holiday scenarios.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountOperator, AccountUser, CompanyAdmin, CompanyUser, Compliance Root User, ComplianceAdmin, ComplianceUser, CSPAdmin, CSPTester, FirmAdmin, FirmUser, ProStoresOperator, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser, TreasuryAdmin, TreasuryUser.
* This API depends on the following active services:*Required* (all):  AvaTaxPro.

*/
func (a *Client) BuildTaxContentFile(params *BuildTaxContentFileParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*BuildTaxContentFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildTaxContentFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BuildTaxContentFile",
		Method:             "POST",
		PathPattern:        "/api/v2/pointofsaledata/build",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildTaxContentFileReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BuildTaxContentFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BuildTaxContentFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BuildTaxContentFileForLocation builds a tax content file for a single location

  Builds a tax content file containing information useful for a retail point-of-sale solution.

Since tax rates may change based on decisions made by a variety of tax authorities, we recommend
that users of this tax content API download new data every day.  Many tax authorities may finalize
decisions on tax changes at unexpected times and may make changes in response to legal issues or
governmental priorities.  Any tax content downloaded for future time periods is subject to change
if tax rates or tax laws change.

A TaxContent file contains a matrix of the taxes that would be charged when you sell any of your
Items at any of your Locations.  To create items, use `CreateItems()`.  To create locations, use
`CreateLocations()`.  The file is built by looking up the tax profile for your location and your
item and calculating taxes for each in turn.  To include a custom `TaxCode` in this tax content
file, first create the custom tax code using `CreateTaxCodes()` to create the custom tax code,
then use `CreateItems()` to create an item that uses the custom tax code.

This data file can be customized for specific partner devices and usage conditions.

The result of this API is the file you requested in the format you requested using the `responseType` field.

This API builds the file on demand, and is limited to files with no more than 7500 scenarios.  To build a tax content
file for a multiple locations in a single file, please use `BuildTaxContentFile`.

NOTE: This API does not work for Tennessee tax holiday scenarios.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountOperator, AccountUser, CompanyAdmin, CompanyUser, Compliance Root User, ComplianceAdmin, ComplianceUser, CSPAdmin, CSPTester, FirmAdmin, FirmUser, ProStoresOperator, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser, TreasuryAdmin, TreasuryUser.
* This API depends on the following active services:*Required* (all):  AvaTaxPro.

*/
func (a *Client) BuildTaxContentFileForLocation(params *BuildTaxContentFileForLocationParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*BuildTaxContentFileForLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildTaxContentFileForLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BuildTaxContentFileForLocation",
		Method:             "GET",
		PathPattern:        "/api/v2/companies/{companyId}/locations/{id}/pointofsaledata",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildTaxContentFileForLocationReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BuildTaxContentFileForLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BuildTaxContentFileForLocation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DownloadTaxRatesByZipCode downloads a file listing tax rates by postal code

  Download a CSV file containing all five digit postal codes in the United States and their sales
and use tax rates for tangible personal property.

Since tax rates may change based on decisions made by a variety of tax authorities, we recommend
that users of this tax content API download new data every day.  Many tax authorities may finalize
decisions on tax changes at unexpected times and may make changes in response to legal issues or
governmental priorities.  Any tax content downloaded for future time periods is subject to change
if tax rates or tax laws change.

This rates file is intended to be used as a default for tax calculation when your software cannot
call the `CreateTransaction` API call.  When using this file, your software will be unable to
handle complex tax rules such as:

* Zip+4 - This tax file contains five digit zip codes only.
* Different product types - This tax file contains tangible personal property tax rates only.
* Mixed sourcing - This tax file cannot be used to resolve origin-based taxes.
* Threshold-based taxes - This tax file does not contain information about thresholds.

If you use this file to provide default tax rates, please ensure that your software calls `CreateTransaction`
to reconcile the actual transaction and determine the difference between the estimated general tax
rate and the final transaction tax.

The file provided by this API is in CSV format with the following columns:

* ZIP_CODE - The five digit zip code for this record.
* STATE_ABBREV - A valid two character US state abbreviation for this record.  Zip codes may span multiple states.
* COUNTY_NAME - A valid county name for this record.  Zip codes may span multiple counties.
* CITY_NAME - A valid city name for this record.  Zip codes may span multiple cities.
* STATE_SALES_TAX - The state component of the sales tax rate.
* STATE_USE_TAX - The state component of the use tax rate.
* COUNTY_SALES_TAX - The county component of the sales tax rate.
* COUNTY_USE_TAX - The county component of the use tax rate.
* CITY_SALES_TAX - The city component of the sales tax rate.
* CITY_USE_TAX - The city component of the use tax rate.
* TOTAL_SALES_TAX - The total tax rate for sales tax for this postal code.  This value may not equal the sum of the state/county/city due to special tax jurisdiction rules.
* TOTAL_USE_TAX - The total tax rate for use tax for this postal code.  This value may not equal the sum of the state/county/city due to special tax jurisdiction rules.
* TAX_SHIPPING_ALONE - This column contains 'Y' if shipping is taxable.
* TAX_SHIPPING_AND_HANDLING_TOGETHER - This column contains 'Y' if shipping and handling are taxable when sent together.

For more detailed tax content, please use the `BuildTaxContentFile` API which allows usage of exact items and exact locations.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountOperator, AccountUser, CompanyAdmin, CompanyUser, Compliance Root User, ComplianceAdmin, ComplianceUser, CSPAdmin, CSPTester, FirmAdmin, FirmUser, ProStoresOperator, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser, TreasuryAdmin, TreasuryUser.

*/
func (a *Client) DownloadTaxRatesByZipCode(params *DownloadTaxRatesByZipCodeParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*DownloadTaxRatesByZipCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadTaxRatesByZipCodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DownloadTaxRatesByZipCode",
		Method:             "GET",
		PathPattern:        "/api/v2/taxratesbyzipcode/download/{date}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadTaxRatesByZipCodeReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DownloadTaxRatesByZipCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DownloadTaxRatesByZipCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TaxRatesByAddress sales tax rates for a specified address

  Usage of this API is subject to rate limits.  Users who exceed the rate limit will receive HTTP
response code 429 - `Too Many Requests`.

This API assumes that you are selling general tangible personal property at a retail point-of-sale
location in the United States only.

For more powerful tax calculation, please consider upgrading to the `CreateTransaction` API,
which supports features including, but not limited to:

* Nexus declarations
* Taxability based on product/service type
* Sourcing rules affecting origin/destination states
* Customers who are exempt from certain taxes
* States that have dollar value thresholds for tax amounts
* Refunds for products purchased on a different date
* Detailed jurisdiction names and state assigned codes
* And more!

Please see [Estimating Tax with REST v2](http://developer.avalara.com/blog/2016/11/04/estimating-tax-with-rest-v2/)
for information on how to upgrade to the full AvaTax CreateTransaction API.
*/
func (a *Client) TaxRatesByAddress(params *TaxRatesByAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TaxRatesByAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTaxRatesByAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TaxRatesByAddress",
		Method:             "GET",
		PathPattern:        "/api/v2/taxrates/byaddress",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TaxRatesByAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TaxRatesByAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TaxRatesByAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TaxRatesByPostalCode sales tax rates for a specified country and postal code this API is only available for u s postal codes

  This API is only available for a US postal codes.

Usage of this API is subject to rate limits.  Users who exceed the rate limit will receive HTTP
response code 429 - `Too Many Requests`.

This API assumes that you are selling general tangible personal property at a retail point-of-sale
location in the United States only.

For more powerful tax calculation, please consider upgrading to the `CreateTransaction` API,
which supports features including, but not limited to:

* Nexus declarations
* Taxability based on product/service type
* Sourcing rules affecting origin/destination states
* Customers who are exempt from certain taxes
* States that have dollar value thresholds for tax amounts
* Refunds for products purchased on a different date
* Detailed jurisdiction names and state assigned codes
* And more!

Please see [Estimating Tax with REST v2](http://developer.avalara.com/blog/2016/11/04/estimating-tax-with-rest-v2/)
for information on how to upgrade to the full AvaTax CreateTransaction API.
*/
func (a *Client) TaxRatesByPostalCode(params *TaxRatesByPostalCodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TaxRatesByPostalCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTaxRatesByPostalCodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TaxRatesByPostalCode",
		Method:             "GET",
		PathPattern:        "/api/v2/taxrates/bypostalcode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TaxRatesByPostalCodeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TaxRatesByPostalCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TaxRatesByPostalCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

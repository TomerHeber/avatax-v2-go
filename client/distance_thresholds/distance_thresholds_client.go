// Code generated by go-swagger; DO NOT EDIT.

package distance_thresholds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new distance thresholds API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for distance thresholds API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDistanceThreshold(params *CreateDistanceThresholdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDistanceThresholdOK, error)

	DeleteDistanceThreshold(params *DeleteDistanceThresholdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDistanceThresholdOK, error)

	GetDistanceThreshold(params *GetDistanceThresholdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDistanceThresholdOK, error)

	ListDistanceThresholds(params *ListDistanceThresholdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDistanceThresholdsOK, error)

	QueryDistanceThresholds(params *QueryDistanceThresholdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryDistanceThresholdsOK, error)

	UpdateDistanceThreshold(params *UpdateDistanceThresholdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDistanceThresholdOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDistanceThreshold creates one or more distance threshold objects

  Create one or more DistanceThreshold objects for this company.

A company-distance-threshold model indicates the distance between a company
and the taxing borders of various countries.  Distance thresholds are necessary
to correctly calculate some value-added taxes.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, CompanyAdmin, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) CreateDistanceThreshold(params *CreateDistanceThresholdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDistanceThresholdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDistanceThresholdParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateDistanceThreshold",
		Method:             "POST",
		PathPattern:        "/api/v2/companies/{companyId}/distancethresholds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDistanceThresholdReader{formats: a.formats},
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
	success, ok := result.(*CreateDistanceThresholdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDistanceThreshold: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDistanceThreshold deletes a single distance threshold object

  Marks the DistanceThreshold object identified by this URL as deleted.

A company-distance-threshold model indicates the distance between a company
and the taxing borders of various countries.  Distance thresholds are necessary
to correctly calculate some value-added taxes.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, CompanyAdmin, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) DeleteDistanceThreshold(params *DeleteDistanceThresholdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDistanceThresholdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDistanceThresholdParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDistanceThreshold",
		Method:             "DELETE",
		PathPattern:        "/api/v2/companies/{companyId}/distancethresholds/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDistanceThresholdReader{formats: a.formats},
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
	success, ok := result.(*DeleteDistanceThresholdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteDistanceThreshold: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDistanceThreshold retrieves a single distance threshold

  Retrieves a single DistanceThreshold object defined by this URL.

A company-distance-threshold model indicates the distance between a company
and the taxing borders of various countries.  Distance thresholds are necessary
to correctly calculate some value-added taxes.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountUser, CompanyAdmin, CompanyUser, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser, TreasuryAdmin, TreasuryUser.

*/
func (a *Client) GetDistanceThreshold(params *GetDistanceThresholdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDistanceThresholdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistanceThresholdParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDistanceThreshold",
		Method:             "GET",
		PathPattern:        "/api/v2/companies/{companyId}/distancethresholds/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDistanceThresholdReader{formats: a.formats},
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
	success, ok := result.(*GetDistanceThresholdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDistanceThreshold: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDistanceThresholds retrieves all distance thresholds for this company

  Lists all DistanceThreshold objects that belong to this company.

A company-distance-threshold model indicates the distance between a company
and the taxing borders of various countries.  Distance thresholds are necessary
to correctly calculate some value-added taxes.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountUser, CompanyAdmin, CompanyUser, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser, TreasuryAdmin, TreasuryUser.

*/
func (a *Client) ListDistanceThresholds(params *ListDistanceThresholdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDistanceThresholdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDistanceThresholdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListDistanceThresholds",
		Method:             "GET",
		PathPattern:        "/api/v2/companies/{companyId}/distancethresholds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDistanceThresholdsReader{formats: a.formats},
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
	success, ok := result.(*ListDistanceThresholdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListDistanceThresholds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryDistanceThresholds retrieves all distance threshold objects

  Lists all DistanceThreshold objects that belong to this account.

A company-distance-threshold model indicates the distance between a company
and the taxing borders of various countries.  Distance thresholds are necessary
to correctly calculate some value-added taxes.

Search for specific objects using the criteria in the `$filter` parameter; full documentation is available on [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/) .
Paginate your results using the `$top`, `$skip`, and `$orderby` parameters.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountUser, CompanyAdmin, CompanyUser, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser, TreasuryAdmin, TreasuryUser.

*/
func (a *Client) QueryDistanceThresholds(params *QueryDistanceThresholdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryDistanceThresholdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDistanceThresholdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryDistanceThresholds",
		Method:             "GET",
		PathPattern:        "/api/v2/distancethresholds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDistanceThresholdsReader{formats: a.formats},
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
	success, ok := result.(*QueryDistanceThresholdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryDistanceThresholds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDistanceThreshold updates a distance threshold object

  Replace the existing DistanceThreshold object at this URL with an updated object.

A company-distance-threshold model indicates the distance between a company
and the taxing borders of various countries.  Distance thresholds are necessary
to correctly calculate some value-added taxes.

All data from the existing object will be replaced with data in the object you PUT.
To set a field's value to null, you may either set its value to null or omit that field from the object you post.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, CompanyAdmin, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) UpdateDistanceThreshold(params *UpdateDistanceThresholdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDistanceThresholdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDistanceThresholdParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateDistanceThreshold",
		Method:             "PUT",
		PathPattern:        "/api/v2/companies/{companyId}/distancethresholds/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDistanceThresholdReader{formats: a.formats},
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
	success, ok := result.(*UpdateDistanceThresholdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateDistanceThreshold: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
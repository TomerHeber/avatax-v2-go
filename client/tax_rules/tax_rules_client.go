// Code generated by go-swagger; DO NOT EDIT.

package tax_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tax rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tax rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTaxRules(params *CreateTaxRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTaxRulesOK, error)

	DeleteTaxRule(params *DeleteTaxRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTaxRuleOK, error)

	GetTaxRule(params *GetTaxRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTaxRuleOK, error)

	ListTaxRules(params *ListTaxRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTaxRulesOK, error)

	QueryTaxRules(params *QueryTaxRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryTaxRulesOK, error)

	UpdateTaxRule(params *UpdateTaxRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTaxRuleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTaxRules creates a new tax rule

  Create one or more custom tax rules attached to this company.

A tax rule represents a rule that changes the default AvaTax behavior for a product or jurisdiction.  Custom tax rules
can be used to change the taxability of an item, to change the tax base of an item, or to change the tax rate
charged when selling an item.  Tax rules can also change tax behavior depending on the `entityUseCode` value submitted
with the transaction.

You can create custom tax rules to customize the behavior of AvaTax to match specific rules that are custom to your
business.  If you have obtained a ruling from a tax auditor that requires custom tax calculations, you can use
custom tax rules to redefine the behavior for your company or item.

Please use custom tax rules carefully and ensure that these tax rules match the behavior agreed upon with your
auditor, legal representative, and accounting team.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, CompanyAdmin, CSPTester, SSTAdmin, TechnicalSupportAdmin.

*/
func (a *Client) CreateTaxRules(params *CreateTaxRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTaxRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTaxRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTaxRules",
		Method:             "POST",
		PathPattern:        "/api/v2/companies/{companyId}/taxrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTaxRulesReader{formats: a.formats},
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
	success, ok := result.(*CreateTaxRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateTaxRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTaxRule deletes a single tax rule

  Mark the custom tax rule identified by this URL as deleted.

A tax rule represents a rule that changes the default AvaTax behavior for a product or jurisdiction.  Custom tax rules
can be used to change the taxability of an item, to change the tax base of an item, or to change the tax rate
charged when selling an item.  Tax rules can also change tax behavior depending on the `entityUseCode` value submitted
with the transaction.

You can create custom tax rules to customize the behavior of AvaTax to match specific rules that are custom to your
business.  If you have obtained a ruling from a tax auditor that requires custom tax calculations, you can use
custom tax rules to redefine the behavior for your company or item.

Please use custom tax rules carefully and ensure that these tax rules match the behavior agreed upon with your
auditor, legal representative, and accounting team.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, CompanyAdmin, CSPTester, SSTAdmin, TechnicalSupportAdmin.

*/
func (a *Client) DeleteTaxRule(params *DeleteTaxRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTaxRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTaxRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTaxRule",
		Method:             "DELETE",
		PathPattern:        "/api/v2/companies/{companyId}/taxrules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTaxRuleReader{formats: a.formats},
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
	success, ok := result.(*DeleteTaxRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteTaxRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTaxRule retrieves a single tax rule

  Get the taxrule object identified by this URL.

A tax rule represents a rule that changes the default AvaTax behavior for a product or jurisdiction.  Custom tax rules
can be used to change the taxability of an item, to change the tax base of an item, or to change the tax rate
charged when selling an item.  Tax rules can also change tax behavior depending on the `entityUseCode` value submitted
with the transaction.

You can create custom tax rules to customize the behavior of AvaTax to match specific rules that are custom to your
business.  If you have obtained a ruling from a tax auditor that requires custom tax calculations, you can use
custom tax rules to redefine the behavior for your company or item.

Please use custom tax rules carefully and ensure that these tax rules match the behavior agreed upon with your
auditor, legal representative, and accounting team.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountUser, CompanyAdmin, CompanyUser, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser.

*/
func (a *Client) GetTaxRule(params *GetTaxRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTaxRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaxRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTaxRule",
		Method:             "GET",
		PathPattern:        "/api/v2/companies/{companyId}/taxrules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTaxRuleReader{formats: a.formats},
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
	success, ok := result.(*GetTaxRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTaxRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListTaxRules retrieves tax rules for this company

  List all taxrule objects attached to this company.

A tax rule represents a rule that changes the default AvaTax behavior for a product or jurisdiction.  Custom tax rules
can be used to change the taxability of an item, to change the tax base of an item, or to change the tax rate
charged when selling an item.  Tax rules can also change tax behavior depending on the `entityUseCode` value submitted
with the transaction.

You can create custom tax rules to customize the behavior of AvaTax to match specific rules that are custom to your
business.  If you have obtained a ruling from a tax auditor that requires custom tax calculations, you can use
custom tax rules to redefine the behavior for your company or item.

Please use custom tax rules carefully and ensure that these tax rules match the behavior agreed upon with your
auditor, legal representative, and accounting team.

Search for specific objects using the criteria in the `$filter` parameter; full documentation is available on [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/) .
Paginate your results using the `$top`, `$skip`, and `$orderby` parameters.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountUser, CompanyAdmin, CompanyUser, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser.

*/
func (a *Client) ListTaxRules(params *ListTaxRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTaxRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTaxRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListTaxRules",
		Method:             "GET",
		PathPattern:        "/api/v2/companies/{companyId}/taxrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTaxRulesReader{formats: a.formats},
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
	success, ok := result.(*ListTaxRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListTaxRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryTaxRules retrieves all tax rules

  Get multiple taxrule objects across all companies.

A tax rule represents a rule that changes the default AvaTax behavior for a product or jurisdiction.  Custom tax rules
can be used to change the taxability of an item, to change the tax base of an item, or to change the tax rate
charged when selling an item.  Tax rules can also change tax behavior depending on the `entityUseCode` value submitted
with the transaction.

You can create custom tax rules to customize the behavior of AvaTax to match specific rules that are custom to your
business.  If you have obtained a ruling from a tax auditor that requires custom tax calculations, you can use
custom tax rules to redefine the behavior for your company or item.

Please use custom tax rules carefully and ensure that these tax rules match the behavior agreed upon with your
auditor, legal representative, and accounting team.

Search for specific objects using the criteria in the `$filter` parameter; full documentation is available on [Filtering in REST](http://developer.avalara.com/avatax/filtering-in-rest/) .
Paginate your results using the `$top`, `$skip`, and `$orderby` parameters.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountUser, CompanyAdmin, CompanyUser, CSPAdmin, CSPTester, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser.

*/
func (a *Client) QueryTaxRules(params *QueryTaxRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryTaxRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryTaxRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryTaxRules",
		Method:             "GET",
		PathPattern:        "/api/v2/taxrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryTaxRulesReader{formats: a.formats},
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
	success, ok := result.(*QueryTaxRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryTaxRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTaxRule updates a single tax rule

  Replace the existing custom tax rule object at this URL with an updated object.

A tax rule represents a rule that changes the default AvaTax behavior for a product or jurisdiction.  Custom tax rules
can be used to change the taxability of an item, to change the tax base of an item, or to change the tax rate
charged when selling an item.  Tax rules can also change tax behavior depending on the `entityUseCode` value submitted
with the transaction.

You can create custom tax rules to customize the behavior of AvaTax to match specific rules that are custom to your
business.  If you have obtained a ruling from a tax auditor that requires custom tax calculations, you can use
custom tax rules to redefine the behavior for your company or item.

Please use custom tax rules carefully and ensure that these tax rules match the behavior agreed upon with your
auditor, legal representative, and accounting team.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, CompanyAdmin, CSPTester, SSTAdmin, TechnicalSupportAdmin.

*/
func (a *Client) UpdateTaxRule(params *UpdateTaxRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTaxRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTaxRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTaxRule",
		Method:             "PUT",
		PathPattern:        "/api/v2/companies/{companyId}/taxrules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTaxRuleReader{formats: a.formats},
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
	success, ok := result.(*UpdateTaxRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateTaxRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
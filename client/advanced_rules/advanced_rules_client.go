// Code generated by go-swagger; DO NOT EDIT.

package advanced_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new advanced rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for advanced rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCompanyLookupFile(params *CreateCompanyLookupFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCompanyLookupFileCreated, error)

	DeleteLookupFile(params *DeleteLookupFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteLookupFileOK, error)

	GetCompanyLookupFiles(params *GetCompanyLookupFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCompanyLookupFilesOK, error)

	GetLookupFile(params *GetLookupFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLookupFileOK, error)

	UpdateLookupFile(params *UpdateLookupFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLookupFileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCompanyLookupFile creates a lookup file for a company
*/
func (a *Client) CreateCompanyLookupFile(params *CreateCompanyLookupFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCompanyLookupFileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCompanyLookupFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCompanyLookupFile",
		Method:             "POST",
		PathPattern:        "/api/v2/advancedrules/accounts/{accountId}/companies/{companyId}/lookupFiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCompanyLookupFileReader{formats: a.formats},
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
	success, ok := result.(*CreateCompanyLookupFileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateCompanyLookupFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteLookupFile deletes a lookup file
*/
func (a *Client) DeleteLookupFile(params *DeleteLookupFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteLookupFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLookupFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteLookupFile",
		Method:             "DELETE",
		PathPattern:        "/api/v2/advancedrules/accounts/{accountId}/lookupFiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLookupFileReader{formats: a.formats},
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
	success, ok := result.(*DeleteLookupFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteLookupFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCompanyLookupFiles gets the lookup files for a company
*/
func (a *Client) GetCompanyLookupFiles(params *GetCompanyLookupFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCompanyLookupFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCompanyLookupFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCompanyLookupFiles",
		Method:             "GET",
		PathPattern:        "/api/v2/advancedrules/accounts/{accountId}/companies/{companyId}/lookupFiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCompanyLookupFilesReader{formats: a.formats},
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
	success, ok := result.(*GetCompanyLookupFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCompanyLookupFiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLookupFile gets a lookup file for an account Id and company lookup file Id
*/
func (a *Client) GetLookupFile(params *GetLookupFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLookupFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLookupFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLookupFile",
		Method:             "GET",
		PathPattern:        "/api/v2/advancedrules/accounts/{accountId}/lookupFiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLookupFileReader{formats: a.formats},
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
	success, ok := result.(*GetLookupFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLookupFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateLookupFile updates a lookup file
*/
func (a *Client) UpdateLookupFile(params *UpdateLookupFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLookupFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLookupFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateLookupFile",
		Method:             "PUT",
		PathPattern:        "/api/v2/advancedrules/accounts/{accountId}/lookupFiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLookupFileReader{formats: a.formats},
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
	success, ok := result.(*UpdateLookupFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateLookupFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
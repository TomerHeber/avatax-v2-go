// Code generated by go-swagger; DO NOT EDIT.

package registrar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new registrar API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for registrar API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAccount(params *CreateAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccountOK, error)

	CreateNotifications(params *CreateNotificationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNotificationsOK, error)

	CreateSubscriptions(params *CreateSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSubscriptionsOK, error)

	DeleteAccount(params *DeleteAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAccountOK, error)

	DeleteNotification(params *DeleteNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNotificationOK, error)

	DeleteSubscription(params *DeleteSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSubscriptionOK, error)

	ListAccountsByTssWriteMode(params *ListAccountsByTssWriteModeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAccountsByTssWriteModeOK, error)

	ResetPassword(params *ResetPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetPasswordOK, error)

	UpdateAccount(params *UpdateAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccountOK, error)

	UpdateNotification(params *UpdateNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNotificationOK, error)

	UpdateSubscription(params *UpdateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSubscriptionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAccount creates a new account

  # For Registrar Use Only
This API is for use by Avalara Registrar administrative users only.

Create a single new account object.
When creating an account object you may attach subscriptions and users as part of the 'Create' call.

### Security Policies

* This API requires one of the following user roles: FirmAdmin, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) CreateAccount(params *CreateAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateAccount",
		Method:             "POST",
		PathPattern:        "/api/v2/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAccountReader{formats: a.formats},
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
	success, ok := result.(*CreateAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateNotifications creates new notifications

  This API is available by invitation only.

Create a single notification.

A notification is a message from Avalara that may have relevance to your business.  You may want
to regularly review notifications and then dismiss them when you are certain that you have addressed
any relevant concerns raised by this notification.

A Global notification is a message which is directed to all the accounts and is set to expire within
a certain time and cannot be dismissed by the user. Make accountId and companyId null to create a global notification.

An example of a notification would be a message about new software, or a change to AvaTax that may
affect you, or a potential issue with your company's tax profile.

### Security Policies

* This API requires one of the following user roles: FirmAdmin, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.
* This API is available by invitation only.  To request access to this feature, please speak to a business development manager and request access to [NotificationsAPI:Create].

*/
func (a *Client) CreateNotifications(params *CreateNotificationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNotificationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateNotifications",
		Method:             "POST",
		PathPattern:        "/api/v2/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNotificationsReader{formats: a.formats},
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
	success, ok := result.(*CreateNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateNotifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateSubscriptions creates a new subscription

  This API is for use by Avalara Registrar administrative users only.

Create one or more new subscription objects attached to this account.
A 'subscription' indicates a licensed subscription to a named Avalara service.
To request or remove subscriptions, please contact Avalara sales or your customer account manager.

### Security Policies

* This API requires one of the following user roles: Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) CreateSubscriptions(params *CreateSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSubscriptions",
		Method:             "POST",
		PathPattern:        "/api/v2/accounts/{accountId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSubscriptionsReader{formats: a.formats},
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
	success, ok := result.(*CreateSubscriptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateSubscriptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAccount deletes a single account

  # For Registrar Use Only
This API is for use by Avalara Registrar administrative users only.

Delete an account.
Deleting an account will delete all companies, all account level users and license keys attached to this account.

### Security Policies

* This API requires the user role SystemAdmin.

*/
func (a *Client) DeleteAccount(params *DeleteAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAccount",
		Method:             "DELETE",
		PathPattern:        "/api/v2/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccountReader{formats: a.formats},
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
	success, ok := result.(*DeleteAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNotification deletes a single notification

  This API is available by invitation only.

Delete the existing notification identified by this URL.

A notification is a message from Avalara that may have relevance to your business.  You may want
to regularly review notifications and then dismiss them when you are certain that you have addressed
any relevant concerns raised by this notification.

An example of a notification would be a message about new software, or a change to AvaTax that may
affect you, or a potential issue with your company's tax profile.

### Security Policies

* This API requires one of the following user roles: FirmAdmin, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.
* This API is available by invitation only.  To request access to this feature, please speak to a business development manager and request access to [NotificationsAPI:Create].

*/
func (a *Client) DeleteNotification(params *DeleteNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNotification",
		Method:             "DELETE",
		PathPattern:        "/api/v2/notifications/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNotificationReader{formats: a.formats},
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
	success, ok := result.(*DeleteNotificationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSubscription deletes a single subscription

  # For Registrar Use Only
This API is for use by Avalara Registrar administrative users only.

Mark the existing account identified by this URL as deleted.

### Security Policies

* This API requires one of the following user roles: Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) DeleteSubscription(params *DeleteSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSubscriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSubscription",
		Method:             "DELETE",
		PathPattern:        "/api/v2/accounts/{accountId}/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSubscriptionReader{formats: a.formats},
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
	success, ok := result.(*DeleteSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAccountsByTssWriteMode retrieves list of accounts by account migration status



### Security Policies

* This API requires one of the following user roles: FirmAdmin, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) ListAccountsByTssWriteMode(params *ListAccountsByTssWriteModeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAccountsByTssWriteModeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAccountsByTssWriteModeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListAccountsByTssWriteMode",
		Method:             "GET",
		PathPattern:        "/api/v2/accounts/ListAccountsByTssWriteMode/{writeMode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAccountsByTssWriteModeReader{formats: a.formats},
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
	success, ok := result.(*ListAccountsByTssWriteModeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListAccountsByTssWriteMode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResetPassword resets a user s password programmatically

  # For Registrar Use Only
This API is for use by Avalara Registrar administrative users only.

Allows a system admin to reset the password for a specific user via the API.
This API is only available for Avalara Registrar Admins, and can be used to reset the password of any
user based on internal Avalara business processes.

### Security Policies

* This API requires one of the following user roles: AccountAdmin, AccountUser, CompanyAdmin, CompanyUser, Compliance Root User, ComplianceAdmin, ComplianceUser, CSPTester, FirmAdmin, FirmUser, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin, TechnicalSupportUser, TreasuryAdmin, TreasuryUser.
* This API is available to Avalara system-level (registrar-level) users only.

*/
func (a *Client) ResetPassword(params *ResetPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResetPassword",
		Method:             "POST",
		PathPattern:        "/api/v2/passwords/{userId}/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetPasswordReader{formats: a.formats},
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
	success, ok := result.(*ResetPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResetPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAccount updates a single account

  # For Registrar Use Only
This API is for use by Avalara Registrar administrative users only.

Replace an existing account object with an updated account object.

### Security Policies

* This API requires one of the following user roles: FirmAdmin, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) UpdateAccount(params *UpdateAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateAccount",
		Method:             "PUT",
		PathPattern:        "/api/v2/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAccountReader{formats: a.formats},
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
	success, ok := result.(*UpdateAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNotification updates a single notification

  This API is available by invitation only.

Replaces the notification identified by this URL with a new notification.

A notification is a message from Avalara that may have relevance to your business.  You may want
to regularly review notifications and then dismiss them when you are certain that you have addressed
any relevant concerns raised by this notification.

An example of a notification would be a message about new software, or a change to AvaTax that may
affect you, or a potential issue with your company's tax profile.

### Security Policies

* This API requires one of the following user roles: FirmAdmin, Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.
* This API is available by invitation only.  To request access to this feature, please speak to a business development manager and request access to [NotificationsAPI:Create].

*/
func (a *Client) UpdateNotification(params *UpdateNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateNotification",
		Method:             "PUT",
		PathPattern:        "/api/v2/notifications/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNotificationReader{formats: a.formats},
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
	success, ok := result.(*UpdateNotificationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSubscription updates a single subscription

  # For Registrar Use Only
This API is for use by Avalara Registrar administrative users only.

Replace the existing subscription object at this URL with an updated object.
A 'subscription' indicates a licensed subscription to a named Avalara service.
To request or remove subscriptions, please contact Avalara sales or your customer account manager.
All data from the existing object will be replaced with data in the object you PUT.
To set a field's value to null, you may either set its value to null or omit that field from the object you post.

### Security Policies

* This API requires one of the following user roles: Registrar, SiteAdmin, SSTAdmin, SystemAdmin, TechnicalSupportAdmin.

*/
func (a *Client) UpdateSubscription(params *UpdateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSubscriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSubscription",
		Method:             "PUT",
		PathPattern:        "/api/v2/accounts/{accountId}/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSubscriptionReader{formats: a.formats},
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
	success, ok := result.(*UpdateSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUserParams creates a new GetUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserParams() *GetUserParams {
	return &GetUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserParamsWithTimeout creates a new GetUserParams object
// with the ability to set a timeout on a request.
func NewGetUserParamsWithTimeout(timeout time.Duration) *GetUserParams {
	return &GetUserParams{
		timeout: timeout,
	}
}

// NewGetUserParamsWithContext creates a new GetUserParams object
// with the ability to set a context for a request.
func NewGetUserParamsWithContext(ctx context.Context) *GetUserParams {
	return &GetUserParams{
		Context: ctx,
	}
}

// NewGetUserParamsWithHTTPClient creates a new GetUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserParamsWithHTTPClient(client *http.Client) *GetUserParams {
	return &GetUserParams{
		HTTPClient: client,
	}
}

/* GetUserParams contains all the parameters to send to the API endpoint
   for the get user operation.

   Typically these are written to a http.Request.
*/
type GetUserParams struct {

	/* DollarInclude.

	   Optional fetch commands.
	*/
	DollarInclude *string

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* AccountID.

	   The accountID of the user you wish to get.

	   Format: int32
	*/
	AccountID int32

	/* ID.

	   The ID of the user to retrieve.

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserParams) WithDefaults() *GetUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := GetUserParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get user params
func (o *GetUserParams) WithTimeout(timeout time.Duration) *GetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user params
func (o *GetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user params
func (o *GetUserParams) WithContext(ctx context.Context) *GetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user params
func (o *GetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) WithHTTPClient(client *http.Client) *GetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarInclude adds the dollarInclude to the get user params
func (o *GetUserParams) WithDollarInclude(dollarInclude *string) *GetUserParams {
	o.SetDollarInclude(dollarInclude)
	return o
}

// SetDollarInclude adds the dollarInclude to the get user params
func (o *GetUserParams) SetDollarInclude(dollarInclude *string) {
	o.DollarInclude = dollarInclude
}

// WithXAvalaraClient adds the xAvalaraClient to the get user params
func (o *GetUserParams) WithXAvalaraClient(xAvalaraClient *string) *GetUserParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the get user params
func (o *GetUserParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithAccountID adds the accountID to the get user params
func (o *GetUserParams) WithAccountID(accountID int32) *GetUserParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get user params
func (o *GetUserParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithID adds the id to the get user params
func (o *GetUserParams) WithID(id int32) *GetUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get user params
func (o *GetUserParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.XAvalaraClient != nil {

		// header param X-Avalara-Client
		if err := r.SetHeaderParam("X-Avalara-Client", *o.XAvalaraClient); err != nil {
			return err
		}
	}

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

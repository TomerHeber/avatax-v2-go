// Code generated by go-swagger; DO NOT EDIT.

package companies

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

	"github.com/TomerHeber/avatax-v2-go/models"
)

// NewCreateFundingRequestParams creates a new CreateFundingRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFundingRequestParams() *CreateFundingRequestParams {
	return &CreateFundingRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFundingRequestParamsWithTimeout creates a new CreateFundingRequestParams object
// with the ability to set a timeout on a request.
func NewCreateFundingRequestParamsWithTimeout(timeout time.Duration) *CreateFundingRequestParams {
	return &CreateFundingRequestParams{
		timeout: timeout,
	}
}

// NewCreateFundingRequestParamsWithContext creates a new CreateFundingRequestParams object
// with the ability to set a context for a request.
func NewCreateFundingRequestParamsWithContext(ctx context.Context) *CreateFundingRequestParams {
	return &CreateFundingRequestParams{
		Context: ctx,
	}
}

// NewCreateFundingRequestParamsWithHTTPClient creates a new CreateFundingRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFundingRequestParamsWithHTTPClient(client *http.Client) *CreateFundingRequestParams {
	return &CreateFundingRequestParams{
		HTTPClient: client,
	}
}

/* CreateFundingRequestParams contains all the parameters to send to the API endpoint
   for the create funding request operation.

   Typically these are written to a http.Request.
*/
type CreateFundingRequestParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The funding initialization request
	*/
	Body *models.FundingInitiateModel

	/* ID.

	   The unique identifier of the company

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create funding request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFundingRequestParams) WithDefaults() *CreateFundingRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create funding request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFundingRequestParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := CreateFundingRequestParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create funding request params
func (o *CreateFundingRequestParams) WithTimeout(timeout time.Duration) *CreateFundingRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create funding request params
func (o *CreateFundingRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create funding request params
func (o *CreateFundingRequestParams) WithContext(ctx context.Context) *CreateFundingRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create funding request params
func (o *CreateFundingRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create funding request params
func (o *CreateFundingRequestParams) WithHTTPClient(client *http.Client) *CreateFundingRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create funding request params
func (o *CreateFundingRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the create funding request params
func (o *CreateFundingRequestParams) WithXAvalaraClient(xAvalaraClient *string) *CreateFundingRequestParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the create funding request params
func (o *CreateFundingRequestParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the create funding request params
func (o *CreateFundingRequestParams) WithBody(body *models.FundingInitiateModel) *CreateFundingRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create funding request params
func (o *CreateFundingRequestParams) SetBody(body *models.FundingInitiateModel) {
	o.Body = body
}

// WithID adds the id to the create funding request params
func (o *CreateFundingRequestParams) WithID(id int32) *CreateFundingRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create funding request params
func (o *CreateFundingRequestParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFundingRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
// Code generated by go-swagger; DO NOT EDIT.

package batches

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

// NewCreateBatchesParams creates a new CreateBatchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBatchesParams() *CreateBatchesParams {
	return &CreateBatchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBatchesParamsWithTimeout creates a new CreateBatchesParams object
// with the ability to set a timeout on a request.
func NewCreateBatchesParamsWithTimeout(timeout time.Duration) *CreateBatchesParams {
	return &CreateBatchesParams{
		timeout: timeout,
	}
}

// NewCreateBatchesParamsWithContext creates a new CreateBatchesParams object
// with the ability to set a context for a request.
func NewCreateBatchesParamsWithContext(ctx context.Context) *CreateBatchesParams {
	return &CreateBatchesParams{
		Context: ctx,
	}
}

// NewCreateBatchesParamsWithHTTPClient creates a new CreateBatchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBatchesParamsWithHTTPClient(client *http.Client) *CreateBatchesParams {
	return &CreateBatchesParams{
		HTTPClient: client,
	}
}

/* CreateBatchesParams contains all the parameters to send to the API endpoint
   for the create batches operation.

   Typically these are written to a http.Request.
*/
type CreateBatchesParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The batch you wish to create.
	*/
	Body []*models.BatchModel

	/* CompanyID.

	   The ID of the company that owns this batch.

	   Format: int32
	*/
	CompanyID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create batches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBatchesParams) WithDefaults() *CreateBatchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create batches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBatchesParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := CreateBatchesParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create batches params
func (o *CreateBatchesParams) WithTimeout(timeout time.Duration) *CreateBatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create batches params
func (o *CreateBatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create batches params
func (o *CreateBatchesParams) WithContext(ctx context.Context) *CreateBatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create batches params
func (o *CreateBatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create batches params
func (o *CreateBatchesParams) WithHTTPClient(client *http.Client) *CreateBatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create batches params
func (o *CreateBatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the create batches params
func (o *CreateBatchesParams) WithXAvalaraClient(xAvalaraClient *string) *CreateBatchesParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the create batches params
func (o *CreateBatchesParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the create batches params
func (o *CreateBatchesParams) WithBody(body []*models.BatchModel) *CreateBatchesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create batches params
func (o *CreateBatchesParams) SetBody(body []*models.BatchModel) {
	o.Body = body
}

// WithCompanyID adds the companyID to the create batches params
func (o *CreateBatchesParams) WithCompanyID(companyID int32) *CreateBatchesParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the create batches params
func (o *CreateBatchesParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param companyId
	if err := r.SetPathParam("companyId", swag.FormatInt32(o.CompanyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

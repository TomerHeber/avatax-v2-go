// Code generated by go-swagger; DO NOT EDIT.

package user_defined_fields

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

// NewUpdateUserDefinedFieldParams creates a new UpdateUserDefinedFieldParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserDefinedFieldParams() *UpdateUserDefinedFieldParams {
	return &UpdateUserDefinedFieldParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserDefinedFieldParamsWithTimeout creates a new UpdateUserDefinedFieldParams object
// with the ability to set a timeout on a request.
func NewUpdateUserDefinedFieldParamsWithTimeout(timeout time.Duration) *UpdateUserDefinedFieldParams {
	return &UpdateUserDefinedFieldParams{
		timeout: timeout,
	}
}

// NewUpdateUserDefinedFieldParamsWithContext creates a new UpdateUserDefinedFieldParams object
// with the ability to set a context for a request.
func NewUpdateUserDefinedFieldParamsWithContext(ctx context.Context) *UpdateUserDefinedFieldParams {
	return &UpdateUserDefinedFieldParams{
		Context: ctx,
	}
}

// NewUpdateUserDefinedFieldParamsWithHTTPClient creates a new UpdateUserDefinedFieldParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserDefinedFieldParamsWithHTTPClient(client *http.Client) *UpdateUserDefinedFieldParams {
	return &UpdateUserDefinedFieldParams{
		HTTPClient: client,
	}
}

/* UpdateUserDefinedFieldParams contains all the parameters to send to the API endpoint
   for the update user defined field operation.

   Typically these are written to a http.Request.
*/
type UpdateUserDefinedFieldParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	// Body.
	Body *models.CompanyUserDefinedFieldModel

	/* CompanyID.

	   The id of the company the user defined field belongs to.

	   Format: int32
	*/
	CompanyID int32

	// ID.
	//
	// Format: int64
	ID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user defined field params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserDefinedFieldParams) WithDefaults() *UpdateUserDefinedFieldParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user defined field params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserDefinedFieldParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := UpdateUserDefinedFieldParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update user defined field params
func (o *UpdateUserDefinedFieldParams) WithTimeout(timeout time.Duration) *UpdateUserDefinedFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user defined field params
func (o *UpdateUserDefinedFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user defined field params
func (o *UpdateUserDefinedFieldParams) WithContext(ctx context.Context) *UpdateUserDefinedFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user defined field params
func (o *UpdateUserDefinedFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user defined field params
func (o *UpdateUserDefinedFieldParams) WithHTTPClient(client *http.Client) *UpdateUserDefinedFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user defined field params
func (o *UpdateUserDefinedFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the update user defined field params
func (o *UpdateUserDefinedFieldParams) WithXAvalaraClient(xAvalaraClient *string) *UpdateUserDefinedFieldParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the update user defined field params
func (o *UpdateUserDefinedFieldParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the update user defined field params
func (o *UpdateUserDefinedFieldParams) WithBody(body *models.CompanyUserDefinedFieldModel) *UpdateUserDefinedFieldParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user defined field params
func (o *UpdateUserDefinedFieldParams) SetBody(body *models.CompanyUserDefinedFieldModel) {
	o.Body = body
}

// WithCompanyID adds the companyID to the update user defined field params
func (o *UpdateUserDefinedFieldParams) WithCompanyID(companyID int32) *UpdateUserDefinedFieldParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the update user defined field params
func (o *UpdateUserDefinedFieldParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the update user defined field params
func (o *UpdateUserDefinedFieldParams) WithID(id *int64) *UpdateUserDefinedFieldParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update user defined field params
func (o *UpdateUserDefinedFieldParams) SetID(id *int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserDefinedFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ID != nil {

		// query param id
		var qrID int64

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
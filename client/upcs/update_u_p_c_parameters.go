// Code generated by go-swagger; DO NOT EDIT.

package upcs

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

// NewUpdateUPCParams creates a new UpdateUPCParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUPCParams() *UpdateUPCParams {
	return &UpdateUPCParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUPCParamsWithTimeout creates a new UpdateUPCParams object
// with the ability to set a timeout on a request.
func NewUpdateUPCParamsWithTimeout(timeout time.Duration) *UpdateUPCParams {
	return &UpdateUPCParams{
		timeout: timeout,
	}
}

// NewUpdateUPCParamsWithContext creates a new UpdateUPCParams object
// with the ability to set a context for a request.
func NewUpdateUPCParamsWithContext(ctx context.Context) *UpdateUPCParams {
	return &UpdateUPCParams{
		Context: ctx,
	}
}

// NewUpdateUPCParamsWithHTTPClient creates a new UpdateUPCParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUPCParamsWithHTTPClient(client *http.Client) *UpdateUPCParams {
	return &UpdateUPCParams{
		HTTPClient: client,
	}
}

/* UpdateUPCParams contains all the parameters to send to the API endpoint
   for the update u p c operation.

   Typically these are written to a http.Request.
*/
type UpdateUPCParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* Body.

	   The UPC you wish to update.
	*/
	Body *models.UPCModel

	/* CompanyID.

	   The ID of the company that this UPC belongs to.

	   Format: int32
	*/
	CompanyID int32

	/* ID.

	   The ID of the UPC you wish to update

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update u p c params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUPCParams) WithDefaults() *UpdateUPCParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update u p c params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUPCParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := UpdateUPCParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update u p c params
func (o *UpdateUPCParams) WithTimeout(timeout time.Duration) *UpdateUPCParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update u p c params
func (o *UpdateUPCParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update u p c params
func (o *UpdateUPCParams) WithContext(ctx context.Context) *UpdateUPCParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update u p c params
func (o *UpdateUPCParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update u p c params
func (o *UpdateUPCParams) WithHTTPClient(client *http.Client) *UpdateUPCParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update u p c params
func (o *UpdateUPCParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the update u p c params
func (o *UpdateUPCParams) WithXAvalaraClient(xAvalaraClient *string) *UpdateUPCParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the update u p c params
func (o *UpdateUPCParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBody adds the body to the update u p c params
func (o *UpdateUPCParams) WithBody(body *models.UPCModel) *UpdateUPCParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update u p c params
func (o *UpdateUPCParams) SetBody(body *models.UPCModel) {
	o.Body = body
}

// WithCompanyID adds the companyID to the update u p c params
func (o *UpdateUPCParams) WithCompanyID(companyID int32) *UpdateUPCParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the update u p c params
func (o *UpdateUPCParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the update u p c params
func (o *UpdateUPCParams) WithID(id int32) *UpdateUPCParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update u p c params
func (o *UpdateUPCParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUPCParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
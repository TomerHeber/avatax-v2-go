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
)

// NewDownloadBatchParams creates a new DownloadBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadBatchParams() *DownloadBatchParams {
	return &DownloadBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadBatchParamsWithTimeout creates a new DownloadBatchParams object
// with the ability to set a timeout on a request.
func NewDownloadBatchParamsWithTimeout(timeout time.Duration) *DownloadBatchParams {
	return &DownloadBatchParams{
		timeout: timeout,
	}
}

// NewDownloadBatchParamsWithContext creates a new DownloadBatchParams object
// with the ability to set a context for a request.
func NewDownloadBatchParamsWithContext(ctx context.Context) *DownloadBatchParams {
	return &DownloadBatchParams{
		Context: ctx,
	}
}

// NewDownloadBatchParamsWithHTTPClient creates a new DownloadBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadBatchParamsWithHTTPClient(client *http.Client) *DownloadBatchParams {
	return &DownloadBatchParams{
		HTTPClient: client,
	}
}

/* DownloadBatchParams contains all the parameters to send to the API endpoint
   for the download batch operation.

   Typically these are written to a http.Request.
*/
type DownloadBatchParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* BatchID.

	   The ID of the batch object

	   Format: int32
	*/
	BatchID int32

	/* CompanyID.

	   The ID of the company that owns this batch

	   Format: int32
	*/
	CompanyID int32

	/* ID.

	   The primary key of this batch file object

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadBatchParams) WithDefaults() *DownloadBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadBatchParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := DownloadBatchParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the download batch params
func (o *DownloadBatchParams) WithTimeout(timeout time.Duration) *DownloadBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download batch params
func (o *DownloadBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download batch params
func (o *DownloadBatchParams) WithContext(ctx context.Context) *DownloadBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download batch params
func (o *DownloadBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download batch params
func (o *DownloadBatchParams) WithHTTPClient(client *http.Client) *DownloadBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download batch params
func (o *DownloadBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the download batch params
func (o *DownloadBatchParams) WithXAvalaraClient(xAvalaraClient *string) *DownloadBatchParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the download batch params
func (o *DownloadBatchParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithBatchID adds the batchID to the download batch params
func (o *DownloadBatchParams) WithBatchID(batchID int32) *DownloadBatchParams {
	o.SetBatchID(batchID)
	return o
}

// SetBatchID adds the batchId to the download batch params
func (o *DownloadBatchParams) SetBatchID(batchID int32) {
	o.BatchID = batchID
}

// WithCompanyID adds the companyID to the download batch params
func (o *DownloadBatchParams) WithCompanyID(companyID int32) *DownloadBatchParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the download batch params
func (o *DownloadBatchParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithID adds the id to the download batch params
func (o *DownloadBatchParams) WithID(id int32) *DownloadBatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the download batch params
func (o *DownloadBatchParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param batchId
	if err := r.SetPathParam("batchId", swag.FormatInt32(o.BatchID)); err != nil {
		return err
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
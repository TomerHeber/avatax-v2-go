// Code generated by go-swagger; DO NOT EDIT.

package items

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

// NewBatchDeleteItemParametersParams creates a new BatchDeleteItemParametersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBatchDeleteItemParametersParams() *BatchDeleteItemParametersParams {
	return &BatchDeleteItemParametersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBatchDeleteItemParametersParamsWithTimeout creates a new BatchDeleteItemParametersParams object
// with the ability to set a timeout on a request.
func NewBatchDeleteItemParametersParamsWithTimeout(timeout time.Duration) *BatchDeleteItemParametersParams {
	return &BatchDeleteItemParametersParams{
		timeout: timeout,
	}
}

// NewBatchDeleteItemParametersParamsWithContext creates a new BatchDeleteItemParametersParams object
// with the ability to set a context for a request.
func NewBatchDeleteItemParametersParamsWithContext(ctx context.Context) *BatchDeleteItemParametersParams {
	return &BatchDeleteItemParametersParams{
		Context: ctx,
	}
}

// NewBatchDeleteItemParametersParamsWithHTTPClient creates a new BatchDeleteItemParametersParams object
// with the ability to set a custom HTTPClient for a request.
func NewBatchDeleteItemParametersParamsWithHTTPClient(client *http.Client) *BatchDeleteItemParametersParams {
	return &BatchDeleteItemParametersParams{
		HTTPClient: client,
	}
}

/* BatchDeleteItemParametersParams contains all the parameters to send to the API endpoint
   for the batch delete item parameters operation.

   Typically these are written to a http.Request.
*/
type BatchDeleteItemParametersParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyID.

	   The ID of the company that owns this item.

	   Format: int32
	*/
	CompanyID int32

	/* ItemID.

	   The ID of the item you wish to delete the parameters.

	   Format: int64
	*/
	ItemID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the batch delete item parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchDeleteItemParametersParams) WithDefaults() *BatchDeleteItemParametersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the batch delete item parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchDeleteItemParametersParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := BatchDeleteItemParametersParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) WithTimeout(timeout time.Duration) *BatchDeleteItemParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) WithContext(ctx context.Context) *BatchDeleteItemParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) WithHTTPClient(client *http.Client) *BatchDeleteItemParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) WithXAvalaraClient(xAvalaraClient *string) *BatchDeleteItemParametersParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyID adds the companyID to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) WithCompanyID(companyID int32) *BatchDeleteItemParametersParams {
	o.SetCompanyID(companyID)
	return o
}

// SetCompanyID adds the companyId to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) SetCompanyID(companyID int32) {
	o.CompanyID = companyID
}

// WithItemID adds the itemID to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) WithItemID(itemID int64) *BatchDeleteItemParametersParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the batch delete item parameters params
func (o *BatchDeleteItemParametersParams) SetItemID(itemID int64) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *BatchDeleteItemParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param companyId
	if err := r.SetPathParam("companyId", swag.FormatInt32(o.CompanyID)); err != nil {
		return err
	}

	// path param itemId
	if err := r.SetPathParam("itemId", swag.FormatInt64(o.ItemID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
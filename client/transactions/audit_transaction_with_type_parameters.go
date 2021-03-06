// Code generated by go-swagger; DO NOT EDIT.

package transactions

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
)

// NewAuditTransactionWithTypeParams creates a new AuditTransactionWithTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditTransactionWithTypeParams() *AuditTransactionWithTypeParams {
	return &AuditTransactionWithTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditTransactionWithTypeParamsWithTimeout creates a new AuditTransactionWithTypeParams object
// with the ability to set a timeout on a request.
func NewAuditTransactionWithTypeParamsWithTimeout(timeout time.Duration) *AuditTransactionWithTypeParams {
	return &AuditTransactionWithTypeParams{
		timeout: timeout,
	}
}

// NewAuditTransactionWithTypeParamsWithContext creates a new AuditTransactionWithTypeParams object
// with the ability to set a context for a request.
func NewAuditTransactionWithTypeParamsWithContext(ctx context.Context) *AuditTransactionWithTypeParams {
	return &AuditTransactionWithTypeParams{
		Context: ctx,
	}
}

// NewAuditTransactionWithTypeParamsWithHTTPClient creates a new AuditTransactionWithTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditTransactionWithTypeParamsWithHTTPClient(client *http.Client) *AuditTransactionWithTypeParams {
	return &AuditTransactionWithTypeParams{
		HTTPClient: client,
	}
}

/* AuditTransactionWithTypeParams contains all the parameters to send to the API endpoint
   for the audit transaction with type operation.

   Typically these are written to a http.Request.
*/
type AuditTransactionWithTypeParams struct {

	/* XAvalaraClient.

	   Identifies the software you are using to call this API.  For more information on the client header, see [Client Headers](https://developer.avalara.com/avatax/client-headers/) .

	   Default: "Swagger UI; 21.12.0; Custom; 1.0"
	*/
	XAvalaraClient *string

	/* CompanyCode.

	   The code identifying the company that owns this transaction
	*/
	CompanyCode string

	/* DocumentType.

	   The document type of the original transaction
	*/
	DocumentType string

	/* TransactionCode.

	   The code identifying the transaction
	*/
	TransactionCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit transaction with type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditTransactionWithTypeParams) WithDefaults() *AuditTransactionWithTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit transaction with type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditTransactionWithTypeParams) SetDefaults() {
	var (
		xAvalaraClientDefault = string("Swagger UI; 21.12.0; Custom; 1.0")
	)

	val := AuditTransactionWithTypeParams{
		XAvalaraClient: &xAvalaraClientDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) WithTimeout(timeout time.Duration) *AuditTransactionWithTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) WithContext(ctx context.Context) *AuditTransactionWithTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) WithHTTPClient(client *http.Client) *AuditTransactionWithTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAvalaraClient adds the xAvalaraClient to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) WithXAvalaraClient(xAvalaraClient *string) *AuditTransactionWithTypeParams {
	o.SetXAvalaraClient(xAvalaraClient)
	return o
}

// SetXAvalaraClient adds the xAvalaraClient to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) SetXAvalaraClient(xAvalaraClient *string) {
	o.XAvalaraClient = xAvalaraClient
}

// WithCompanyCode adds the companyCode to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) WithCompanyCode(companyCode string) *AuditTransactionWithTypeParams {
	o.SetCompanyCode(companyCode)
	return o
}

// SetCompanyCode adds the companyCode to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) SetCompanyCode(companyCode string) {
	o.CompanyCode = companyCode
}

// WithDocumentType adds the documentType to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) WithDocumentType(documentType string) *AuditTransactionWithTypeParams {
	o.SetDocumentType(documentType)
	return o
}

// SetDocumentType adds the documentType to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) SetDocumentType(documentType string) {
	o.DocumentType = documentType
}

// WithTransactionCode adds the transactionCode to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) WithTransactionCode(transactionCode string) *AuditTransactionWithTypeParams {
	o.SetTransactionCode(transactionCode)
	return o
}

// SetTransactionCode adds the transactionCode to the audit transaction with type params
func (o *AuditTransactionWithTypeParams) SetTransactionCode(transactionCode string) {
	o.TransactionCode = transactionCode
}

// WriteToRequest writes these params to a swagger request
func (o *AuditTransactionWithTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param companyCode
	if err := r.SetPathParam("companyCode", o.CompanyCode); err != nil {
		return err
	}

	// path param documentType
	if err := r.SetPathParam("documentType", o.DocumentType); err != nil {
		return err
	}

	// path param transactionCode
	if err := r.SetPathParam("transactionCode", o.TransactionCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

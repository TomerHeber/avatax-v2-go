// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// AdjustTransactionReader is a Reader for the AdjustTransaction structure.
type AdjustTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdjustTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdjustTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdjustTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdjustTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdjustTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdjustTransactionOK creates a AdjustTransactionOK with default headers values
func NewAdjustTransactionOK() *AdjustTransactionOK {
	return &AdjustTransactionOK{}
}

/* AdjustTransactionOK describes a response with status code 200, with default header values.

Success
*/
type AdjustTransactionOK struct {
	Payload *models.TransactionModel
}

func (o *AdjustTransactionOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyCode}/transactions/{transactionCode}/adjust][%d] adjustTransactionOK  %+v", 200, o.Payload)
}
func (o *AdjustTransactionOK) GetPayload() *models.TransactionModel {
	return o.Payload
}

func (o *AdjustTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdjustTransactionBadRequest creates a AdjustTransactionBadRequest with default headers values
func NewAdjustTransactionBadRequest() *AdjustTransactionBadRequest {
	return &AdjustTransactionBadRequest{}
}

/* AdjustTransactionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdjustTransactionBadRequest struct {
}

func (o *AdjustTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyCode}/transactions/{transactionCode}/adjust][%d] adjustTransactionBadRequest ", 400)
}

func (o *AdjustTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdjustTransactionUnauthorized creates a AdjustTransactionUnauthorized with default headers values
func NewAdjustTransactionUnauthorized() *AdjustTransactionUnauthorized {
	return &AdjustTransactionUnauthorized{}
}

/* AdjustTransactionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdjustTransactionUnauthorized struct {
}

func (o *AdjustTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyCode}/transactions/{transactionCode}/adjust][%d] adjustTransactionUnauthorized ", 401)
}

func (o *AdjustTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdjustTransactionNotFound creates a AdjustTransactionNotFound with default headers values
func NewAdjustTransactionNotFound() *AdjustTransactionNotFound {
	return &AdjustTransactionNotFound{}
}

/* AdjustTransactionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdjustTransactionNotFound struct {
}

func (o *AdjustTransactionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyCode}/transactions/{transactionCode}/adjust][%d] adjustTransactionNotFound ", 404)
}

func (o *AdjustTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
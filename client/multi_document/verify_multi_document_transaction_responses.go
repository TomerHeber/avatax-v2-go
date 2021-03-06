// Code generated by go-swagger; DO NOT EDIT.

package multi_document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// VerifyMultiDocumentTransactionReader is a Reader for the VerifyMultiDocumentTransaction structure.
type VerifyMultiDocumentTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyMultiDocumentTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerifyMultiDocumentTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVerifyMultiDocumentTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewVerifyMultiDocumentTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVerifyMultiDocumentTransactionOK creates a VerifyMultiDocumentTransactionOK with default headers values
func NewVerifyMultiDocumentTransactionOK() *VerifyMultiDocumentTransactionOK {
	return &VerifyMultiDocumentTransactionOK{}
}

/* VerifyMultiDocumentTransactionOK describes a response with status code 200, with default header values.

Success
*/
type VerifyMultiDocumentTransactionOK struct {
	Payload *models.MultiDocumentModel
}

func (o *VerifyMultiDocumentTransactionOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/multidocument/verify][%d] verifyMultiDocumentTransactionOK  %+v", 200, o.Payload)
}
func (o *VerifyMultiDocumentTransactionOK) GetPayload() *models.MultiDocumentModel {
	return o.Payload
}

func (o *VerifyMultiDocumentTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MultiDocumentModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyMultiDocumentTransactionBadRequest creates a VerifyMultiDocumentTransactionBadRequest with default headers values
func NewVerifyMultiDocumentTransactionBadRequest() *VerifyMultiDocumentTransactionBadRequest {
	return &VerifyMultiDocumentTransactionBadRequest{}
}

/* VerifyMultiDocumentTransactionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VerifyMultiDocumentTransactionBadRequest struct {
}

func (o *VerifyMultiDocumentTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/multidocument/verify][%d] verifyMultiDocumentTransactionBadRequest ", 400)
}

func (o *VerifyMultiDocumentTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVerifyMultiDocumentTransactionUnauthorized creates a VerifyMultiDocumentTransactionUnauthorized with default headers values
func NewVerifyMultiDocumentTransactionUnauthorized() *VerifyMultiDocumentTransactionUnauthorized {
	return &VerifyMultiDocumentTransactionUnauthorized{}
}

/* VerifyMultiDocumentTransactionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type VerifyMultiDocumentTransactionUnauthorized struct {
}

func (o *VerifyMultiDocumentTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/multidocument/verify][%d] verifyMultiDocumentTransactionUnauthorized ", 401)
}

func (o *VerifyMultiDocumentTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

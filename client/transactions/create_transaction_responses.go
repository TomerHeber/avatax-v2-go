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

// CreateTransactionReader is a Reader for the CreateTransaction structure.
type CreateTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTransactionOK creates a CreateTransactionOK with default headers values
func NewCreateTransactionOK() *CreateTransactionOK {
	return &CreateTransactionOK{}
}

/* CreateTransactionOK describes a response with status code 200, with default header values.

Success
*/
type CreateTransactionOK struct {
	Payload *models.TransactionModel
}

func (o *CreateTransactionOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/create][%d] createTransactionOK  %+v", 200, o.Payload)
}
func (o *CreateTransactionOK) GetPayload() *models.TransactionModel {
	return o.Payload
}

func (o *CreateTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionBadRequest creates a CreateTransactionBadRequest with default headers values
func NewCreateTransactionBadRequest() *CreateTransactionBadRequest {
	return &CreateTransactionBadRequest{}
}

/* CreateTransactionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateTransactionBadRequest struct {
}

func (o *CreateTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/create][%d] createTransactionBadRequest ", 400)
}

func (o *CreateTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTransactionUnauthorized creates a CreateTransactionUnauthorized with default headers values
func NewCreateTransactionUnauthorized() *CreateTransactionUnauthorized {
	return &CreateTransactionUnauthorized{}
}

/* CreateTransactionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateTransactionUnauthorized struct {
}

func (o *CreateTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/create][%d] createTransactionUnauthorized ", 401)
}

func (o *CreateTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
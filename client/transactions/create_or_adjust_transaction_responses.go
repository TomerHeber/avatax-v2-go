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

// CreateOrAdjustTransactionReader is a Reader for the CreateOrAdjustTransaction structure.
type CreateOrAdjustTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrAdjustTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrAdjustTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOrAdjustTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateOrAdjustTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrAdjustTransactionOK creates a CreateOrAdjustTransactionOK with default headers values
func NewCreateOrAdjustTransactionOK() *CreateOrAdjustTransactionOK {
	return &CreateOrAdjustTransactionOK{}
}

/* CreateOrAdjustTransactionOK describes a response with status code 200, with default header values.

Success
*/
type CreateOrAdjustTransactionOK struct {
	Payload *models.TransactionModel
}

func (o *CreateOrAdjustTransactionOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/createoradjust][%d] createOrAdjustTransactionOK  %+v", 200, o.Payload)
}
func (o *CreateOrAdjustTransactionOK) GetPayload() *models.TransactionModel {
	return o.Payload
}

func (o *CreateOrAdjustTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrAdjustTransactionBadRequest creates a CreateOrAdjustTransactionBadRequest with default headers values
func NewCreateOrAdjustTransactionBadRequest() *CreateOrAdjustTransactionBadRequest {
	return &CreateOrAdjustTransactionBadRequest{}
}

/* CreateOrAdjustTransactionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateOrAdjustTransactionBadRequest struct {
}

func (o *CreateOrAdjustTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/createoradjust][%d] createOrAdjustTransactionBadRequest ", 400)
}

func (o *CreateOrAdjustTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrAdjustTransactionUnauthorized creates a CreateOrAdjustTransactionUnauthorized with default headers values
func NewCreateOrAdjustTransactionUnauthorized() *CreateOrAdjustTransactionUnauthorized {
	return &CreateOrAdjustTransactionUnauthorized{}
}

/* CreateOrAdjustTransactionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateOrAdjustTransactionUnauthorized struct {
}

func (o *CreateOrAdjustTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/transactions/createoradjust][%d] createOrAdjustTransactionUnauthorized ", 401)
}

func (o *CreateOrAdjustTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
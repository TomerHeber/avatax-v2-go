// Code generated by go-swagger; DO NOT EDIT.

package batches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// CreateTransactionBatchReader is a Reader for the CreateTransactionBatch structure.
type CreateTransactionBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransactionBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTransactionBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTransactionBatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTransactionBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTransactionBatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTransactionBatchOK creates a CreateTransactionBatchOK with default headers values
func NewCreateTransactionBatchOK() *CreateTransactionBatchOK {
	return &CreateTransactionBatchOK{}
}

/* CreateTransactionBatchOK describes a response with status code 200, with default header values.

Success
*/
type CreateTransactionBatchOK struct {
	Payload *models.CreateTransactionBatchResponseModel
}

func (o *CreateTransactionBatchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/batches/transactions][%d] createTransactionBatchOK  %+v", 200, o.Payload)
}
func (o *CreateTransactionBatchOK) GetPayload() *models.CreateTransactionBatchResponseModel {
	return o.Payload
}

func (o *CreateTransactionBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateTransactionBatchResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionBatchBadRequest creates a CreateTransactionBatchBadRequest with default headers values
func NewCreateTransactionBatchBadRequest() *CreateTransactionBatchBadRequest {
	return &CreateTransactionBatchBadRequest{}
}

/* CreateTransactionBatchBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateTransactionBatchBadRequest struct {
}

func (o *CreateTransactionBatchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/batches/transactions][%d] createTransactionBatchBadRequest ", 400)
}

func (o *CreateTransactionBatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTransactionBatchUnauthorized creates a CreateTransactionBatchUnauthorized with default headers values
func NewCreateTransactionBatchUnauthorized() *CreateTransactionBatchUnauthorized {
	return &CreateTransactionBatchUnauthorized{}
}

/* CreateTransactionBatchUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateTransactionBatchUnauthorized struct {
}

func (o *CreateTransactionBatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/batches/transactions][%d] createTransactionBatchUnauthorized ", 401)
}

func (o *CreateTransactionBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTransactionBatchNotFound creates a CreateTransactionBatchNotFound with default headers values
func NewCreateTransactionBatchNotFound() *CreateTransactionBatchNotFound {
	return &CreateTransactionBatchNotFound{}
}

/* CreateTransactionBatchNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateTransactionBatchNotFound struct {
}

func (o *CreateTransactionBatchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/batches/transactions][%d] createTransactionBatchNotFound ", 404)
}

func (o *CreateTransactionBatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
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

// CancelBatchReader is a Reader for the CancelBatch structure.
type CancelBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelBatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelBatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelBatchOK creates a CancelBatchOK with default headers values
func NewCancelBatchOK() *CancelBatchOK {
	return &CancelBatchOK{}
}

/* CancelBatchOK describes a response with status code 200, with default header values.

Success
*/
type CancelBatchOK struct {
	Payload *models.BatchModel
}

func (o *CancelBatchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/batches/{id}/cancel][%d] cancelBatchOK  %+v", 200, o.Payload)
}
func (o *CancelBatchOK) GetPayload() *models.BatchModel {
	return o.Payload
}

func (o *CancelBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelBatchBadRequest creates a CancelBatchBadRequest with default headers values
func NewCancelBatchBadRequest() *CancelBatchBadRequest {
	return &CancelBatchBadRequest{}
}

/* CancelBatchBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CancelBatchBadRequest struct {
}

func (o *CancelBatchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/batches/{id}/cancel][%d] cancelBatchBadRequest ", 400)
}

func (o *CancelBatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelBatchUnauthorized creates a CancelBatchUnauthorized with default headers values
func NewCancelBatchUnauthorized() *CancelBatchUnauthorized {
	return &CancelBatchUnauthorized{}
}

/* CancelBatchUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CancelBatchUnauthorized struct {
}

func (o *CancelBatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/batches/{id}/cancel][%d] cancelBatchUnauthorized ", 401)
}

func (o *CancelBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelBatchNotFound creates a CancelBatchNotFound with default headers values
func NewCancelBatchNotFound() *CancelBatchNotFound {
	return &CancelBatchNotFound{}
}

/* CancelBatchNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CancelBatchNotFound struct {
}

func (o *CancelBatchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/batches/{id}/cancel][%d] cancelBatchNotFound ", 404)
}

func (o *CancelBatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
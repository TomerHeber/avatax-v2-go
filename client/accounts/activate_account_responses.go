// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// ActivateAccountReader is a Reader for the ActivateAccount structure.
type ActivateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivateAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewActivateAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewActivateAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActivateAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewActivateAccountOK creates a ActivateAccountOK with default headers values
func NewActivateAccountOK() *ActivateAccountOK {
	return &ActivateAccountOK{}
}

/* ActivateAccountOK describes a response with status code 200, with default header values.

Success
*/
type ActivateAccountOK struct {
	Payload *models.AccountModel
}

func (o *ActivateAccountOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts/{id}/activate][%d] activateAccountOK  %+v", 200, o.Payload)
}
func (o *ActivateAccountOK) GetPayload() *models.AccountModel {
	return o.Payload
}

func (o *ActivateAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateAccountBadRequest creates a ActivateAccountBadRequest with default headers values
func NewActivateAccountBadRequest() *ActivateAccountBadRequest {
	return &ActivateAccountBadRequest{}
}

/* ActivateAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ActivateAccountBadRequest struct {
}

func (o *ActivateAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts/{id}/activate][%d] activateAccountBadRequest ", 400)
}

func (o *ActivateAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActivateAccountUnauthorized creates a ActivateAccountUnauthorized with default headers values
func NewActivateAccountUnauthorized() *ActivateAccountUnauthorized {
	return &ActivateAccountUnauthorized{}
}

/* ActivateAccountUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ActivateAccountUnauthorized struct {
}

func (o *ActivateAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts/{id}/activate][%d] activateAccountUnauthorized ", 401)
}

func (o *ActivateAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActivateAccountNotFound creates a ActivateAccountNotFound with default headers values
func NewActivateAccountNotFound() *ActivateAccountNotFound {
	return &ActivateAccountNotFound{}
}

/* ActivateAccountNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ActivateAccountNotFound struct {
}

func (o *ActivateAccountNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts/{id}/activate][%d] activateAccountNotFound ", 404)
}

func (o *ActivateAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package e_commerce_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// CreateECommerceTokenReader is a Reader for the CreateECommerceToken structure.
type CreateECommerceTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateECommerceTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateECommerceTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateECommerceTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateECommerceTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateECommerceTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateECommerceTokenOK creates a CreateECommerceTokenOK with default headers values
func NewCreateECommerceTokenOK() *CreateECommerceTokenOK {
	return &CreateECommerceTokenOK{}
}

/* CreateECommerceTokenOK describes a response with status code 200, with default header values.

Success
*/
type CreateECommerceTokenOK struct {
	Payload *models.ECommerceTokenOutputModel
}

func (o *CreateECommerceTokenOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/ecommercetokens][%d] createECommerceTokenOK  %+v", 200, o.Payload)
}
func (o *CreateECommerceTokenOK) GetPayload() *models.ECommerceTokenOutputModel {
	return o.Payload
}

func (o *CreateECommerceTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ECommerceTokenOutputModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateECommerceTokenBadRequest creates a CreateECommerceTokenBadRequest with default headers values
func NewCreateECommerceTokenBadRequest() *CreateECommerceTokenBadRequest {
	return &CreateECommerceTokenBadRequest{}
}

/* CreateECommerceTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateECommerceTokenBadRequest struct {
}

func (o *CreateECommerceTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/ecommercetokens][%d] createECommerceTokenBadRequest ", 400)
}

func (o *CreateECommerceTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateECommerceTokenUnauthorized creates a CreateECommerceTokenUnauthorized with default headers values
func NewCreateECommerceTokenUnauthorized() *CreateECommerceTokenUnauthorized {
	return &CreateECommerceTokenUnauthorized{}
}

/* CreateECommerceTokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateECommerceTokenUnauthorized struct {
}

func (o *CreateECommerceTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/ecommercetokens][%d] createECommerceTokenUnauthorized ", 401)
}

func (o *CreateECommerceTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateECommerceTokenInternalServerError creates a CreateECommerceTokenInternalServerError with default headers values
func NewCreateECommerceTokenInternalServerError() *CreateECommerceTokenInternalServerError {
	return &CreateECommerceTokenInternalServerError{}
}

/* CreateECommerceTokenInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CreateECommerceTokenInternalServerError struct {
}

func (o *CreateECommerceTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/ecommercetokens][%d] createECommerceTokenInternalServerError ", 500)
}

func (o *CreateECommerceTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
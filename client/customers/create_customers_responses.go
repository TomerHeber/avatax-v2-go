// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// CreateCustomersReader is a Reader for the CreateCustomers structure.
type CreateCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCustomersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCustomersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCustomersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCustomersOK creates a CreateCustomersOK with default headers values
func NewCreateCustomersOK() *CreateCustomersOK {
	return &CreateCustomersOK{}
}

/* CreateCustomersOK describes a response with status code 200, with default header values.

Success
*/
type CreateCustomersOK struct {
	Payload []*models.CustomerModel
}

func (o *CreateCustomersOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/customers][%d] createCustomersOK  %+v", 200, o.Payload)
}
func (o *CreateCustomersOK) GetPayload() []*models.CustomerModel {
	return o.Payload
}

func (o *CreateCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomersBadRequest creates a CreateCustomersBadRequest with default headers values
func NewCreateCustomersBadRequest() *CreateCustomersBadRequest {
	return &CreateCustomersBadRequest{}
}

/* CreateCustomersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateCustomersBadRequest struct {
}

func (o *CreateCustomersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/customers][%d] createCustomersBadRequest ", 400)
}

func (o *CreateCustomersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCustomersUnauthorized creates a CreateCustomersUnauthorized with default headers values
func NewCreateCustomersUnauthorized() *CreateCustomersUnauthorized {
	return &CreateCustomersUnauthorized{}
}

/* CreateCustomersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateCustomersUnauthorized struct {
}

func (o *CreateCustomersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/customers][%d] createCustomersUnauthorized ", 401)
}

func (o *CreateCustomersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCustomersNotFound creates a CreateCustomersNotFound with default headers values
func NewCreateCustomersNotFound() *CreateCustomersNotFound {
	return &CreateCustomersNotFound{}
}

/* CreateCustomersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateCustomersNotFound struct {
}

func (o *CreateCustomersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/customers][%d] createCustomersNotFound ", 404)
}

func (o *CreateCustomersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package registrar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// CreateAccountReader is a Reader for the CreateAccount structure.
type CreateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAccountOK creates a CreateAccountOK with default headers values
func NewCreateAccountOK() *CreateAccountOK {
	return &CreateAccountOK{}
}

/* CreateAccountOK describes a response with status code 200, with default header values.

Success
*/
type CreateAccountOK struct {
	Payload []*models.AccountModel
}

func (o *CreateAccountOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts][%d] createAccountOK  %+v", 200, o.Payload)
}
func (o *CreateAccountOK) GetPayload() []*models.AccountModel {
	return o.Payload
}

func (o *CreateAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountBadRequest creates a CreateAccountBadRequest with default headers values
func NewCreateAccountBadRequest() *CreateAccountBadRequest {
	return &CreateAccountBadRequest{}
}

/* CreateAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAccountBadRequest struct {
}

func (o *CreateAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts][%d] createAccountBadRequest ", 400)
}

func (o *CreateAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccountUnauthorized creates a CreateAccountUnauthorized with default headers values
func NewCreateAccountUnauthorized() *CreateAccountUnauthorized {
	return &CreateAccountUnauthorized{}
}

/* CreateAccountUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAccountUnauthorized struct {
}

func (o *CreateAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts][%d] createAccountUnauthorized ", 401)
}

func (o *CreateAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

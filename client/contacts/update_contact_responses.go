// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// UpdateContactReader is a Reader for the UpdateContact structure.
type UpdateContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateContactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateContactOK creates a UpdateContactOK with default headers values
func NewUpdateContactOK() *UpdateContactOK {
	return &UpdateContactOK{}
}

/* UpdateContactOK describes a response with status code 200, with default header values.

Success
*/
type UpdateContactOK struct {
	Payload *models.ContactModel
}

func (o *UpdateContactOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/contacts/{id}][%d] updateContactOK  %+v", 200, o.Payload)
}
func (o *UpdateContactOK) GetPayload() *models.ContactModel {
	return o.Payload
}

func (o *UpdateContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContactBadRequest creates a UpdateContactBadRequest with default headers values
func NewUpdateContactBadRequest() *UpdateContactBadRequest {
	return &UpdateContactBadRequest{}
}

/* UpdateContactBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateContactBadRequest struct {
}

func (o *UpdateContactBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/contacts/{id}][%d] updateContactBadRequest ", 400)
}

func (o *UpdateContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateContactUnauthorized creates a UpdateContactUnauthorized with default headers values
func NewUpdateContactUnauthorized() *UpdateContactUnauthorized {
	return &UpdateContactUnauthorized{}
}

/* UpdateContactUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateContactUnauthorized struct {
}

func (o *UpdateContactUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/contacts/{id}][%d] updateContactUnauthorized ", 401)
}

func (o *UpdateContactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateContactNotFound creates a UpdateContactNotFound with default headers values
func NewUpdateContactNotFound() *UpdateContactNotFound {
	return &UpdateContactNotFound{}
}

/* UpdateContactNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateContactNotFound struct {
}

func (o *UpdateContactNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/contacts/{id}][%d] updateContactNotFound ", 404)
}

func (o *UpdateContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

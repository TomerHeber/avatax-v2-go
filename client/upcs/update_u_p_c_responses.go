// Code generated by go-swagger; DO NOT EDIT.

package upcs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// UpdateUPCReader is a Reader for the UpdateUPC structure.
type UpdateUPCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUPCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUPCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUPCBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateUPCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUPCNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUPCOK creates a UpdateUPCOK with default headers values
func NewUpdateUPCOK() *UpdateUPCOK {
	return &UpdateUPCOK{}
}

/* UpdateUPCOK describes a response with status code 200, with default header values.

Success
*/
type UpdateUPCOK struct {
	Payload *models.UPCModel
}

func (o *UpdateUPCOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/upcs/{id}][%d] updateUPCOK  %+v", 200, o.Payload)
}
func (o *UpdateUPCOK) GetPayload() *models.UPCModel {
	return o.Payload
}

func (o *UpdateUPCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UPCModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUPCBadRequest creates a UpdateUPCBadRequest with default headers values
func NewUpdateUPCBadRequest() *UpdateUPCBadRequest {
	return &UpdateUPCBadRequest{}
}

/* UpdateUPCBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateUPCBadRequest struct {
}

func (o *UpdateUPCBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/upcs/{id}][%d] updateUPCBadRequest ", 400)
}

func (o *UpdateUPCBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUPCUnauthorized creates a UpdateUPCUnauthorized with default headers values
func NewUpdateUPCUnauthorized() *UpdateUPCUnauthorized {
	return &UpdateUPCUnauthorized{}
}

/* UpdateUPCUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateUPCUnauthorized struct {
}

func (o *UpdateUPCUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/upcs/{id}][%d] updateUPCUnauthorized ", 401)
}

func (o *UpdateUPCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUPCNotFound creates a UpdateUPCNotFound with default headers values
func NewUpdateUPCNotFound() *UpdateUPCNotFound {
	return &UpdateUPCNotFound{}
}

/* UpdateUPCNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateUPCNotFound struct {
}

func (o *UpdateUPCNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/upcs/{id}][%d] updateUPCNotFound ", 404)
}

func (o *UpdateUPCNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
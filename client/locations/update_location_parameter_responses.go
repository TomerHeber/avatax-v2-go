// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// UpdateLocationParameterReader is a Reader for the UpdateLocationParameter structure.
type UpdateLocationParameterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLocationParameterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLocationParameterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLocationParameterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateLocationParameterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLocationParameterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLocationParameterOK creates a UpdateLocationParameterOK with default headers values
func NewUpdateLocationParameterOK() *UpdateLocationParameterOK {
	return &UpdateLocationParameterOK{}
}

/* UpdateLocationParameterOK describes a response with status code 200, with default header values.

Success
*/
type UpdateLocationParameterOK struct {
	Payload *models.LocationParameterModel
}

func (o *UpdateLocationParameterOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/locations/{locationId}/parameters/{id}][%d] updateLocationParameterOK  %+v", 200, o.Payload)
}
func (o *UpdateLocationParameterOK) GetPayload() *models.LocationParameterModel {
	return o.Payload
}

func (o *UpdateLocationParameterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationParameterModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLocationParameterBadRequest creates a UpdateLocationParameterBadRequest with default headers values
func NewUpdateLocationParameterBadRequest() *UpdateLocationParameterBadRequest {
	return &UpdateLocationParameterBadRequest{}
}

/* UpdateLocationParameterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateLocationParameterBadRequest struct {
}

func (o *UpdateLocationParameterBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/locations/{locationId}/parameters/{id}][%d] updateLocationParameterBadRequest ", 400)
}

func (o *UpdateLocationParameterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLocationParameterUnauthorized creates a UpdateLocationParameterUnauthorized with default headers values
func NewUpdateLocationParameterUnauthorized() *UpdateLocationParameterUnauthorized {
	return &UpdateLocationParameterUnauthorized{}
}

/* UpdateLocationParameterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateLocationParameterUnauthorized struct {
}

func (o *UpdateLocationParameterUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/locations/{locationId}/parameters/{id}][%d] updateLocationParameterUnauthorized ", 401)
}

func (o *UpdateLocationParameterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLocationParameterNotFound creates a UpdateLocationParameterNotFound with default headers values
func NewUpdateLocationParameterNotFound() *UpdateLocationParameterNotFound {
	return &UpdateLocationParameterNotFound{}
}

/* UpdateLocationParameterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateLocationParameterNotFound struct {
}

func (o *UpdateLocationParameterNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/locations/{locationId}/parameters/{id}][%d] updateLocationParameterNotFound ", 404)
}

func (o *UpdateLocationParameterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

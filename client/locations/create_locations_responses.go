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

// CreateLocationsReader is a Reader for the CreateLocations structure.
type CreateLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateLocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateLocationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLocationsOK creates a CreateLocationsOK with default headers values
func NewCreateLocationsOK() *CreateLocationsOK {
	return &CreateLocationsOK{}
}

/* CreateLocationsOK describes a response with status code 200, with default header values.

Success
*/
type CreateLocationsOK struct {
	Payload []*models.LocationModel
}

func (o *CreateLocationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/locations][%d] createLocationsOK  %+v", 200, o.Payload)
}
func (o *CreateLocationsOK) GetPayload() []*models.LocationModel {
	return o.Payload
}

func (o *CreateLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLocationsBadRequest creates a CreateLocationsBadRequest with default headers values
func NewCreateLocationsBadRequest() *CreateLocationsBadRequest {
	return &CreateLocationsBadRequest{}
}

/* CreateLocationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateLocationsBadRequest struct {
}

func (o *CreateLocationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/locations][%d] createLocationsBadRequest ", 400)
}

func (o *CreateLocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLocationsUnauthorized creates a CreateLocationsUnauthorized with default headers values
func NewCreateLocationsUnauthorized() *CreateLocationsUnauthorized {
	return &CreateLocationsUnauthorized{}
}

/* CreateLocationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateLocationsUnauthorized struct {
}

func (o *CreateLocationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/locations][%d] createLocationsUnauthorized ", 401)
}

func (o *CreateLocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLocationsNotFound creates a CreateLocationsNotFound with default headers values
func NewCreateLocationsNotFound() *CreateLocationsNotFound {
	return &CreateLocationsNotFound{}
}

/* CreateLocationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateLocationsNotFound struct {
}

func (o *CreateLocationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/locations][%d] createLocationsNotFound ", 404)
}

func (o *CreateLocationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

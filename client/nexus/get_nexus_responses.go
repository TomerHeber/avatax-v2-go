// Code generated by go-swagger; DO NOT EDIT.

package nexus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// GetNexusReader is a Reader for the GetNexus structure.
type GetNexusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNexusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNexusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNexusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNexusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNexusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNexusOK creates a GetNexusOK with default headers values
func NewGetNexusOK() *GetNexusOK {
	return &GetNexusOK{}
}

/* GetNexusOK describes a response with status code 200, with default header values.

Success
*/
type GetNexusOK struct {
	Payload *models.NexusModel
}

func (o *GetNexusOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/nexus/{id}][%d] getNexusOK  %+v", 200, o.Payload)
}
func (o *GetNexusOK) GetPayload() *models.NexusModel {
	return o.Payload
}

func (o *GetNexusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NexusModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNexusBadRequest creates a GetNexusBadRequest with default headers values
func NewGetNexusBadRequest() *GetNexusBadRequest {
	return &GetNexusBadRequest{}
}

/* GetNexusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNexusBadRequest struct {
}

func (o *GetNexusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/nexus/{id}][%d] getNexusBadRequest ", 400)
}

func (o *GetNexusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNexusUnauthorized creates a GetNexusUnauthorized with default headers values
func NewGetNexusUnauthorized() *GetNexusUnauthorized {
	return &GetNexusUnauthorized{}
}

/* GetNexusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNexusUnauthorized struct {
}

func (o *GetNexusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/nexus/{id}][%d] getNexusUnauthorized ", 401)
}

func (o *GetNexusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNexusNotFound creates a GetNexusNotFound with default headers values
func NewGetNexusNotFound() *GetNexusNotFound {
	return &GetNexusNotFound{}
}

/* GetNexusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetNexusNotFound struct {
}

func (o *GetNexusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/nexus/{id}][%d] getNexusNotFound ", 404)
}

func (o *GetNexusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

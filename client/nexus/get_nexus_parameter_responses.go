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

// GetNexusParameterReader is a Reader for the GetNexusParameter structure.
type GetNexusParameterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNexusParameterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNexusParameterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNexusParameterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNexusParameterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNexusParameterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNexusParameterOK creates a GetNexusParameterOK with default headers values
func NewGetNexusParameterOK() *GetNexusParameterOK {
	return &GetNexusParameterOK{}
}

/* GetNexusParameterOK describes a response with status code 200, with default header values.

Success
*/
type GetNexusParameterOK struct {
	Payload *models.NexusParameterDetailModel
}

func (o *GetNexusParameterOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/nexus/{nexusId}/parameters/{id}][%d] getNexusParameterOK  %+v", 200, o.Payload)
}
func (o *GetNexusParameterOK) GetPayload() *models.NexusParameterDetailModel {
	return o.Payload
}

func (o *GetNexusParameterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NexusParameterDetailModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNexusParameterBadRequest creates a GetNexusParameterBadRequest with default headers values
func NewGetNexusParameterBadRequest() *GetNexusParameterBadRequest {
	return &GetNexusParameterBadRequest{}
}

/* GetNexusParameterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNexusParameterBadRequest struct {
}

func (o *GetNexusParameterBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/nexus/{nexusId}/parameters/{id}][%d] getNexusParameterBadRequest ", 400)
}

func (o *GetNexusParameterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNexusParameterUnauthorized creates a GetNexusParameterUnauthorized with default headers values
func NewGetNexusParameterUnauthorized() *GetNexusParameterUnauthorized {
	return &GetNexusParameterUnauthorized{}
}

/* GetNexusParameterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNexusParameterUnauthorized struct {
}

func (o *GetNexusParameterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/nexus/{nexusId}/parameters/{id}][%d] getNexusParameterUnauthorized ", 401)
}

func (o *GetNexusParameterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNexusParameterNotFound creates a GetNexusParameterNotFound with default headers values
func NewGetNexusParameterNotFound() *GetNexusParameterNotFound {
	return &GetNexusParameterNotFound{}
}

/* GetNexusParameterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetNexusParameterNotFound struct {
}

func (o *GetNexusParameterNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/nexus/{nexusId}/parameters/{id}][%d] getNexusParameterNotFound ", 404)
}

func (o *GetNexusParameterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

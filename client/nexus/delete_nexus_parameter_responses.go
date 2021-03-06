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

// DeleteNexusParameterReader is a Reader for the DeleteNexusParameter structure.
type DeleteNexusParameterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNexusParameterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNexusParameterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNexusParameterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNexusParameterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNexusParameterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNexusParameterOK creates a DeleteNexusParameterOK with default headers values
func NewDeleteNexusParameterOK() *DeleteNexusParameterOK {
	return &DeleteNexusParameterOK{}
}

/* DeleteNexusParameterOK describes a response with status code 200, with default header values.

Success
*/
type DeleteNexusParameterOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteNexusParameterOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/nexus/{nexusId}/parameters/{id}][%d] deleteNexusParameterOK  %+v", 200, o.Payload)
}
func (o *DeleteNexusParameterOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteNexusParameterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNexusParameterBadRequest creates a DeleteNexusParameterBadRequest with default headers values
func NewDeleteNexusParameterBadRequest() *DeleteNexusParameterBadRequest {
	return &DeleteNexusParameterBadRequest{}
}

/* DeleteNexusParameterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteNexusParameterBadRequest struct {
}

func (o *DeleteNexusParameterBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/nexus/{nexusId}/parameters/{id}][%d] deleteNexusParameterBadRequest ", 400)
}

func (o *DeleteNexusParameterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNexusParameterUnauthorized creates a DeleteNexusParameterUnauthorized with default headers values
func NewDeleteNexusParameterUnauthorized() *DeleteNexusParameterUnauthorized {
	return &DeleteNexusParameterUnauthorized{}
}

/* DeleteNexusParameterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteNexusParameterUnauthorized struct {
}

func (o *DeleteNexusParameterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/nexus/{nexusId}/parameters/{id}][%d] deleteNexusParameterUnauthorized ", 401)
}

func (o *DeleteNexusParameterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNexusParameterNotFound creates a DeleteNexusParameterNotFound with default headers values
func NewDeleteNexusParameterNotFound() *DeleteNexusParameterNotFound {
	return &DeleteNexusParameterNotFound{}
}

/* DeleteNexusParameterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteNexusParameterNotFound struct {
}

func (o *DeleteNexusParameterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/nexus/{nexusId}/parameters/{id}][%d] deleteNexusParameterNotFound ", 404)
}

func (o *DeleteNexusParameterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

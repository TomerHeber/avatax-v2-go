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

// GetLocationParameterReader is a Reader for the GetLocationParameter structure.
type GetLocationParameterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocationParameterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLocationParameterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLocationParameterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLocationParameterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLocationParameterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLocationParameterOK creates a GetLocationParameterOK with default headers values
func NewGetLocationParameterOK() *GetLocationParameterOK {
	return &GetLocationParameterOK{}
}

/* GetLocationParameterOK describes a response with status code 200, with default header values.

Success
*/
type GetLocationParameterOK struct {
	Payload *models.LocationParameterModel
}

func (o *GetLocationParameterOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/locations/{locationId}/parameters/{id}][%d] getLocationParameterOK  %+v", 200, o.Payload)
}
func (o *GetLocationParameterOK) GetPayload() *models.LocationParameterModel {
	return o.Payload
}

func (o *GetLocationParameterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationParameterModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationParameterBadRequest creates a GetLocationParameterBadRequest with default headers values
func NewGetLocationParameterBadRequest() *GetLocationParameterBadRequest {
	return &GetLocationParameterBadRequest{}
}

/* GetLocationParameterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLocationParameterBadRequest struct {
}

func (o *GetLocationParameterBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/locations/{locationId}/parameters/{id}][%d] getLocationParameterBadRequest ", 400)
}

func (o *GetLocationParameterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLocationParameterUnauthorized creates a GetLocationParameterUnauthorized with default headers values
func NewGetLocationParameterUnauthorized() *GetLocationParameterUnauthorized {
	return &GetLocationParameterUnauthorized{}
}

/* GetLocationParameterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLocationParameterUnauthorized struct {
}

func (o *GetLocationParameterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/locations/{locationId}/parameters/{id}][%d] getLocationParameterUnauthorized ", 401)
}

func (o *GetLocationParameterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLocationParameterNotFound creates a GetLocationParameterNotFound with default headers values
func NewGetLocationParameterNotFound() *GetLocationParameterNotFound {
	return &GetLocationParameterNotFound{}
}

/* GetLocationParameterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetLocationParameterNotFound struct {
}

func (o *GetLocationParameterNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/locations/{locationId}/parameters/{id}][%d] getLocationParameterNotFound ", 404)
}

func (o *GetLocationParameterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
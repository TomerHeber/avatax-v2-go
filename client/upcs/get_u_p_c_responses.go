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

// GetUPCReader is a Reader for the GetUPC structure.
type GetUPCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUPCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUPCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUPCBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUPCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUPCNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUPCOK creates a GetUPCOK with default headers values
func NewGetUPCOK() *GetUPCOK {
	return &GetUPCOK{}
}

/* GetUPCOK describes a response with status code 200, with default header values.

Success
*/
type GetUPCOK struct {
	Payload *models.UPCModel
}

func (o *GetUPCOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/upcs/{id}][%d] getUPCOK  %+v", 200, o.Payload)
}
func (o *GetUPCOK) GetPayload() *models.UPCModel {
	return o.Payload
}

func (o *GetUPCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UPCModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUPCBadRequest creates a GetUPCBadRequest with default headers values
func NewGetUPCBadRequest() *GetUPCBadRequest {
	return &GetUPCBadRequest{}
}

/* GetUPCBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetUPCBadRequest struct {
}

func (o *GetUPCBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/upcs/{id}][%d] getUPCBadRequest ", 400)
}

func (o *GetUPCBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUPCUnauthorized creates a GetUPCUnauthorized with default headers values
func NewGetUPCUnauthorized() *GetUPCUnauthorized {
	return &GetUPCUnauthorized{}
}

/* GetUPCUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetUPCUnauthorized struct {
}

func (o *GetUPCUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/upcs/{id}][%d] getUPCUnauthorized ", 401)
}

func (o *GetUPCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUPCNotFound creates a GetUPCNotFound with default headers values
func NewGetUPCNotFound() *GetUPCNotFound {
	return &GetUPCNotFound{}
}

/* GetUPCNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetUPCNotFound struct {
}

func (o *GetUPCNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/upcs/{id}][%d] getUPCNotFound ", 404)
}

func (o *GetUPCNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
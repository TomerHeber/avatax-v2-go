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

// DeclareNexusByAddressReader is a Reader for the DeclareNexusByAddress structure.
type DeclareNexusByAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeclareNexusByAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeclareNexusByAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeclareNexusByAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeclareNexusByAddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeclareNexusByAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeclareNexusByAddressOK creates a DeclareNexusByAddressOK with default headers values
func NewDeclareNexusByAddressOK() *DeclareNexusByAddressOK {
	return &DeclareNexusByAddressOK{}
}

/* DeclareNexusByAddressOK describes a response with status code 200, with default header values.

Success
*/
type DeclareNexusByAddressOK struct {
	Payload []*models.NexusByAddressModel
}

func (o *DeclareNexusByAddressOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/nexus/byaddress][%d] declareNexusByAddressOK  %+v", 200, o.Payload)
}
func (o *DeclareNexusByAddressOK) GetPayload() []*models.NexusByAddressModel {
	return o.Payload
}

func (o *DeclareNexusByAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeclareNexusByAddressBadRequest creates a DeclareNexusByAddressBadRequest with default headers values
func NewDeclareNexusByAddressBadRequest() *DeclareNexusByAddressBadRequest {
	return &DeclareNexusByAddressBadRequest{}
}

/* DeclareNexusByAddressBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeclareNexusByAddressBadRequest struct {
}

func (o *DeclareNexusByAddressBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/nexus/byaddress][%d] declareNexusByAddressBadRequest ", 400)
}

func (o *DeclareNexusByAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeclareNexusByAddressUnauthorized creates a DeclareNexusByAddressUnauthorized with default headers values
func NewDeclareNexusByAddressUnauthorized() *DeclareNexusByAddressUnauthorized {
	return &DeclareNexusByAddressUnauthorized{}
}

/* DeclareNexusByAddressUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeclareNexusByAddressUnauthorized struct {
}

func (o *DeclareNexusByAddressUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/nexus/byaddress][%d] declareNexusByAddressUnauthorized ", 401)
}

func (o *DeclareNexusByAddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeclareNexusByAddressNotFound creates a DeclareNexusByAddressNotFound with default headers values
func NewDeclareNexusByAddressNotFound() *DeclareNexusByAddressNotFound {
	return &DeclareNexusByAddressNotFound{}
}

/* DeclareNexusByAddressNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeclareNexusByAddressNotFound struct {
}

func (o *DeclareNexusByAddressNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/nexus/byaddress][%d] declareNexusByAddressNotFound ", 404)
}

func (o *DeclareNexusByAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package firm_client_linkages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// RevokeFirmClientLinkageReader is a Reader for the RevokeFirmClientLinkage structure.
type RevokeFirmClientLinkageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeFirmClientLinkageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeFirmClientLinkageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeFirmClientLinkageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRevokeFirmClientLinkageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeFirmClientLinkageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokeFirmClientLinkageOK creates a RevokeFirmClientLinkageOK with default headers values
func NewRevokeFirmClientLinkageOK() *RevokeFirmClientLinkageOK {
	return &RevokeFirmClientLinkageOK{}
}

/* RevokeFirmClientLinkageOK describes a response with status code 200, with default header values.

Success
*/
type RevokeFirmClientLinkageOK struct {
	Payload *models.FirmClientLinkageOutputModel
}

func (o *RevokeFirmClientLinkageOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/firmclientlinkages/{id}/revoke][%d] revokeFirmClientLinkageOK  %+v", 200, o.Payload)
}
func (o *RevokeFirmClientLinkageOK) GetPayload() *models.FirmClientLinkageOutputModel {
	return o.Payload
}

func (o *RevokeFirmClientLinkageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirmClientLinkageOutputModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeFirmClientLinkageBadRequest creates a RevokeFirmClientLinkageBadRequest with default headers values
func NewRevokeFirmClientLinkageBadRequest() *RevokeFirmClientLinkageBadRequest {
	return &RevokeFirmClientLinkageBadRequest{}
}

/* RevokeFirmClientLinkageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RevokeFirmClientLinkageBadRequest struct {
}

func (o *RevokeFirmClientLinkageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/firmclientlinkages/{id}/revoke][%d] revokeFirmClientLinkageBadRequest ", 400)
}

func (o *RevokeFirmClientLinkageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeFirmClientLinkageUnauthorized creates a RevokeFirmClientLinkageUnauthorized with default headers values
func NewRevokeFirmClientLinkageUnauthorized() *RevokeFirmClientLinkageUnauthorized {
	return &RevokeFirmClientLinkageUnauthorized{}
}

/* RevokeFirmClientLinkageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RevokeFirmClientLinkageUnauthorized struct {
}

func (o *RevokeFirmClientLinkageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/firmclientlinkages/{id}/revoke][%d] revokeFirmClientLinkageUnauthorized ", 401)
}

func (o *RevokeFirmClientLinkageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeFirmClientLinkageNotFound creates a RevokeFirmClientLinkageNotFound with default headers values
func NewRevokeFirmClientLinkageNotFound() *RevokeFirmClientLinkageNotFound {
	return &RevokeFirmClientLinkageNotFound{}
}

/* RevokeFirmClientLinkageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RevokeFirmClientLinkageNotFound struct {
}

func (o *RevokeFirmClientLinkageNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/firmclientlinkages/{id}/revoke][%d] revokeFirmClientLinkageNotFound ", 404)
}

func (o *RevokeFirmClientLinkageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
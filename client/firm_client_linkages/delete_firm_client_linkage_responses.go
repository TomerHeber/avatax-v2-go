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

// DeleteFirmClientLinkageReader is a Reader for the DeleteFirmClientLinkage structure.
type DeleteFirmClientLinkageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFirmClientLinkageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFirmClientLinkageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFirmClientLinkageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFirmClientLinkageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFirmClientLinkageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFirmClientLinkageOK creates a DeleteFirmClientLinkageOK with default headers values
func NewDeleteFirmClientLinkageOK() *DeleteFirmClientLinkageOK {
	return &DeleteFirmClientLinkageOK{}
}

/* DeleteFirmClientLinkageOK describes a response with status code 200, with default header values.

Success
*/
type DeleteFirmClientLinkageOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteFirmClientLinkageOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/firmclientlinkages/{id}][%d] deleteFirmClientLinkageOK  %+v", 200, o.Payload)
}
func (o *DeleteFirmClientLinkageOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteFirmClientLinkageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFirmClientLinkageBadRequest creates a DeleteFirmClientLinkageBadRequest with default headers values
func NewDeleteFirmClientLinkageBadRequest() *DeleteFirmClientLinkageBadRequest {
	return &DeleteFirmClientLinkageBadRequest{}
}

/* DeleteFirmClientLinkageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteFirmClientLinkageBadRequest struct {
}

func (o *DeleteFirmClientLinkageBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/firmclientlinkages/{id}][%d] deleteFirmClientLinkageBadRequest ", 400)
}

func (o *DeleteFirmClientLinkageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFirmClientLinkageUnauthorized creates a DeleteFirmClientLinkageUnauthorized with default headers values
func NewDeleteFirmClientLinkageUnauthorized() *DeleteFirmClientLinkageUnauthorized {
	return &DeleteFirmClientLinkageUnauthorized{}
}

/* DeleteFirmClientLinkageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteFirmClientLinkageUnauthorized struct {
}

func (o *DeleteFirmClientLinkageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/firmclientlinkages/{id}][%d] deleteFirmClientLinkageUnauthorized ", 401)
}

func (o *DeleteFirmClientLinkageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFirmClientLinkageNotFound creates a DeleteFirmClientLinkageNotFound with default headers values
func NewDeleteFirmClientLinkageNotFound() *DeleteFirmClientLinkageNotFound {
	return &DeleteFirmClientLinkageNotFound{}
}

/* DeleteFirmClientLinkageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteFirmClientLinkageNotFound struct {
}

func (o *DeleteFirmClientLinkageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/firmclientlinkages/{id}][%d] deleteFirmClientLinkageNotFound ", 404)
}

func (o *DeleteFirmClientLinkageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
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

// CreateFirmClientLinkageReader is a Reader for the CreateFirmClientLinkage structure.
type CreateFirmClientLinkageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFirmClientLinkageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFirmClientLinkageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFirmClientLinkageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFirmClientLinkageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFirmClientLinkageCreated creates a CreateFirmClientLinkageCreated with default headers values
func NewCreateFirmClientLinkageCreated() *CreateFirmClientLinkageCreated {
	return &CreateFirmClientLinkageCreated{}
}

/* CreateFirmClientLinkageCreated describes a response with status code 201, with default header values.

Success
*/
type CreateFirmClientLinkageCreated struct {
	Payload *models.FirmClientLinkageOutputModel
}

func (o *CreateFirmClientLinkageCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/firmclientlinkages][%d] createFirmClientLinkageCreated  %+v", 201, o.Payload)
}
func (o *CreateFirmClientLinkageCreated) GetPayload() *models.FirmClientLinkageOutputModel {
	return o.Payload
}

func (o *CreateFirmClientLinkageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirmClientLinkageOutputModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFirmClientLinkageBadRequest creates a CreateFirmClientLinkageBadRequest with default headers values
func NewCreateFirmClientLinkageBadRequest() *CreateFirmClientLinkageBadRequest {
	return &CreateFirmClientLinkageBadRequest{}
}

/* CreateFirmClientLinkageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateFirmClientLinkageBadRequest struct {
}

func (o *CreateFirmClientLinkageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/firmclientlinkages][%d] createFirmClientLinkageBadRequest ", 400)
}

func (o *CreateFirmClientLinkageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFirmClientLinkageUnauthorized creates a CreateFirmClientLinkageUnauthorized with default headers values
func NewCreateFirmClientLinkageUnauthorized() *CreateFirmClientLinkageUnauthorized {
	return &CreateFirmClientLinkageUnauthorized{}
}

/* CreateFirmClientLinkageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateFirmClientLinkageUnauthorized struct {
}

func (o *CreateFirmClientLinkageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/firmclientlinkages][%d] createFirmClientLinkageUnauthorized ", 401)
}

func (o *CreateFirmClientLinkageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
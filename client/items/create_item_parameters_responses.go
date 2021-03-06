// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// CreateItemParametersReader is a Reader for the CreateItemParameters structure.
type CreateItemParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateItemParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateItemParametersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateItemParametersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateItemParametersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateItemParametersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateItemParametersCreated creates a CreateItemParametersCreated with default headers values
func NewCreateItemParametersCreated() *CreateItemParametersCreated {
	return &CreateItemParametersCreated{}
}

/* CreateItemParametersCreated describes a response with status code 201, with default header values.

Success
*/
type CreateItemParametersCreated struct {
	Payload []*models.ItemParameterModel
}

func (o *CreateItemParametersCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/items/{itemId}/parameters][%d] createItemParametersCreated  %+v", 201, o.Payload)
}
func (o *CreateItemParametersCreated) GetPayload() []*models.ItemParameterModel {
	return o.Payload
}

func (o *CreateItemParametersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateItemParametersBadRequest creates a CreateItemParametersBadRequest with default headers values
func NewCreateItemParametersBadRequest() *CreateItemParametersBadRequest {
	return &CreateItemParametersBadRequest{}
}

/* CreateItemParametersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateItemParametersBadRequest struct {
}

func (o *CreateItemParametersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/items/{itemId}/parameters][%d] createItemParametersBadRequest ", 400)
}

func (o *CreateItemParametersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateItemParametersUnauthorized creates a CreateItemParametersUnauthorized with default headers values
func NewCreateItemParametersUnauthorized() *CreateItemParametersUnauthorized {
	return &CreateItemParametersUnauthorized{}
}

/* CreateItemParametersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateItemParametersUnauthorized struct {
}

func (o *CreateItemParametersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/items/{itemId}/parameters][%d] createItemParametersUnauthorized ", 401)
}

func (o *CreateItemParametersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateItemParametersNotFound creates a CreateItemParametersNotFound with default headers values
func NewCreateItemParametersNotFound() *CreateItemParametersNotFound {
	return &CreateItemParametersNotFound{}
}

/* CreateItemParametersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateItemParametersNotFound struct {
}

func (o *CreateItemParametersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/items/{itemId}/parameters][%d] createItemParametersNotFound ", 404)
}

func (o *CreateItemParametersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

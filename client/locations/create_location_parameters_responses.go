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

// CreateLocationParametersReader is a Reader for the CreateLocationParameters structure.
type CreateLocationParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLocationParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateLocationParametersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLocationParametersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateLocationParametersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateLocationParametersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLocationParametersCreated creates a CreateLocationParametersCreated with default headers values
func NewCreateLocationParametersCreated() *CreateLocationParametersCreated {
	return &CreateLocationParametersCreated{}
}

/* CreateLocationParametersCreated describes a response with status code 201, with default header values.

Success
*/
type CreateLocationParametersCreated struct {
	Payload []*models.LocationParameterModel
}

func (o *CreateLocationParametersCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/locations/{locationId}/parameters][%d] createLocationParametersCreated  %+v", 201, o.Payload)
}
func (o *CreateLocationParametersCreated) GetPayload() []*models.LocationParameterModel {
	return o.Payload
}

func (o *CreateLocationParametersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLocationParametersBadRequest creates a CreateLocationParametersBadRequest with default headers values
func NewCreateLocationParametersBadRequest() *CreateLocationParametersBadRequest {
	return &CreateLocationParametersBadRequest{}
}

/* CreateLocationParametersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateLocationParametersBadRequest struct {
}

func (o *CreateLocationParametersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/locations/{locationId}/parameters][%d] createLocationParametersBadRequest ", 400)
}

func (o *CreateLocationParametersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLocationParametersUnauthorized creates a CreateLocationParametersUnauthorized with default headers values
func NewCreateLocationParametersUnauthorized() *CreateLocationParametersUnauthorized {
	return &CreateLocationParametersUnauthorized{}
}

/* CreateLocationParametersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateLocationParametersUnauthorized struct {
}

func (o *CreateLocationParametersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/locations/{locationId}/parameters][%d] createLocationParametersUnauthorized ", 401)
}

func (o *CreateLocationParametersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLocationParametersNotFound creates a CreateLocationParametersNotFound with default headers values
func NewCreateLocationParametersNotFound() *CreateLocationParametersNotFound {
	return &CreateLocationParametersNotFound{}
}

/* CreateLocationParametersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateLocationParametersNotFound struct {
}

func (o *CreateLocationParametersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/locations/{locationId}/parameters][%d] createLocationParametersNotFound ", 404)
}

func (o *CreateLocationParametersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package companies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// CreateCompanyParametersReader is a Reader for the CreateCompanyParameters structure.
type CreateCompanyParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCompanyParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCompanyParametersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCompanyParametersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCompanyParametersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCompanyParametersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCompanyParametersCreated creates a CreateCompanyParametersCreated with default headers values
func NewCreateCompanyParametersCreated() *CreateCompanyParametersCreated {
	return &CreateCompanyParametersCreated{}
}

/* CreateCompanyParametersCreated describes a response with status code 201, with default header values.

Success
*/
type CreateCompanyParametersCreated struct {
	Payload []*models.CompanyParameterDetailModel
}

func (o *CreateCompanyParametersCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/parameters][%d] createCompanyParametersCreated  %+v", 201, o.Payload)
}
func (o *CreateCompanyParametersCreated) GetPayload() []*models.CompanyParameterDetailModel {
	return o.Payload
}

func (o *CreateCompanyParametersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCompanyParametersBadRequest creates a CreateCompanyParametersBadRequest with default headers values
func NewCreateCompanyParametersBadRequest() *CreateCompanyParametersBadRequest {
	return &CreateCompanyParametersBadRequest{}
}

/* CreateCompanyParametersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateCompanyParametersBadRequest struct {
}

func (o *CreateCompanyParametersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/parameters][%d] createCompanyParametersBadRequest ", 400)
}

func (o *CreateCompanyParametersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCompanyParametersUnauthorized creates a CreateCompanyParametersUnauthorized with default headers values
func NewCreateCompanyParametersUnauthorized() *CreateCompanyParametersUnauthorized {
	return &CreateCompanyParametersUnauthorized{}
}

/* CreateCompanyParametersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateCompanyParametersUnauthorized struct {
}

func (o *CreateCompanyParametersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/parameters][%d] createCompanyParametersUnauthorized ", 401)
}

func (o *CreateCompanyParametersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCompanyParametersNotFound creates a CreateCompanyParametersNotFound with default headers values
func NewCreateCompanyParametersNotFound() *CreateCompanyParametersNotFound {
	return &CreateCompanyParametersNotFound{}
}

/* CreateCompanyParametersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateCompanyParametersNotFound struct {
}

func (o *CreateCompanyParametersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/parameters][%d] createCompanyParametersNotFound ", 404)
}

func (o *CreateCompanyParametersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

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

// SetCompanyConfigurationReader is a Reader for the SetCompanyConfiguration structure.
type SetCompanyConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetCompanyConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetCompanyConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetCompanyConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetCompanyConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetCompanyConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetCompanyConfigurationOK creates a SetCompanyConfigurationOK with default headers values
func NewSetCompanyConfigurationOK() *SetCompanyConfigurationOK {
	return &SetCompanyConfigurationOK{}
}

/* SetCompanyConfigurationOK describes a response with status code 200, with default header values.

Success
*/
type SetCompanyConfigurationOK struct {
	Payload []*models.CompanyConfigurationModel
}

func (o *SetCompanyConfigurationOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{id}/configuration][%d] setCompanyConfigurationOK  %+v", 200, o.Payload)
}
func (o *SetCompanyConfigurationOK) GetPayload() []*models.CompanyConfigurationModel {
	return o.Payload
}

func (o *SetCompanyConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetCompanyConfigurationBadRequest creates a SetCompanyConfigurationBadRequest with default headers values
func NewSetCompanyConfigurationBadRequest() *SetCompanyConfigurationBadRequest {
	return &SetCompanyConfigurationBadRequest{}
}

/* SetCompanyConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SetCompanyConfigurationBadRequest struct {
}

func (o *SetCompanyConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{id}/configuration][%d] setCompanyConfigurationBadRequest ", 400)
}

func (o *SetCompanyConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCompanyConfigurationUnauthorized creates a SetCompanyConfigurationUnauthorized with default headers values
func NewSetCompanyConfigurationUnauthorized() *SetCompanyConfigurationUnauthorized {
	return &SetCompanyConfigurationUnauthorized{}
}

/* SetCompanyConfigurationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SetCompanyConfigurationUnauthorized struct {
}

func (o *SetCompanyConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{id}/configuration][%d] setCompanyConfigurationUnauthorized ", 401)
}

func (o *SetCompanyConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCompanyConfigurationNotFound creates a SetCompanyConfigurationNotFound with default headers values
func NewSetCompanyConfigurationNotFound() *SetCompanyConfigurationNotFound {
	return &SetCompanyConfigurationNotFound{}
}

/* SetCompanyConfigurationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SetCompanyConfigurationNotFound struct {
}

func (o *SetCompanyConfigurationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{id}/configuration][%d] setCompanyConfigurationNotFound ", 404)
}

func (o *SetCompanyConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package data_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// CreateDataSourcesReader is a Reader for the CreateDataSources structure.
type CreateDataSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDataSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDataSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDataSourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDataSourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDataSourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDataSourcesOK creates a CreateDataSourcesOK with default headers values
func NewCreateDataSourcesOK() *CreateDataSourcesOK {
	return &CreateDataSourcesOK{}
}

/* CreateDataSourcesOK describes a response with status code 200, with default header values.

Success
*/
type CreateDataSourcesOK struct {
	Payload []*models.DataSourceModel
}

func (o *CreateDataSourcesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/datasources][%d] createDataSourcesOK  %+v", 200, o.Payload)
}
func (o *CreateDataSourcesOK) GetPayload() []*models.DataSourceModel {
	return o.Payload
}

func (o *CreateDataSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataSourcesBadRequest creates a CreateDataSourcesBadRequest with default headers values
func NewCreateDataSourcesBadRequest() *CreateDataSourcesBadRequest {
	return &CreateDataSourcesBadRequest{}
}

/* CreateDataSourcesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateDataSourcesBadRequest struct {
}

func (o *CreateDataSourcesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/datasources][%d] createDataSourcesBadRequest ", 400)
}

func (o *CreateDataSourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDataSourcesUnauthorized creates a CreateDataSourcesUnauthorized with default headers values
func NewCreateDataSourcesUnauthorized() *CreateDataSourcesUnauthorized {
	return &CreateDataSourcesUnauthorized{}
}

/* CreateDataSourcesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateDataSourcesUnauthorized struct {
}

func (o *CreateDataSourcesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/datasources][%d] createDataSourcesUnauthorized ", 401)
}

func (o *CreateDataSourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDataSourcesNotFound creates a CreateDataSourcesNotFound with default headers values
func NewCreateDataSourcesNotFound() *CreateDataSourcesNotFound {
	return &CreateDataSourcesNotFound{}
}

/* CreateDataSourcesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateDataSourcesNotFound struct {
}

func (o *CreateDataSourcesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/datasources][%d] createDataSourcesNotFound ", 404)
}

func (o *CreateDataSourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package companies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeFilingStatusReader is a Reader for the ChangeFilingStatus structure.
type ChangeFilingStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeFilingStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeFilingStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChangeFilingStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewChangeFilingStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChangeFilingStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChangeFilingStatusOK creates a ChangeFilingStatusOK with default headers values
func NewChangeFilingStatusOK() *ChangeFilingStatusOK {
	return &ChangeFilingStatusOK{}
}

/* ChangeFilingStatusOK describes a response with status code 200, with default header values.

Success
*/
type ChangeFilingStatusOK struct {
	Payload string
}

func (o *ChangeFilingStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{id}/filingstatus][%d] changeFilingStatusOK  %+v", 200, o.Payload)
}
func (o *ChangeFilingStatusOK) GetPayload() string {
	return o.Payload
}

func (o *ChangeFilingStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeFilingStatusBadRequest creates a ChangeFilingStatusBadRequest with default headers values
func NewChangeFilingStatusBadRequest() *ChangeFilingStatusBadRequest {
	return &ChangeFilingStatusBadRequest{}
}

/* ChangeFilingStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ChangeFilingStatusBadRequest struct {
}

func (o *ChangeFilingStatusBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{id}/filingstatus][%d] changeFilingStatusBadRequest ", 400)
}

func (o *ChangeFilingStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeFilingStatusUnauthorized creates a ChangeFilingStatusUnauthorized with default headers values
func NewChangeFilingStatusUnauthorized() *ChangeFilingStatusUnauthorized {
	return &ChangeFilingStatusUnauthorized{}
}

/* ChangeFilingStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ChangeFilingStatusUnauthorized struct {
}

func (o *ChangeFilingStatusUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{id}/filingstatus][%d] changeFilingStatusUnauthorized ", 401)
}

func (o *ChangeFilingStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeFilingStatusNotFound creates a ChangeFilingStatusNotFound with default headers values
func NewChangeFilingStatusNotFound() *ChangeFilingStatusNotFound {
	return &ChangeFilingStatusNotFound{}
}

/* ChangeFilingStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ChangeFilingStatusNotFound struct {
}

func (o *ChangeFilingStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{id}/filingstatus][%d] changeFilingStatusNotFound ", 404)
}

func (o *ChangeFilingStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package ava_file_forms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// GetAvaFileFormReader is a Reader for the GetAvaFileForm structure.
type GetAvaFileFormReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvaFileFormReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAvaFileFormOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAvaFileFormBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAvaFileFormUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAvaFileFormNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAvaFileFormOK creates a GetAvaFileFormOK with default headers values
func NewGetAvaFileFormOK() *GetAvaFileFormOK {
	return &GetAvaFileFormOK{}
}

/* GetAvaFileFormOK describes a response with status code 200, with default header values.

Success
*/
type GetAvaFileFormOK struct {
	Payload *models.AvaFileFormModel
}

func (o *GetAvaFileFormOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/avafileforms/{id}][%d] getAvaFileFormOK  %+v", 200, o.Payload)
}
func (o *GetAvaFileFormOK) GetPayload() *models.AvaFileFormModel {
	return o.Payload
}

func (o *GetAvaFileFormOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvaFileFormModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvaFileFormBadRequest creates a GetAvaFileFormBadRequest with default headers values
func NewGetAvaFileFormBadRequest() *GetAvaFileFormBadRequest {
	return &GetAvaFileFormBadRequest{}
}

/* GetAvaFileFormBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAvaFileFormBadRequest struct {
}

func (o *GetAvaFileFormBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/avafileforms/{id}][%d] getAvaFileFormBadRequest ", 400)
}

func (o *GetAvaFileFormBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAvaFileFormUnauthorized creates a GetAvaFileFormUnauthorized with default headers values
func NewGetAvaFileFormUnauthorized() *GetAvaFileFormUnauthorized {
	return &GetAvaFileFormUnauthorized{}
}

/* GetAvaFileFormUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAvaFileFormUnauthorized struct {
}

func (o *GetAvaFileFormUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/avafileforms/{id}][%d] getAvaFileFormUnauthorized ", 401)
}

func (o *GetAvaFileFormUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAvaFileFormNotFound creates a GetAvaFileFormNotFound with default headers values
func NewGetAvaFileFormNotFound() *GetAvaFileFormNotFound {
	return &GetAvaFileFormNotFound{}
}

/* GetAvaFileFormNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAvaFileFormNotFound struct {
}

func (o *GetAvaFileFormNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/avafileforms/{id}][%d] getAvaFileFormNotFound ", 404)
}

func (o *GetAvaFileFormNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
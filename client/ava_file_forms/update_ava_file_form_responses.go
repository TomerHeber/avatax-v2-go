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

// UpdateAvaFileFormReader is a Reader for the UpdateAvaFileForm structure.
type UpdateAvaFileFormReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAvaFileFormReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAvaFileFormOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAvaFileFormBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAvaFileFormUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAvaFileFormNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAvaFileFormOK creates a UpdateAvaFileFormOK with default headers values
func NewUpdateAvaFileFormOK() *UpdateAvaFileFormOK {
	return &UpdateAvaFileFormOK{}
}

/* UpdateAvaFileFormOK describes a response with status code 200, with default header values.

Success
*/
type UpdateAvaFileFormOK struct {
	Payload *models.AvaFileFormModel
}

func (o *UpdateAvaFileFormOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/avafileforms/{id}][%d] updateAvaFileFormOK  %+v", 200, o.Payload)
}
func (o *UpdateAvaFileFormOK) GetPayload() *models.AvaFileFormModel {
	return o.Payload
}

func (o *UpdateAvaFileFormOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvaFileFormModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAvaFileFormBadRequest creates a UpdateAvaFileFormBadRequest with default headers values
func NewUpdateAvaFileFormBadRequest() *UpdateAvaFileFormBadRequest {
	return &UpdateAvaFileFormBadRequest{}
}

/* UpdateAvaFileFormBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateAvaFileFormBadRequest struct {
}

func (o *UpdateAvaFileFormBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/avafileforms/{id}][%d] updateAvaFileFormBadRequest ", 400)
}

func (o *UpdateAvaFileFormBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAvaFileFormUnauthorized creates a UpdateAvaFileFormUnauthorized with default headers values
func NewUpdateAvaFileFormUnauthorized() *UpdateAvaFileFormUnauthorized {
	return &UpdateAvaFileFormUnauthorized{}
}

/* UpdateAvaFileFormUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateAvaFileFormUnauthorized struct {
}

func (o *UpdateAvaFileFormUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/avafileforms/{id}][%d] updateAvaFileFormUnauthorized ", 401)
}

func (o *UpdateAvaFileFormUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAvaFileFormNotFound creates a UpdateAvaFileFormNotFound with default headers values
func NewUpdateAvaFileFormNotFound() *UpdateAvaFileFormNotFound {
	return &UpdateAvaFileFormNotFound{}
}

/* UpdateAvaFileFormNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateAvaFileFormNotFound struct {
}

func (o *UpdateAvaFileFormNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/avafileforms/{id}][%d] updateAvaFileFormNotFound ", 404)
}

func (o *UpdateAvaFileFormNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
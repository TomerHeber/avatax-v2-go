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

// DeleteAvaFileFormReader is a Reader for the DeleteAvaFileForm structure.
type DeleteAvaFileFormReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAvaFileFormReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAvaFileFormOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAvaFileFormBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAvaFileFormUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAvaFileFormNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAvaFileFormOK creates a DeleteAvaFileFormOK with default headers values
func NewDeleteAvaFileFormOK() *DeleteAvaFileFormOK {
	return &DeleteAvaFileFormOK{}
}

/* DeleteAvaFileFormOK describes a response with status code 200, with default header values.

Success
*/
type DeleteAvaFileFormOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteAvaFileFormOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/avafileforms/{id}][%d] deleteAvaFileFormOK  %+v", 200, o.Payload)
}
func (o *DeleteAvaFileFormOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteAvaFileFormOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAvaFileFormBadRequest creates a DeleteAvaFileFormBadRequest with default headers values
func NewDeleteAvaFileFormBadRequest() *DeleteAvaFileFormBadRequest {
	return &DeleteAvaFileFormBadRequest{}
}

/* DeleteAvaFileFormBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAvaFileFormBadRequest struct {
}

func (o *DeleteAvaFileFormBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/avafileforms/{id}][%d] deleteAvaFileFormBadRequest ", 400)
}

func (o *DeleteAvaFileFormBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAvaFileFormUnauthorized creates a DeleteAvaFileFormUnauthorized with default headers values
func NewDeleteAvaFileFormUnauthorized() *DeleteAvaFileFormUnauthorized {
	return &DeleteAvaFileFormUnauthorized{}
}

/* DeleteAvaFileFormUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteAvaFileFormUnauthorized struct {
}

func (o *DeleteAvaFileFormUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/avafileforms/{id}][%d] deleteAvaFileFormUnauthorized ", 401)
}

func (o *DeleteAvaFileFormUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAvaFileFormNotFound creates a DeleteAvaFileFormNotFound with default headers values
func NewDeleteAvaFileFormNotFound() *DeleteAvaFileFormNotFound {
	return &DeleteAvaFileFormNotFound{}
}

/* DeleteAvaFileFormNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAvaFileFormNotFound struct {
}

func (o *DeleteAvaFileFormNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/avafileforms/{id}][%d] deleteAvaFileFormNotFound ", 404)
}

func (o *DeleteAvaFileFormNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package notices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// CreateNoticeResponsibilityTypeReader is a Reader for the CreateNoticeResponsibilityType structure.
type CreateNoticeResponsibilityTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNoticeResponsibilityTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNoticeResponsibilityTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNoticeResponsibilityTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateNoticeResponsibilityTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateNoticeResponsibilityTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNoticeResponsibilityTypeOK creates a CreateNoticeResponsibilityTypeOK with default headers values
func NewCreateNoticeResponsibilityTypeOK() *CreateNoticeResponsibilityTypeOK {
	return &CreateNoticeResponsibilityTypeOK{}
}

/* CreateNoticeResponsibilityTypeOK describes a response with status code 200, with default header values.

Success
*/
type CreateNoticeResponsibilityTypeOK struct {
	Payload *models.NoticeResponsibilityModel
}

func (o *CreateNoticeResponsibilityTypeOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/notices/responsibilities][%d] createNoticeResponsibilityTypeOK  %+v", 200, o.Payload)
}
func (o *CreateNoticeResponsibilityTypeOK) GetPayload() *models.NoticeResponsibilityModel {
	return o.Payload
}

func (o *CreateNoticeResponsibilityTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NoticeResponsibilityModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNoticeResponsibilityTypeBadRequest creates a CreateNoticeResponsibilityTypeBadRequest with default headers values
func NewCreateNoticeResponsibilityTypeBadRequest() *CreateNoticeResponsibilityTypeBadRequest {
	return &CreateNoticeResponsibilityTypeBadRequest{}
}

/* CreateNoticeResponsibilityTypeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateNoticeResponsibilityTypeBadRequest struct {
}

func (o *CreateNoticeResponsibilityTypeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/notices/responsibilities][%d] createNoticeResponsibilityTypeBadRequest ", 400)
}

func (o *CreateNoticeResponsibilityTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNoticeResponsibilityTypeUnauthorized creates a CreateNoticeResponsibilityTypeUnauthorized with default headers values
func NewCreateNoticeResponsibilityTypeUnauthorized() *CreateNoticeResponsibilityTypeUnauthorized {
	return &CreateNoticeResponsibilityTypeUnauthorized{}
}

/* CreateNoticeResponsibilityTypeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateNoticeResponsibilityTypeUnauthorized struct {
}

func (o *CreateNoticeResponsibilityTypeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/notices/responsibilities][%d] createNoticeResponsibilityTypeUnauthorized ", 401)
}

func (o *CreateNoticeResponsibilityTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNoticeResponsibilityTypeNotFound creates a CreateNoticeResponsibilityTypeNotFound with default headers values
func NewCreateNoticeResponsibilityTypeNotFound() *CreateNoticeResponsibilityTypeNotFound {
	return &CreateNoticeResponsibilityTypeNotFound{}
}

/* CreateNoticeResponsibilityTypeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateNoticeResponsibilityTypeNotFound struct {
}

func (o *CreateNoticeResponsibilityTypeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/notices/responsibilities][%d] createNoticeResponsibilityTypeNotFound ", 404)
}

func (o *CreateNoticeResponsibilityTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package registrar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// UpdateNotificationReader is a Reader for the UpdateNotification structure.
type UpdateNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNotificationOK creates a UpdateNotificationOK with default headers values
func NewUpdateNotificationOK() *UpdateNotificationOK {
	return &UpdateNotificationOK{}
}

/* UpdateNotificationOK describes a response with status code 200, with default header values.

Success
*/
type UpdateNotificationOK struct {
	Payload *models.NotificationModel
}

func (o *UpdateNotificationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/notifications/{id}][%d] updateNotificationOK  %+v", 200, o.Payload)
}
func (o *UpdateNotificationOK) GetPayload() *models.NotificationModel {
	return o.Payload
}

func (o *UpdateNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationBadRequest creates a UpdateNotificationBadRequest with default headers values
func NewUpdateNotificationBadRequest() *UpdateNotificationBadRequest {
	return &UpdateNotificationBadRequest{}
}

/* UpdateNotificationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateNotificationBadRequest struct {
}

func (o *UpdateNotificationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/notifications/{id}][%d] updateNotificationBadRequest ", 400)
}

func (o *UpdateNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNotificationUnauthorized creates a UpdateNotificationUnauthorized with default headers values
func NewUpdateNotificationUnauthorized() *UpdateNotificationUnauthorized {
	return &UpdateNotificationUnauthorized{}
}

/* UpdateNotificationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateNotificationUnauthorized struct {
}

func (o *UpdateNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/notifications/{id}][%d] updateNotificationUnauthorized ", 401)
}

func (o *UpdateNotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNotificationNotFound creates a UpdateNotificationNotFound with default headers values
func NewUpdateNotificationNotFound() *UpdateNotificationNotFound {
	return &UpdateNotificationNotFound{}
}

/* UpdateNotificationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateNotificationNotFound struct {
}

func (o *UpdateNotificationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/notifications/{id}][%d] updateNotificationNotFound ", 404)
}

func (o *UpdateNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

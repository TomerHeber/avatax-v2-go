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

// UpdateSubscriptionReader is a Reader for the UpdateSubscription structure.
type UpdateSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateSubscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSubscriptionOK creates a UpdateSubscriptionOK with default headers values
func NewUpdateSubscriptionOK() *UpdateSubscriptionOK {
	return &UpdateSubscriptionOK{}
}

/* UpdateSubscriptionOK describes a response with status code 200, with default header values.

Success
*/
type UpdateSubscriptionOK struct {
	Payload *models.SubscriptionModel
}

func (o *UpdateSubscriptionOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/accounts/{accountId}/subscriptions/{id}][%d] updateSubscriptionOK  %+v", 200, o.Payload)
}
func (o *UpdateSubscriptionOK) GetPayload() *models.SubscriptionModel {
	return o.Payload
}

func (o *UpdateSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionBadRequest creates a UpdateSubscriptionBadRequest with default headers values
func NewUpdateSubscriptionBadRequest() *UpdateSubscriptionBadRequest {
	return &UpdateSubscriptionBadRequest{}
}

/* UpdateSubscriptionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateSubscriptionBadRequest struct {
}

func (o *UpdateSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/accounts/{accountId}/subscriptions/{id}][%d] updateSubscriptionBadRequest ", 400)
}

func (o *UpdateSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSubscriptionUnauthorized creates a UpdateSubscriptionUnauthorized with default headers values
func NewUpdateSubscriptionUnauthorized() *UpdateSubscriptionUnauthorized {
	return &UpdateSubscriptionUnauthorized{}
}

/* UpdateSubscriptionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateSubscriptionUnauthorized struct {
}

func (o *UpdateSubscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/accounts/{accountId}/subscriptions/{id}][%d] updateSubscriptionUnauthorized ", 401)
}

func (o *UpdateSubscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSubscriptionNotFound creates a UpdateSubscriptionNotFound with default headers values
func NewUpdateSubscriptionNotFound() *UpdateSubscriptionNotFound {
	return &UpdateSubscriptionNotFound{}
}

/* UpdateSubscriptionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateSubscriptionNotFound struct {
}

func (o *UpdateSubscriptionNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/accounts/{accountId}/subscriptions/{id}][%d] updateSubscriptionNotFound ", 404)
}

func (o *UpdateSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
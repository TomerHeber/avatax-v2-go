// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// GetSubscriptionReader is a Reader for the GetSubscription structure.
type GetSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSubscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSubscriptionOK creates a GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

/* GetSubscriptionOK describes a response with status code 200, with default header values.

Success
*/
type GetSubscriptionOK struct {
	Payload *models.SubscriptionModel
}

func (o *GetSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{accountId}/subscriptions/{id}][%d] getSubscriptionOK  %+v", 200, o.Payload)
}
func (o *GetSubscriptionOK) GetPayload() *models.SubscriptionModel {
	return o.Payload
}

func (o *GetSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionBadRequest creates a GetSubscriptionBadRequest with default headers values
func NewGetSubscriptionBadRequest() *GetSubscriptionBadRequest {
	return &GetSubscriptionBadRequest{}
}

/* GetSubscriptionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSubscriptionBadRequest struct {
}

func (o *GetSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{accountId}/subscriptions/{id}][%d] getSubscriptionBadRequest ", 400)
}

func (o *GetSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionUnauthorized creates a GetSubscriptionUnauthorized with default headers values
func NewGetSubscriptionUnauthorized() *GetSubscriptionUnauthorized {
	return &GetSubscriptionUnauthorized{}
}

/* GetSubscriptionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSubscriptionUnauthorized struct {
}

func (o *GetSubscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{accountId}/subscriptions/{id}][%d] getSubscriptionUnauthorized ", 401)
}

func (o *GetSubscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionNotFound creates a GetSubscriptionNotFound with default headers values
func NewGetSubscriptionNotFound() *GetSubscriptionNotFound {
	return &GetSubscriptionNotFound{}
}

/* GetSubscriptionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSubscriptionNotFound struct {
}

func (o *GetSubscriptionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{accountId}/subscriptions/{id}][%d] getSubscriptionNotFound ", 404)
}

func (o *GetSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
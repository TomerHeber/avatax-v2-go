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

// DeleteSubscriptionReader is a Reader for the DeleteSubscription structure.
type DeleteSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSubscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSubscriptionOK creates a DeleteSubscriptionOK with default headers values
func NewDeleteSubscriptionOK() *DeleteSubscriptionOK {
	return &DeleteSubscriptionOK{}
}

/* DeleteSubscriptionOK describes a response with status code 200, with default header values.

Success
*/
type DeleteSubscriptionOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteSubscriptionOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{accountId}/subscriptions/{id}][%d] deleteSubscriptionOK  %+v", 200, o.Payload)
}
func (o *DeleteSubscriptionOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubscriptionBadRequest creates a DeleteSubscriptionBadRequest with default headers values
func NewDeleteSubscriptionBadRequest() *DeleteSubscriptionBadRequest {
	return &DeleteSubscriptionBadRequest{}
}

/* DeleteSubscriptionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteSubscriptionBadRequest struct {
}

func (o *DeleteSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{accountId}/subscriptions/{id}][%d] deleteSubscriptionBadRequest ", 400)
}

func (o *DeleteSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubscriptionUnauthorized creates a DeleteSubscriptionUnauthorized with default headers values
func NewDeleteSubscriptionUnauthorized() *DeleteSubscriptionUnauthorized {
	return &DeleteSubscriptionUnauthorized{}
}

/* DeleteSubscriptionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteSubscriptionUnauthorized struct {
}

func (o *DeleteSubscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{accountId}/subscriptions/{id}][%d] deleteSubscriptionUnauthorized ", 401)
}

func (o *DeleteSubscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubscriptionNotFound creates a DeleteSubscriptionNotFound with default headers values
func NewDeleteSubscriptionNotFound() *DeleteSubscriptionNotFound {
	return &DeleteSubscriptionNotFound{}
}

/* DeleteSubscriptionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteSubscriptionNotFound struct {
}

func (o *DeleteSubscriptionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{accountId}/subscriptions/{id}][%d] deleteSubscriptionNotFound ", 404)
}

func (o *DeleteSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
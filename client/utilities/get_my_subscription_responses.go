// Code generated by go-swagger; DO NOT EDIT.

package utilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// GetMySubscriptionReader is a Reader for the GetMySubscription structure.
type GetMySubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMySubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMySubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMySubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMySubscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMySubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMySubscriptionOK creates a GetMySubscriptionOK with default headers values
func NewGetMySubscriptionOK() *GetMySubscriptionOK {
	return &GetMySubscriptionOK{}
}

/* GetMySubscriptionOK describes a response with status code 200, with default header values.

Success
*/
type GetMySubscriptionOK struct {
	Payload *models.SubscriptionModel
}

func (o *GetMySubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/utilities/subscriptions/{serviceTypeId}][%d] getMySubscriptionOK  %+v", 200, o.Payload)
}
func (o *GetMySubscriptionOK) GetPayload() *models.SubscriptionModel {
	return o.Payload
}

func (o *GetMySubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMySubscriptionBadRequest creates a GetMySubscriptionBadRequest with default headers values
func NewGetMySubscriptionBadRequest() *GetMySubscriptionBadRequest {
	return &GetMySubscriptionBadRequest{}
}

/* GetMySubscriptionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetMySubscriptionBadRequest struct {
}

func (o *GetMySubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/utilities/subscriptions/{serviceTypeId}][%d] getMySubscriptionBadRequest ", 400)
}

func (o *GetMySubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMySubscriptionUnauthorized creates a GetMySubscriptionUnauthorized with default headers values
func NewGetMySubscriptionUnauthorized() *GetMySubscriptionUnauthorized {
	return &GetMySubscriptionUnauthorized{}
}

/* GetMySubscriptionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetMySubscriptionUnauthorized struct {
}

func (o *GetMySubscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/utilities/subscriptions/{serviceTypeId}][%d] getMySubscriptionUnauthorized ", 401)
}

func (o *GetMySubscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMySubscriptionNotFound creates a GetMySubscriptionNotFound with default headers values
func NewGetMySubscriptionNotFound() *GetMySubscriptionNotFound {
	return &GetMySubscriptionNotFound{}
}

/* GetMySubscriptionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetMySubscriptionNotFound struct {
}

func (o *GetMySubscriptionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/utilities/subscriptions/{serviceTypeId}][%d] getMySubscriptionNotFound ", 404)
}

func (o *GetMySubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}